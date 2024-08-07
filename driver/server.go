package driver

import (
	"context"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/app/usecase"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/collector"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/collector/aiseg2"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/event"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/repository"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/service"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/driver/configuration"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/infrastructure/tesla"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/interfaces/worker"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/interfaces/worker/chargecontroller"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/internal/logx"
)

type (
	ServerRegistry interface {
		logx.Provider
		configuration.ServerConfigProvider
		configuration.AppConfigProvider

		usecase.Provider

		service.VehicleServiceProvider
		service.ChargeServiceProvider

		repository.MigrationProvider
		repository.VehicleRepositoryProvider
		repository.PowerMetricRepositoryProvider
		repository.ChargeSettingRepositoryProvider

		collector.Provider
		worker.PubSubProvider
	}
	serverRegistry struct {
		EssentialRegistry

		usecase usecase.Usecase

		vehicleSvc *service.VehicleService
		chargeSvc  *service.ChargeService

		vehicleRepository repository.VehicleRepository

		collector collector.Collector
		worker.PubSubProvider
		chargeControllerSubscriber chargecontroller.Subscriber
	}
)

var _ ServerRegistry = new(serverRegistry)

func NewServerRegistry(ctx context.Context) (ServerRegistry, error) {
	registry, err := NewEssentialRegistry(ctx)
	if err != nil {
		return nil, err
	}
	r := &serverRegistry{
		EssentialRegistry: registry,
	}
	r.PubSubProvider = worker.NewPubSub(r)
	return r, nil
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

func (r *serverRegistry) ChargeService() *service.ChargeService {
	if r.chargeSvc == nil {
		r.chargeSvc = service.NewChargeService(r)
	}
	return r.chargeSvc
}

func (r *serverRegistry) Usecase() usecase.Usecase {
	if r.usecase == nil {
		r.usecase = usecase.New(r)
	}
	return r.usecase
}

func (r *serverRegistry) Collector() collector.Collector {
	if r.collector == nil {
		switch r.AppConfig().Collector {
		case "aiseg2":
			c, err := aiseg2.NewClient()
			if err != nil {
				panic(err)
			}
			r.collector = c
		default:
			panic("unknown collector")
		}
	}
	return r.collector
}

func (r *serverRegistry) MetricInsertedEventSubscriber() event.MetricInsertedEventSubscriber {
	if r.chargeControllerSubscriber == nil {
		r.chargeControllerSubscriber = chargecontroller.New(r)
	}
	return r.chargeControllerSubscriber
}
