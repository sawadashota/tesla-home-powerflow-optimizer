package configuration

import "github.com/kelseyhightower/envconfig"

type (
	ServerConfig struct {
		Port int `envconfig:"SERVER_PORT" default:"8088"`
	}
	ServerConfigProvider interface {
		ServerConfig() *ServerConfig
	}
	serverConfigProvider struct {
		config *ServerConfig
	}
)

func NewServerConfig() (*ServerConfig, error) {
	c := new(ServerConfig)
	if err := envconfig.Process("", c); err != nil {
		return nil, err
	}
	return c, nil
}

func NewServerConfigProvider(config *ServerConfig) ServerConfigProvider {
	return &serverConfigProvider{config: config}
}

func (p *serverConfigProvider) ServerConfig() *ServerConfig {
	return p.config
}
