package configuration

import (
	"strings"

	"github.com/kelseyhightower/envconfig"
	"github.com/teslamotors/vehicle-command/pkg/protocol"
	"golang.org/x/oauth2"
)

type (
	TeslaOAuthConfig struct {
		BrowserPort       int    `envconfig:"SIGNIN_WITH_TESLA_BROWSER_PORT" default:"3090"`
		OAuthIssuer       string `envconfig:"TESLA_OAUTH_ISSUER" default:"https://auth.tesla.com/oauth2/v3"`
		OAuthClientID     string `envconfig:"TESLA_OAUTH_CLIENT_ID" required:"true"`
		OAuthClientSecret string `envconfig:"TESLA_OAUTH_CLIENT_SECRET" required:"true"`
		OAuthRedirectURI  string `envconfig:"TESLA_OAUTH_REDIRECT_URI" default:"http://localhost:3090/oauth/signin/callback"`
		OAuthScope        string `envconfig:"TESLA_OAUTH_SCOPE" default:"openid offline_access vehicle_device_data vehicle_charging_cmds"`
		PrivateKeyPath    string `envconfig:"TESLA_PRIVATE_KEY_PATH" default:"private_key.pem"`

		APIHost string `envconfig:"TESLA_API_HOST" default:"https://fleet-api.prd.na.vn.cloud.tesla.com"`
	}
	TeslaOAuthConfigProvider interface {
		TeslaOAuthConfig() *TeslaOAuthConfig
	}
	teslaOAuthConfigProvider struct {
		config *TeslaOAuthConfig
	}
)

func NewTeslaOAuthConfig() (*TeslaOAuthConfig, error) {
	c := new(TeslaOAuthConfig)
	if err := envconfig.Process("", c); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *TeslaOAuthConfig) Config() oauth2.Config {
	return oauth2.Config{
		ClientID:     c.OAuthClientID,
		ClientSecret: c.OAuthClientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL:  c.OAuthIssuer + "/token/",
			AuthURL:   c.OAuthIssuer + "/authorize",
			AuthStyle: oauth2.AuthStyleInParams,
		},
		RedirectURL: c.OAuthRedirectURI,
		Scopes:      c.OAuthScopes(),
	}
}

func NewTeslaOAuthConfigProvider(config *TeslaOAuthConfig) TeslaOAuthConfigProvider {
	return &teslaOAuthConfigProvider{config: config}
}

func (c *TeslaOAuthConfig) OAuthScopes() []string {
	return strings.Split(c.OAuthScope, " ")
}

func (p *teslaOAuthConfigProvider) TeslaOAuthConfig() *TeslaOAuthConfig {
	return p.config
}

func (c *TeslaOAuthConfig) PrivateKey() protocol.ECDHPrivateKey {
	key, err := protocol.LoadPrivateKey(c.PrivateKeyPath)
	if err != nil {
		panic(err)
	}
	return key
}

func (c *TeslaOAuthConfig) Audience() string {
	return c.APIHost
}
