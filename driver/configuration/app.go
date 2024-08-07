package configuration

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type (
	AppConfig struct {
		LogLevel          string `envconfig:"LOG_LEVEL" default:"INFO"`
		TeslaVIN          string `envconfig:"TESLA_VIN" required:"true"`
		SqliteDSN         string `envconfig:"SQLITE_DSN" default:"file:sqlite.db?cache=shared&_fk=1"`
		Collector         string `envconfig:"COLLECTOR" default:"aiseg2"`
		CollectorInterval int    `envconfig:"COLLECTOR_INTERVAL" default:"5"`
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

func (c *AppConfig) CollectorIntervalDuration() time.Duration {
	return time.Duration(c.CollectorInterval) * time.Minute
}

func NewAppConfigProvider(config *AppConfig) AppConfigProvider {
	return &appConfigProvider{config: config}
}

func (p *appConfigProvider) AppConfig() *AppConfig {
	return p.config
}
