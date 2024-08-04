package driver

import (
	"context"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/driver/configuration"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/app/usecase"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/service"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/repository"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/infrastructure/tesla"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/internal/logx"
)

type (
	ServerRegistry interface {
		logx.Provider
		configuration.ServerConfigProvider
		configuration.AppConfigProvider

		usecase.Provider

		service.VehicleServiceProvider
	}
	serverRegistry struct {
		EssentialRegistry

		usecase usecase.Usecase

		vehicleSvc *service.VehicleService

		vehicleRepository repository.VehicleRepository
	}
)

var _ ServerRegistry = new(serverRegistry)

func NewServerRegistry(ctx context.Context) (ServerRegistry, error) {
	registry, err := NewEssentialRegistry(ctx)
	if err != nil {
		return nil, err
	}
	return &serverRegistry{
		EssentialRegistry: registry,
	}, nil
}

func (r *serverRegistry) VehicleRepository() repository.VehicleRepository {
	if r.vehicleRepository == nil {
		r.vehicleRepository = tesla.NewClient(r.Context(), r)
	}
	return r.vehicleRepository
}

func (r *serverRegistry) VehicleService() *service.VehicleService {
	if r.vehicleSvc == nil {
		r.vehicleSvc = service.NewVehicleService(r)
	}
	return r.vehicleSvc
}

func (r *serverRegistry) Usecase() usecase.Usecase {
	if r.usecase == nil {
		r.usecase = usecase.New(r)
	}
	return r.usecase
}
