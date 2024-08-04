package tesla

import (
	"context"
	"log/slog"

	"golang.org/x/oauth2"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/model"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/repository"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/driver/configuration"
)

type (
	tokenSource struct {
		ctx     context.Context
		conf    oauth2.Config
		cache   *oauth2.Token
		subject string
		scope   string
		r       tokenSourceDependencies
	}
	tokenSourceDependencies interface {
		configuration.TeslaOAuthConfigProvider
		repository.GrantRepositoryProvider
	}
)

func newTokenSource(ctx context.Context, r tokenSourceDependencies) oauth2.TokenSource {
	return &tokenSource{
		ctx:  ctx,
		conf: r.TeslaOAuthConfig().Config(),
		r:    r,
	}
}

func (t *tokenSource) Token() (*oauth2.Token, error) {
	token, err := t.getToken(t.ctx)
	if err != nil {
		return nil, err
	}
	if token.Valid() {
		slog.Info("using stored token")
		return token, nil
	}

	slog.Info("refreshing token...")
	return t.refreshToken(t.ctx, token)
}

func (t *tokenSource) getToken(ctx context.Context) (*oauth2.Token, error) {
	if t.cache != nil {
		return t.cache, nil
	}

	found, err := t.r.GrantRepository().FindLatestOne(ctx)
	if err != nil {
		return nil, err
	}
	t.subject = found.Subject
	t.scope = found.Scope
	return &oauth2.Token{
		AccessToken:  found.AccessToken,
		TokenType:    "Bearer",
		RefreshToken: found.RefreshToken,
		Expiry:       found.Expiry,
	}, nil
}

func (t *tokenSource) refreshToken(ctx context.Context, token *oauth2.Token) (*oauth2.Token, error) {
	newToken, err := t.conf.TokenSource(ctx, token).Token()
	if err != nil {
		return nil, err
	}

	newGrant := &model.Grant{
		Subject:      t.subject,
		AccessToken:  newToken.AccessToken,
		RefreshToken: newToken.RefreshToken,
		Scope:        t.scope,
		Expiry:       newToken.Expiry,
	}
	if err := t.r.GrantRepository().SaveOne(ctx, newGrant); err != nil {
		return nil, err
	}
	return newToken, nil
}
