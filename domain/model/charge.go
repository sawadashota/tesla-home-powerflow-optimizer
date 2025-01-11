package model

import (
	"strings"
	"time"
)

type ChargingState string

const (
	ChargingStateUnknown  ChargingState = "Unknown"
	ChargingStateCharging ChargingState = "Charging"
	ChargingStateStopped  ChargingState = "Stopped"
)

type ChargeSetting struct {
	Enabled                     bool                 `validate:"required"`
	ChargeStartThreshold        int                  `validate:"required,min=100"`
	PowerUsageIncreaseThreshold int                  `validate:"required,min=100"`
	PowerUsageDecreaseThreshold int                  `validate:"required,max=-100"`
	UpdateInterval              time.Duration        `validate:"required,min=10m"`
	MinCharge                   MinimumChargeSetting `validate:"required"`
}

var ChargeSettingDefault = &ChargeSetting{
	ChargeStartThreshold:        300,
	PowerUsageIncreaseThreshold: 300,
	PowerUsageDecreaseThreshold: -400,
	UpdateInterval:              30 * time.Minute,

	MinCharge: MinimumChargeSetting{
		Threshold: 0,
		TimeRange: TimeRange{
			Start: NewTimeOnly(22, 0, 0),
			End:   NewTimeOnly(6, 0, 0),
		},
		Amperage: 16,
	},
}

func (s *ChargeSetting) Validate() error {
	return Validate(s)
}

type MinimumChargeSetting struct {
	Threshold int       `validate:"required,min=0,max=100"`
	TimeRange TimeRange `validate:"required"`
	Amperage  int       `validate:"required,min=5,max=16"`
}

func (s *MinimumChargeSetting) Enabled() bool {
	return s.Threshold > 0
}

type TimeOnly time.Time

func TimeOnlyNow() TimeOnly {
	return NewTimeOnlyFromTime(time.Now())
}

func NewTimeOnlyFromTime(t time.Time) TimeOnly {
	return NewTimeOnly(t.Hour(), t.Minute(), t.Second())
}

func NewTimeOnly(hour, minute, second int) TimeOnly {
	t := time.Date(1970, 1, 1, hour, minute, second, 0, time.Local)
	return TimeOnly(t)
}

func ParseTimeOnly(s string) (TimeOnly, error) {
	layout := "15:04:05"
	if len(strings.Split(s, ":")) == 2 {
		layout = "15:04"
	}
	t, err := time.Parse(layout, s)
	if err != nil {
		return TimeOnly{}, err
	}
	return NewTimeOnlyFromTime(t), nil
}

func (t TimeOnly) String() string {
	return time.Time(t).Format("15:04:05")
}

func (t TimeOnly) HourMinute() string {
	return time.Time(t).Format("15:04")
}

func (t TimeOnly) Equal(u TimeOnly) bool {
	return time.Time(t).Equal(time.Time(u))
}

func (t TimeOnly) After(u TimeOnly) bool {
	return time.Time(t).After(time.Time(u))
}

func (t TimeOnly) AfterOrEqual(u TimeOnly) bool {
	return t.After(u) || t.Equal(u)
}

func (t TimeOnly) Before(u TimeOnly) bool {
	return time.Time(t).Before(time.Time(u))
}

func (t TimeOnly) BeforeOrEqual(u TimeOnly) bool {
	return t.Before(u) || t.Equal(u)
}

type TimeRange struct {
	Start TimeOnly `validate:"required"`
	End   TimeOnly `validate:"required"`
}

// InRange returns true if the given time is within the range.
// ex. start: 21:00, end: 6:00, t: 23:00 => true
// ex. start: 21:00, end: 6:00, t: 7:00 => false
// ex. start: 6:00, end: 21:00, t: 23:00 => false
func (r *TimeRange) InRange(t TimeOnly) bool {
	if r.Start.Before(r.End) {
		return t.AfterOrEqual(r.Start) && t.BeforeOrEqual(r.End)
	}
	return t.AfterOrEqual(r.Start) || t.BeforeOrEqual(r.End)
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
