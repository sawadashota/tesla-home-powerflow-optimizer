package service_test

import (
	"testing"
	"time"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/model"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/service"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/internal/randx"
)

func TestDecideChargingAmps(t *testing.T) {
	newPowerMetric := func(watt int) *model.PowerMetric {
		return &model.PowerMetric{
			Watt:      watt,
			Timestamp: time.Now(),
		}
	}

	const (
		middleAmps = 10
		maxAmps    = 16
	)
	type ChargeStateOption func(*model.VehicleChargeState)
	newChargeState := func(opts ...ChargeStateOption) *model.VehicleChargeState {
		v := &model.VehicleChargeState{
			VIN:                     randx.MustString(17, randx.Alpha),
			BatteryLevel:            73,
			BatteryRange:            200,
			ChargeAmps:              middleAmps,
			ChargeCurrentRequest:    middleAmps,
			ChargeCurrentRequestMax: 16,
			ChargeEnableRequest:     true,
			ChargeLimitSoc:          80,
			ChargePortDoorOpen:      true,
			ChargePortLatch:         "Engaged",
			ChargerActualCurrent:    middleAmps,
			ChargerVoltage:          203,
			ChargingState:           model.ChargingStateCharging,
			MinutesToFullCharge:     79,
			Timestamp:               time.Now(),
			UsableBatteryLevel:      73,
		}

		for _, opt := range opts {
			opt(v)
		}
		return v
	}
	stateStopped := newChargeState(func(state *model.VehicleChargeState) {
		state.ChargeAmps = 0
		state.ChargeCurrentRequest = 0
		state.ChargerActualCurrent = 0
		state.ChargerVoltage = 2
		state.ChargingState = model.ChargingStateStopped
		state.MinutesToFullCharge = 0.0
	})
	stateMiddleAmps := newChargeState()
	changingInProgress := newChargeState(func(state *model.VehicleChargeState) {
		const amps = middleAmps
		state.ChargeAmps = amps - 1
		state.ChargeCurrentRequest = amps
		state.ChargerActualCurrent = amps - 1
	})
	stateMinimumAmps := newChargeState(func(state *model.VehicleChargeState) {
		const amps = service.MinimumChargeAmps
		state.ChargeAmps = amps
		state.ChargeCurrentRequest = amps
		state.ChargerActualCurrent = amps
	})

	setting := &model.ChargeSetting{
		ChargeStartThreshold:        300,
		PowerUsageIncreaseThreshold: 200,
		PowerUsageDecreaseThreshold: -400,
		UpdateInterval:              10 * time.Minute,
	}
	type args struct {
		metric  *model.PowerMetric
		setting *model.ChargeSetting
		state   *model.VehicleChargeState
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "start charging: greater than ChargeStartThreshold",
			args: args{
				metric:  newPowerMetric(setting.ChargeStartThreshold + 1),
				setting: setting,
				state:   stateStopped,
			},
			want: service.MinimumChargeAmps,
		},
		{
			name: "start charging: equal to ChargeStartThreshold",
			args: args{
				metric:  newPowerMetric(setting.ChargeStartThreshold),
				setting: setting,
				state:   stateStopped,
			},
			want: service.MinimumChargeAmps,
		},
		{
			name: "keep stopped: less than ChargeStartThreshold",
			args: args{
				metric:  newPowerMetric(setting.ChargeStartThreshold - 1),
				setting: setting,
				state:   stateStopped,
			},
			want: 0,
		},
		{
			name: "increase an amp: equal to PowerUsageIncreaseThreshold",
			args: args{
				metric:  newPowerMetric(setting.PowerUsageIncreaseThreshold),
				setting: setting,
				state:   stateMiddleAmps,
			},
			want: middleAmps + 1,
		},
		{
			name: "keep an amp: changing in progress",
			args: args{
				metric:  newPowerMetric(setting.PowerUsageIncreaseThreshold),
				setting: setting,
				state:   changingInProgress,
			},
			want: middleAmps,
		},
		{
			name: "increase 3 amps: equal to 3 * PowerUsageIncreaseThreshold",
			args: args{
				metric:  newPowerMetric(setting.PowerUsageIncreaseThreshold * 3),
				setting: setting,
				state:   stateMiddleAmps,
			},
			want: middleAmps + 3,
		},
		{
			name: "set max amps",
			args: args{
				metric:  newPowerMetric(setting.PowerUsageIncreaseThreshold * 20),
				setting: setting,
				state:   stateMiddleAmps,
			},
			want: maxAmps,
		},
		{
			name: "decrease 3 amps: equal to 3 * PowerUsageIncreaseThreshold",
			args: args{
				metric:  newPowerMetric(setting.PowerUsageDecreaseThreshold * 3),
				setting: setting,
				state:   stateMiddleAmps,
			},
			want: middleAmps - 3,
		},
		{
			name: "keep the same amp: less than PowerUsageIncreaseThreshold",
			args: args{
				metric:  newPowerMetric(setting.PowerUsageIncreaseThreshold - 1),
				setting: setting,
				state:   stateMiddleAmps,
			},
			want: middleAmps,
		},
		{
			name: "keep the same amp: more than PowerUsageDecreaseThreshold",
			args: args{
				metric:  newPowerMetric(setting.PowerUsageDecreaseThreshold + 1),
				setting: setting,
				state:   stateMiddleAmps,
			},
			want: middleAmps,
		},
		{
			name: "stop charging: watt is less than ChargeStartThreshold",
			args: args{
				// W = (MinimumChargeAmps * setting.ChargerVoltage) - setting.ChargeStartThreshold= 5 * 203 - 300 = 715
				metric:  newPowerMetric(-716),
				setting: setting,
				state:   stateMinimumAmps,
			},
			want: 0,
		},
		{
			name: "keep the minimum amp: watt is equal to ChargeStartThreshold",
			args: args{
				// W =(MinimumChargeAmps * setting.ChargerVoltage) - setting.ChargeStartThreshold= 5 * 203 - 300 = 715
				metric:  newPowerMetric(-715),
				setting: setting,
				state:   stateMinimumAmps,
			},
			want: service.MinimumChargeAmps,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := service.DecideChargingAmps(tt.args.metric, tt.args.setting, tt.args.state); got != tt.want {
				t.Errorf("DecideChargingAmps() = %v, want %v", got, tt.want)
			}
		})
	}
}
