package driver

import (
	"log/slog"
	"net/http"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/repository"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/driver/configuration"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/infrastructure/sqlite"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/internal/httpclient"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/internal/logx"
)

type (
	EssentialRegistry interface {
		configuration.Provider
		logx.Provider
		httpclient.Provider

		repository.Provider
	}
	essentialRegistry struct {
		config     configuration.Configuration
		logger     *slog.Logger
		httpclient http.Client

		repository.Provider
	}
)

var _ EssentialRegistry = new(essentialRegistry)

func NewEssentialRegistry() (EssentialRegistry, error) {
	config, err := configuration.New()
	if err != nil {
		return nil, err
	}
	repo, err := sqlite.NewProvider(config)
	return &essentialRegistry{
		config:   config,
		Provider: repo,
	}, nil
}

func (r *essentialRegistry) Configuration() configuration.Configuration {
	return r.config
}

func (r *essentialRegistry) Logger() *slog.Logger {
	if r.logger == nil {
		r.logger = logx.New(r.Configuration().AppConfig().LogLevel)
	}
	return r.logger
}

func (r *essentialRegistry) HTTPClient() *http.Client {
	return http.DefaultClient
}
