package configuration

import (
	"github.com/kelseyhightower/envconfig"
	"strings"
)

type (
	TeslaOAuthConfig struct {
		BrowserPort      int    `envconfig:"SIGNIN_WITH_TESLA_BROWSER_PORT" default:"3090"`
		OAuthIssuer      string `envconfig:"TESLA_OAUTH_ISSUER" default:"https://auth.tesla.com/oauth2/v3"`
		OAuthClientID    string `envconfig:"TESLA_OAUTH_CLIENT_ID" required:"true"`
		OAuthRedirectURI string `envconfig:"TESLA_OAUTH_REDIRECT_URI" default:"http://localhost:3090/oauth/signin/callback"`
		OAuthScope       string `envconfig:"TESLA_OAUTH_SCOPE" default:"openid offline_access vehicle_device_data vehicle_charging_cmds"`
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

func NewTeslaOAuthConfigProvider(config *TeslaOAuthConfig) TeslaOAuthConfigProvider {
	return &teslaOAuthConfigProvider{config: config}
}

func (p *teslaOAuthConfigProvider) TeslaOAuthConfig() *TeslaOAuthConfig {
	return p.config
}

type (
	TeslaAPIConfig struct {
		APIHost string `envconfig:"TESLA_API_HOST" default:"https://fleet-api.prd.na.vn.cloud.tesla.com"`
	}
	TeslaAPIConfigProvider interface {
		TeslaAPIConfig() *TeslaAPIConfig
	}
	teslaAPIConfigProvider struct {
		config *TeslaAPIConfig
	}
)

func NewTeslaAPIConfig() (*TeslaAPIConfig, error) {
	c := new(TeslaAPIConfig)
	if err := envconfig.Process("", c); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *TeslaOAuthConfig) OAuthScopes() []string {
	return strings.Split(c.OAuthScope, " ")
}

func NewTeslaAPIConfigProvider(config *TeslaAPIConfig) TeslaAPIConfigProvider {
	return &teslaAPIConfigProvider{config: config}
}

func (p *teslaAPIConfigProvider) TeslaAPIConfig() *TeslaAPIConfig {
	return p.config
}
