package aiseg2

import (
	"net/url"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ORIGIN   string `envconfig:"AISEG2_ORIGIN" required:"true"`
	User     string `envconfig:"AISEG2_USER" required:"true"`
	Password string `envconfig:"AISEG2_PASSWORD" required:"true"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c *Config) url(path string) string {
	u, err := url.Parse(c.ORIGIN)
	if err != nil {
		panic(err)
	}
	u.Path = path
	return u.String()
}
