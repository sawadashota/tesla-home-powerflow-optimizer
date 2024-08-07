package model

import "time"

type VehicleState string

const (
	VehicleStateOnline  VehicleState = "online"
	VehicleStateOffline VehicleState = "offline"
)

type VehicleSummary struct {
	VIN   string       `validate:"required"`
	State VehicleState `validate:"required,oneof=online offline"`
}

type VehicleData struct {
	VIN         string       `validate:"required"`
	State       VehicleState `validate:"required,oneof=online offline"`
	ChargeState VehicleChargeState
}

func (v *VehicleData) Validate() error {
	if err := Validate(v); err != nil {
		return err
	}
	return v.ChargeState.Validate()
}

type VehicleChargeState struct {
	VIN                     string        `validate:"required"`
	BatteryLevel            int           `validate:"required,min=0,max=100"`
	BatteryRange            float32       `validate:"required,min=0"`
	ChargeAmps              int           `validate:"required,min=0"`
	ChargeCurrentRequest    int           `validate:"required,min=0"`
	ChargeCurrentRequestMax int           `validate:"required,min=0"`
	ChargeEnableRequest     bool          `validate:"required,min=0"`
	ChargeLimitSoc          int           `validate:"required,min=0,max=100"`
	ChargePortDoorOpen      bool          `validate:"required"`
	ChargePortLatch         string        `validate:"required"`
	ChargerActualCurrent    int           `validate:"required,min=0"`
	ChargerVoltage          int           `validate:"required,min=0"`
	ChargingState           ChargingState `validate:"required"`
	MinutesToFullCharge     int           `validate:"required,min=0"`
	Timestamp               time.Time     `validate:"required"`
	UsableBatteryLevel      int           `validate:"required,min=0,max=100"`
}

func (v *VehicleChargeState) Validate() error {
	return Validate(v)
}

// IsOnline returns true if the vehicle is online.
// Tesla goes sleep mode after 2 minutes of inactivity.
func (v *VehicleChargeState) IsOnline() bool {
	return time.Now().Sub(v.Timestamp) < time.Minute
}

func (v *VehicleChargeState) IsCharging() bool {
	return v.ChargingState == ChargingStateCharging
}
