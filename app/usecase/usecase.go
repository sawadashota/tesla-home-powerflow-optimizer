package usecase

import (
	"context"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/driver/configuration"

	"github.com/morikuni/failure/v2"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/model"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/service"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/pkg/restapi"
)

type (
	Usecase interface {
		restapi.StrictServerInterface
	}
	usecase struct {
		r dependencies
	}
	dependencies interface {
		configuration.AppConfigProvider
		service.VehicleServiceProvider
	}
	Provider interface {
		Usecase() Usecase
	}
)

func New(r dependencies) Usecase {
	return &usecase{
		r: r,
	}
}

func (u *usecase) GetVehicleData(ctx context.Context, _ restapi.GetVehicleDataRequestObject) (restapi.GetVehicleDataResponseObject, error) {
	v, err := u.r.VehicleService().GetVehicleData(ctx, u.r.AppConfig().TeslaVIN)
	if err != nil {
		switch failure.CodeOf(err) {
		case model.ErrCodeDeviceOffline:
			return restapi.GetVehicleData400JSONResponse{
				Code:    restapi.FailedPrecondition,
				Message: err.Error(),
			}, nil
		default:
			return restapi.GetVehicleData500JSONResponse{
				Code:    restapi.InternalServerError,
				Message: err.Error(),
			}, nil
		}
	}
	return restapi.GetVehicleData200JSONResponse{
		Vin:   v.VIN,
		State: restapi.VehicleDataState(v.State),
		ChargeState: restapi.ChargeState{
			BatteryLevel:            v.ChargeState.BatteryLevel,
			BatteryRange:            v.ChargeState.BatteryRange,
			ChargeAmps:              v.ChargeState.ChargeAmps,
			ChargeCurrentRequest:    v.ChargeState.ChargeCurrentRequest,
			ChargeCurrentRequestMax: v.ChargeState.ChargeCurrentRequestMax,
			ChargeEnableRequest:     v.ChargeState.ChargeEnableRequest,
			ChargeLimitSoc:          v.ChargeState.ChargeLimitSoc,
			ChargePortDoorOpen:      v.ChargeState.ChargePortDoorOpen,
			ChargePortLatch:         v.ChargeState.ChargingState,
			ChargerActualCurrent:    v.ChargeState.ChargerActualCurrent,
			ChargerVoltage:          v.ChargeState.ChargerVoltage,
			ChargingState:           v.ChargeState.ChargingState,
			MinutesToFullCharge:     v.ChargeState.MinutesToFullCharge,
			TimeToFullCharge:        v.ChargeState.TimeToFullCharge,
			Timestamp:               v.ChargeState.Timestamp.Unix(),
			UsableBatteryLevel:      v.ChargeState.UsableBatteryLevel,
		},
	}, nil
}

func (u *usecase) SettingGetVehicleChargeSetting(ctx context.Context, request restapi.SettingGetVehicleChargeSettingRequestObject) (restapi.SettingGetVehicleChargeSettingResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
