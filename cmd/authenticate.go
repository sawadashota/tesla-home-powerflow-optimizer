package cmd

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/lestrrat-go/jwx/jwk"

	"github.com/lestrrat-go/jwx/jwt"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/model"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/driver"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/internal/randx"
	"github.com/spf13/cobra"
	"github.com/toqueteos/webbrowser"
	"golang.org/x/oauth2"
)

func newAuthenticateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "authenticate",
		Short: "Sign in with your Tesla account to call the Tesla API",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			r, err := driver.NewOidcRegistry()
			if err != nil {
				return err
			}
			r.Logger().Info("Migrating database...")
			if err := r.Migrate(ctx); err != nil {
				return err
			}

			r.Logger().Info("Starting sign in with Tesla")
			grant, err := signInWithTesla(ctx, r)
			if err != nil {
				return err
			}
			return r.GrantService().Save(ctx, grant)
		},
	}

	return cmd
}

var oidcWelcomeHTML = template.Must(template.New("").Parse(`<html>
<body>
<h1>Welcome to the Tesla Home PowerFlow Optimizer!</h1>
<p><a href="{{ .URL }}">Sign in with Tesla</a></p>
</body>
</html>`))

var oidcErrorHTML = template.Must(template.New("").Parse(`<html>
<body>
<h1>An error occurred</h1>
<h2>{{ .Name }}</h2>
<p>{{ .Description }}</p>
<p>{{ .Hint }}</p>
<p>{{ .Debug }}</p>
</body>
</html>`))

var oidcCompleteHTML = template.Must(template.New("").Parse(`<html>
<body>
<h1>Complete!</h1>
<p>Please close this tab.</p>
</body>
</html>`))

type oidcErrorHTMLParams struct {
	Name        string
	Description string
	Hint        string
	Debug       string
}

func signInWithTesla(ctx context.Context, r driver.OidcRegistry) (*model.Grant, error) {
	serverLocation := fmt.Sprintf("http://localhost:%d", r.TeslaOAuthConfig().BrowserPort)
	redirectURL, err := url.Parse(r.TeslaOAuthConfig().OAuthRedirectURI)
	if err != nil {
		return nil, err
	}

	conf := oauth2.Config{
		ClientID: r.TeslaOAuthConfig().OAuthClientID,
		Endpoint: oauth2.Endpoint{
			TokenURL:  r.TeslaOAuthConfig().OAuthIssuer + "/token/",
			AuthURL:   r.TeslaOAuthConfig().OAuthIssuer + "/authorize",
			AuthStyle: oauth2.AuthStyleInParams,
		},
		RedirectURL: r.TeslaOAuthConfig().OAuthRedirectURI,
		Scopes:      r.TeslaOAuthConfig().OAuthScopes(),
	}
	fmt.Println("ClientID:", r.TeslaOAuthConfig().OAuthClientID)
	var generateAuthCodeURL = func() (string, []rune, string, error) {
		state, err := randx.RuneSequence(24, randx.AlphaLower)
		if err != nil {
			return "", nil, "", err
		}

		nonce, err := randx.RuneSequence(24, randx.AlphaLower)
		if err != nil {
			return "", nil, "", err
		}
		verifier := oauth2.GenerateVerifier()

		authCodeURL := conf.AuthCodeURL(
			string(state),
			oauth2.SetAuthURLParam("audience", strings.Join([]string{r.TeslaAPIConfig().APIHost}, "+")),
			oauth2.SetAuthURLParam("nonce", string(nonce)),
			oauth2.S256ChallengeOption(verifier),
			//oauth2.SetAuthURLParam("prompt", strings.Join(prompt, "+")),
		)
		return authCodeURL, state, verifier, nil
	}
	authCodeURL, state, verifier, err := generateAuthCodeURL()
	if err != nil {
		return nil, err
	}

	_ = webbrowser.Open(serverLocation)

	fmt.Println("Setting up home route on " + serverLocation)
	fmt.Println("Press ctrl + c to end the process.")
	fmt.Printf("If your browser does not open automatically, navigate to:\n\n\t%s\n\n", serverLocation)

	router := http.NewServeMux()
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", r.TeslaOAuthConfig().BrowserPort),
		Handler: router,
	}

	var onDone = func() {
		fmt.Println("Shutting down Sign in with Tesla server...")
		go func() {
			shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := srv.Shutdown(shutdownCtx); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "failed to shut down")
			}
		}()
	}

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_ = oidcWelcomeHTML.Execute(w, &struct{ URL string }{URL: authCodeURL})
	})

	var grant *model.Grant
	router.HandleFunc(redirectURL.Path, func(w http.ResponseWriter, req *http.Request) {
		if len(req.URL.Query().Get("error")) > 0 {
			fmt.Printf("Got error: %s\n", req.URL.Query().Get("error_description"))

			w.WriteHeader(http.StatusInternalServerError)
			_ = oidcErrorHTML.Execute(w, &oidcErrorHTMLParams{
				Name:        req.URL.Query().Get("error"),
				Description: req.URL.Query().Get("error_description"),
				Hint:        req.URL.Query().Get("error_hint"),
				Debug:       req.URL.Query().Get("error_debug"),
			})

			onDone()
			return
		}

		if req.URL.Query().Get("state") != string(state) {
			_, _ = fmt.Fprintf(os.Stderr, "States do not match. Expected %s, got %s\n", string(state), req.URL.Query().Get("state"))

			w.WriteHeader(http.StatusInternalServerError)
			_ = oidcErrorHTML.Execute(w, &oidcErrorHTMLParams{
				Name:        "States do not match",
				Description: "Expected state " + string(state) + " but got " + req.URL.Query().Get("state"),
			})
			onDone()
			return
		}

		code := req.URL.Query().Get("code")
		token, err := conf.Exchange(
			ctx,
			code,
			oauth2.VerifierOption(verifier),
			oauth2.SetAuthURLParam("audience", strings.Join([]string{r.TeslaAPIConfig().APIHost}, "+")),
		)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Unable to exchange code for token: %s\n", err)

			w.WriteHeader(http.StatusInternalServerError)
			_ = oidcErrorHTML.Execute(w, &oidcErrorHTMLParams{
				Name: err.Error(),
			})
			onDone()
			return
		}
		idt := token.Extra("id_token").(string)
		idToken, err := parseIDToken(ctx, r.TeslaOAuthConfig().OAuthIssuer, idt)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Unable to parse ID Token: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			_ = oidcErrorHTML.Execute(w, &oidcErrorHTMLParams{
				Name: err.Error(),
			})
			onDone()
			return
		}

		grant = &model.Grant{
			Subject:      idToken.Subject(),
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
			Scope:        r.TeslaOAuthConfig().OAuthScope,
			Expiry:       token.Expiry,
		}
		fmt.Printf("Subject:\n\t%s\n", grant.Subject)
		//fmt.Printf("Access Token:\n\t%s\n", grant.AccessToken)
		//fmt.Printf("Refresh Token:\n\t%s\n", grant.RefreshToken)
		//fmt.Printf("Expires in:\n\t%s\n", grant.Expiry.Format(time.RFC3339))

		w.WriteHeader(http.StatusOK)
		_ = oidcCompleteHTML.Execute(w, nil)
		time.Sleep(3 * time.Second)

		onDone()
	})

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		_, _ = fmt.Fprintf(os.Stderr, "failed to serve Sign in with Tesla server: %v\n", err)
	} else {
		fmt.Println("shut down server successfully")
	}

	if grant == nil {
		return nil, errors.New("failed to sign in with Tesla")
	}

	return grant, nil
}

// parseIDToken parses the ID Token using the JWK set from the issuer.
// jwks_uri is different between third-party and the other.
// third-party: /discovery/thirdparty/keys
// the other: /discovery/keys
func parseIDToken(ctx context.Context, issuer string, idToken string) (jwt.Token, error) {
	set, err := jwk.Fetch(ctx, issuer+"/discovery/thirdparty/keys")
	if err != nil {
		return nil, err
	}

	return jwt.Parse([]byte(idToken), jwt.WithKeySet(set))
}
