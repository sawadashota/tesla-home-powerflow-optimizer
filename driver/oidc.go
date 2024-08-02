package driver

import (
	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/repository"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/service"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/driver/configuration"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/internal/logx"
)

type (
	OidcRegistry interface {
		logx.Provider
		configuration.TeslaOAuthConfigProvider
		configuration.TeslaAPIConfigProvider
		repository.MigrationProvider
		service.GrantServiceProvider
	}
	oidcRegistry struct {
		EssentialRegistry
		grantService *service.GrantService
	}
)

var _ OidcRegistry = new(oidcRegistry)

func NewOidcRegistry() (OidcRegistry, error) {
	registry, err := NewEssentialRegistry()
	if err != nil {
		return nil, err
	}
	return &oidcRegistry{
		EssentialRegistry: registry,
	}, nil
}

func (r *oidcRegistry) TeslaOAuthConfig() *configuration.TeslaOAuthConfig {
	return r.Configuration().TeslaOAuthConfig()
}

func (r *oidcRegistry) TeslaAPIConfig() *configuration.TeslaAPIConfig {
	return r.Configuration().TeslaAPIConfig()
}

func (r *oidcRegistry) GrantService() *service.GrantService {
	if r.grantService == nil {
		r.grantService = service.NewGrantService(r)
	}
	return r.grantService
}
