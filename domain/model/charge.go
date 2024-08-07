package model

import "time"

type ChargingState string

const (
	ChargingStateUnknown  ChargingState = "Unknown"
	ChargingStateCharging ChargingState = "Charging"
	ChargingStateStopped  ChargingState = "Stopped"
)

type ChargeSetting struct {
	Enabled                     bool          `validate:"required"`
	ChargeStartThreshold        int           `validate:"required,min=100"`
	PowerUsageIncreaseThreshold int           `validate:"required,min=100"`
	PowerUsageDecreaseThreshold int           `validate:"required,max=-100"`
	UpdateInterval              time.Duration `validate:"required,min=10m"`
}

var ChargeSettingDefault = &ChargeSetting{
	ChargeStartThreshold:        300,
	PowerUsageIncreaseThreshold: 300,
	PowerUsageDecreaseThreshold: -400,
	UpdateInterval:              30 * time.Minute,
}

func (s *ChargeSetting) Validate() error {
	return Validate(s)
}

type ChargeCommandOperation string

const (
	ChargeCommandOperationStart ChargeCommandOperation = "start"
	ChargeCommandOperationStop  ChargeCommandOperation = "stop"
	ChargeCommandOperationSet   ChargeCommandOperation = "set"
)

type ChargeCommandHistory struct {
	VIN       string                 `validate:"required"`
	Operation ChargeCommandOperation `validate:"required,oneof=start stop set"`
	Amps      int                    `validate:"required,min=0"`
	Timestamp time.Time
}

func (h *ChargeCommandHistory) Validate() error {
	return Validate(h)
}
