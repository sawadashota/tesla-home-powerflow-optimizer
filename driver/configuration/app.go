package configuration

import "github.com/kelseyhightower/envconfig"

type (
	AppConfig struct {
		LogLevel  string `envconfig:"LOG_LEVEL" default:"INFO"`
		TeslaVIN  string `envconfig:"TESLA_VIN" required:"true"`
		SqliteDSN string `envconfig:"SQLITE_DSN" default:"file:sqlite.db?cache=shared&_fk=1"`
	}
	AppConfigProvider interface {
		AppConfig() *AppConfig
	}
	appConfigProvider struct {
		config *AppConfig
	}
)

func NewAppConfig() (*AppConfig, error) {
	c := new(AppConfig)
	if err := envconfig.Process("", c); err != nil {
		return nil, err
	}
	return c, nil
}

func NewAppConfigProvider(config *AppConfig) AppConfigProvider {
	return &appConfigProvider{config: config}
}

func (p *appConfigProvider) AppConfig() *AppConfig {
	return p.config
}
