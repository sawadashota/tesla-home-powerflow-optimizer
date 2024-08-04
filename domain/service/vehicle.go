package service

import (
	"context"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/model"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/repository"
)

type (
	VehicleService struct {
		r VehicleServiceDependencies
	}
	VehicleServiceDependencies interface {
		repository.VehicleRepositoryProvider
	}
	VehicleServiceProvider interface {
		VehicleService() *VehicleService
	}
)

func NewVehicleService(r VehicleServiceDependencies) *VehicleService {
	return &VehicleService{r: r}
}

// GetVehicleData of online vehicle
// if the vehicle is offline, it will return a model.ErrCodeDeviceOffline error
func (s *VehicleService) GetVehicleData(ctx context.Context, vin string) (*model.VehicleData, error) {
	return s.r.VehicleRepository().GetVehicleData(ctx, vin)
}
