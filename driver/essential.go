package driver

import (
	"context"
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
		Context() context.Context
		configuration.Provider
		logx.Provider
		httpclient.Provider

		repository.Provider
	}
	essentialRegistry struct {
		ctx context.Context
		configurationProvider
		logger     *slog.Logger
		httpclient http.Client

		repositoryProvider
	}
)

var _ EssentialRegistry = new(essentialRegistry)

func NewEssentialRegistry(ctx context.Context) (EssentialRegistry, error) {
	config, err := configuration.New()
	if err != nil {
		return nil, err
	}
	repo, err := sqlite.NewProvider(config)
	return &essentialRegistry{
		ctx:                   ctx,
		configurationProvider: config,
		repositoryProvider:    repo,
	}, nil
}

func (r *essentialRegistry) Context() context.Context {
	return r.ctx
}

func (r *essentialRegistry) Logger() *slog.Logger {
	if r.logger == nil {
		r.logger = logx.New(r.AppConfig().LogLevel)
	}
	return r.logger
}

func (r *essentialRegistry) HTTPClient() *http.Client {
	return http.DefaultClient
}
