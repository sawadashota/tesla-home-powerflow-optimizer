package repository

import (
	"context"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/model"
)

type (
	VehicleRepository interface {
		// GetVehicle summary
		//
		// It can be used to check if the vehicle is online or offline.
		// If the vehicle is not found, it will return a model.ErrCodeNotFound error.
		// 	You should check if the vehicle is registered in the Tesla account or the VIN is correct.
		GetVehicle(ctx context.Context, vin string) (*model.VehicleSummary, error)

		// GetVehicleData of online vehicle
		//
		// If the vehicle is offline, it will return a model.ErrCodeDeviceOffline error.
		// 	You should wake up the vehicle first before calling this method.
		// If the vehicle is not found, it will return a model.ErrCodeNotFound error.
		// 	You should check if the vehicle is registered in the Tesla account or the VIN is correct.
		GetVehicleData(ctx context.Context, vin string) (*model.VehicleData, error)

		// WakeUp the vehicle from sleep, which is a state to minimize idle energy consumption.
		//
		// If the vehicle is not found, it will return a model.ErrCodeNotFound error.
		// 	You should check if the vehicle is registered in the Tesla account or the VIN is correct.
		WakeUp(ctx context.Context, vin string) error

		// StartCharge the vehicle
		//
		// If the vehicle is not found, it will return a model.ErrCodeNotFound error.
		StartCharge(ctx context.Context, vin string) error

		// StopCharge the vehicle
		//
		// If the vehicle is not found, it will return a model.ErrCodeNotFound error.
		StopCharge(ctx context.Context, vin string) error

		// SetChargeLimit of the vehicle
		//
		// If the vehicle is not found, it will return a model.ErrCodeNotFound error.
		SetChargeLimit(ctx context.Context, vin string, percent int32) error

		// SetChargeAmps of the vehicle
		//
		// If the vehicle is not found, it will return a model.ErrCodeNotFound error.
		SetChargeAmps(ctx context.Context, vin string, amps int32) error
	}
	VehicleRepositoryProvider interface {
		VehicleRepository() VehicleRepository
	}
)
