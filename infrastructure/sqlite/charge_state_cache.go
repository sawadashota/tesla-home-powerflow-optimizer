package sqlite

import (
	"context"

	"github.com/morikuni/failure/v2"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/model"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/repository"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/ent"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/ent/chargestatecache"
)

type chargeStateCache struct {
	client *ent.Client
}

var _ repository.VehicleChargeStateCache = new(chargeStateCache)

func newChargeStateCache(client *ent.Client) *chargeStateCache {
	return &chargeStateCache{
		client: client,
	}
}

func (c *chargeStateCache) FindOne(ctx context.Context, vin string) (*model.VehicleChargeState, error) {
	found, err := c.client.ChargeStateCache.Query().Where(chargestatecache.VinEQ(vin)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, failure.New(model.ErrCodeNotFound, failure.Message("vehicle charge state cache not found"))
		}
		return nil, err
	}
	return &model.VehicleChargeState{
		VIN:                     found.Vin,
		BatteryLevel:            found.BatteryLevel,
		BatteryRange:            found.BatteryRange,
		ChargeAmps:              found.ChargeAmps,
		ChargeCurrentRequest:    found.ChargeCurrentRequest,
		ChargeCurrentRequestMax: found.ChargeCurrentRequestMax,
		ChargeEnableRequest:     found.ChargeEnableRequest,
		ChargeLimitSoc:          found.ChargeLimitSoc,
		ChargePortDoorOpen:      found.ChargePortDoorOpen,
		ChargePortLatch:         found.ChargePortLatch,
		ChargerActualCurrent:    found.ChargerActualCurrent,
		ChargerVoltage:          found.ChargerVoltage,
		ChargingState:           model.ChargingState(found.ChargingState),
		MinutesToFullCharge:     found.MinutesToFullCharge,
		Timestamp:               found.Timestamp,
		UsableBatteryLevel:      found.UsableBatteryLevel,
	}, nil
}

func (c *chargeStateCache) SaveOne(ctx context.Context, state *model.VehicleChargeState) error {
	found, err := c.client.ChargeStateCache.Query().Where(chargestatecache.VinEQ(state.VIN)).First(ctx)
	if found != nil {
		_, err = c.client.ChargeStateCache.Update().
			SetVin(state.VIN).
			SetBatteryLevel(state.BatteryLevel).
			SetBatteryRange(state.BatteryRange).
			SetChargeAmps(state.ChargeAmps).
			SetChargeCurrentRequest(state.ChargeCurrentRequest).
			SetChargeCurrentRequestMax(state.ChargeCurrentRequestMax).
			SetChargeEnableRequest(state.ChargeEnableRequest).
			SetChargeLimitSoc(state.ChargeLimitSoc).
			SetChargePortDoorOpen(state.ChargePortDoorOpen).
			SetChargePortLatch(state.ChargePortLatch).
			SetChargerActualCurrent(state.ChargerActualCurrent).
			SetChargerVoltage(state.ChargerVoltage).
			SetChargingState(string(state.ChargingState)).
			SetMinutesToFullCharge(state.MinutesToFullCharge).
			SetTimestamp(state.Timestamp).
			SetUsableBatteryLevel(state.UsableBatteryLevel).
			Save(ctx)
		return err
	}
	_, err = c.client.ChargeStateCache.Create().
		SetVin(state.VIN).
		SetBatteryLevel(state.BatteryLevel).
		SetBatteryRange(state.BatteryRange).
		SetChargeAmps(state.ChargeAmps).
		SetChargeCurrentRequest(state.ChargeCurrentRequest).
		SetChargeCurrentRequestMax(state.ChargeCurrentRequestMax).
		SetChargeEnableRequest(state.ChargeEnableRequest).
		SetChargeLimitSoc(state.ChargeLimitSoc).
		SetChargePortDoorOpen(state.ChargePortDoorOpen).
		SetChargePortLatch(state.ChargePortLatch).
		SetChargerActualCurrent(state.ChargerActualCurrent).
		SetChargerVoltage(state.ChargerVoltage).
		SetChargingState(string(state.ChargingState)).
		SetMinutesToFullCharge(state.MinutesToFullCharge).
		SetTimestamp(state.Timestamp).
		SetUsableBatteryLevel(state.UsableBatteryLevel).
		Save(ctx)
	return err
}
