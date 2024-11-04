package service

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/morikuni/failure/v2"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/model"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/repository"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/driver/configuration"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/internal/logx"
)

type (
	ChargeService struct {
		r ChargeServiceDependencies
	}
	ChargeServiceDependencies interface {
		logx.Provider
		configuration.AppConfigProvider
		repository.PowerMetricRepositoryProvider
		repository.ChargeSettingRepositoryProvider
		repository.VehicleChargeStateCacheProvider
		repository.VehicleRepositoryProvider
		repository.ChargeCommandHistoryRepositoryProvider
	}
	ChargeServiceProvider interface {
		ChargeService() *ChargeService
	}
)

func NewChargeService(r ChargeServiceDependencies) *ChargeService {
	return &ChargeService{r: r}
}

func (s *ChargeService) RefreshChargeState(ctx context.Context, vin string) (*model.VehicleChargeState, error) {
	state, err := s.r.VehicleChargeStateCache().FindOne(ctx, vin)
	if err == nil && state.IsFresh() {
		s.r.Logger().Info("cache is fresh", slog.String("vin", vin))
		return state, nil
	} else if err != nil && !failure.Is(err, model.ErrCodeNotFound) {
		return nil, err
	}

	s.r.Logger().Info("fetching state from the vehicle...", slog.String("vin", vin))

	if err := s.r.VehicleRepository().WakeUp(ctx, vin); err != nil {
		return nil, err
	}
	return s.waitUntilWakedUp(ctx, vin)
}

// Adjust evaluates the current power usage, charge settings, and latest charge status cache,
// and makes necessary adjustments to the charging process. This may involve starting
// or stopping the charging process, as well as increasing or decreasing the charging amperage.
func (s *ChargeService) Adjust(ctx context.Context) error {
	setting, err := s.r.ChargeSettingRepository().FindOne(ctx)
	if err != nil {
		return err
	}
	if !setting.Enabled {
		s.r.Logger().Info("charge setting is disabled")
		return nil
	}

	const metricCount = 2
	metrics, err := s.r.PowerMetricRepository().FindLatestN(ctx, metricCount)
	if err != nil {
		return err
	}
	if len(metrics) < metricCount ||
		metrics.MaximumInterval() > s.r.AppConfig().CollectorIntervalDuration()*2 ||
		metrics.LatestTimestamp().Add(s.r.AppConfig().CollectorIntervalDuration()*2).Before(time.Now()) {
		s.r.Logger().Info("invalid metrics. not enough data, too old or too long interval",
			slog.Int("count", len(metrics)),
			slog.Time("latest_timestamp", metrics.LatestTimestamp()),
			slog.Duration("interval", metrics.MaximumInterval()),
		)
		return nil
	}

	// check last charge command history
	history, err := s.r.ChargeCommandHistoryRepository().FindLatestOne(ctx)
	if err != nil {
		if !failure.Is(err, model.ErrCodeNotFound) {
			return err
		}
	} else {
		if time.Since(history.Timestamp) < setting.UpdateInterval {
			s.r.Logger().Info(
				"last charge command is too recent. skip at the moment.",
				slog.Time("last_timestamp", history.Timestamp),
				slog.String("interval", time.Since(history.Timestamp).String()),
				slog.String("interval_remains", (setting.UpdateInterval-time.Since(history.Timestamp)).String()),
			)
			return nil
		}
	}

	state, err := s.describeChargeState(ctx, s.r.AppConfig().TeslaVIN)
	if err != nil {
		return err
	}

	// check if the decision is stable
	var decision int
	decisions := make(map[int]struct{}, metricCount)
	for _, metric := range metrics {
		decision = DecideChargingAmps(metric, setting, state)
		decisions[decision] = struct{}{}
	}
	if len(decisions) > 1 {
		// skip at the moment. wait for decision stabilization
		arr := make([]int, 0, len(decisions))
		for d := range decisions {
			arr = append(arr, d)
		}
		s.r.Logger().Info("decision is not stable. skip at the moment.", slog.String("decisions", fmt.Sprint(arr)))
		return nil
	}

	// adjust the charge state
	if state.IsCharging() {
		if state.ChargeAmps == decision {
			s.r.Logger().Info("no need to adjust because decision is not changed", slog.Int("decision", decision), slog.Int("current", state.ChargeAmps))
			return nil
		}
	} else {
		if decision == 0 {
			s.r.Logger().Info("no need to adjust because already stopped", slog.Int("decision", decision))
			return nil
		}
	}

	if err := s.updateChargeState(ctx, state.VIN, state, decision); err != nil {
		return err
	}

	_, err = s.describeVehicleChargeStateWithDelay(ctx, state.VIN)
	return err
}

func (s *ChargeService) DescribeSetting(ctx context.Context) (*model.ChargeSetting, error) {
	return s.r.ChargeSettingRepository().FindOne(ctx)
}

func (s *ChargeService) SaveSetting(ctx context.Context, setting *model.ChargeSetting) error {
	return s.r.ChargeSettingRepository().SaveOne(ctx, setting)
}

func (s *ChargeService) updateChargeState(ctx context.Context, vin string, state *model.VehicleChargeState, amps int) error {
	if !state.IsFresh() {
		if err := s.r.VehicleRepository().WakeUp(ctx, vin); err != nil {
			return err
		}

		newState, err := s.waitUntilWakedUp(ctx, vin)
		if err != nil {
			return err
		}
		state = newState
	}

	if state.ChargePortLatch != "Engaged" {
		s.r.Logger().Info("charge port is not engaged", slog.String("vin", vin))
		return s.r.ChargeSettingRepository().SetEnabled(ctx, false)
	}
	if state.ChargeLimitSoc == state.BatteryLevel && !state.IsCharging() {
		s.r.Logger().Info("battery level is full", slog.String("vin", vin), slog.Int("battery_level", state.BatteryLevel))
		return s.r.ChargeSettingRepository().SetEnabled(ctx, false)
	}
	if amps == 0 {
		if !state.IsCharging() {
			s.r.Logger().Info("already stopped charging. nothing to do", slog.String("vin", vin))
			return nil
		}
		if err := s.r.VehicleRepository().StopCharge(ctx, vin); err != nil {
			return err
		}
		return s.r.ChargeCommandHistoryRepository().CreateOne(ctx, &model.ChargeCommandHistory{
			VIN:       state.VIN,
			Operation: model.ChargeCommandOperationStop,
			Amps:      0,
			Timestamp: time.Now(),
		})
	}

	history := &model.ChargeCommandHistory{
		VIN:       state.VIN,
		Operation: model.ChargeCommandOperationSet,
		Amps:      amps,
		Timestamp: time.Now(),
	}
	if !state.IsCharging() {
		if err := s.r.VehicleRepository().StartCharge(ctx, vin); err != nil {
			return err
		}
		history.Operation = model.ChargeCommandOperationStart
	}
	if state.ChargeCurrentRequest != amps {
		if err := s.r.VehicleRepository().SetChargeAmps(ctx, vin, int32(amps)); err != nil {
			return err
		}
	}
	return s.r.ChargeCommandHistoryRepository().CreateOne(ctx, history)
}

func (s *ChargeService) describeChargeState(ctx context.Context, vin string) (*model.VehicleChargeState, error) {
	state, err := s.r.VehicleChargeStateCache().FindOne(ctx, vin)
	if err == nil {
		return state, nil
	}
	if !failure.Is(err, model.ErrCodeNotFound) {
		return nil, err
	}
	s.r.Logger().Info("cache not found. fetch from the vehicle", slog.String("vin", vin))

	if err := s.r.VehicleRepository().WakeUp(ctx, vin); err != nil {
		return nil, err
	}
	return s.waitUntilWakedUp(ctx, vin)
}

func (s *ChargeService) waitUntilWakedUp(ctx context.Context, vin string) (*model.VehicleChargeState, error) {
	vehicle, err := s.r.VehicleRepository().WaitUntilWakedUp(ctx, vin)
	if err != nil {
		return nil, err
	}
	if err := s.r.VehicleChargeStateCache().SaveOne(ctx, &vehicle.ChargeState); err != nil {
		return nil, err
	}
	return &vehicle.ChargeState, nil
}

// describeVehicleChargeStateWithDelay describes the vehicle charge state with a delay due to update.
func (s *ChargeService) describeVehicleChargeStateWithDelay(ctx context.Context, vin string) (*model.VehicleChargeState, error) {
	time.Sleep(20 * time.Second)
	vehicle, err := s.r.VehicleRepository().GetVehicleData(ctx, vin)
	if err != nil {
		return nil, err
	}
	if err := s.r.VehicleChargeStateCache().SaveOne(ctx, &vehicle.ChargeState); err != nil {
		return nil, err
	}
	return &vehicle.ChargeState, nil
}

const MinimumChargeAmps = 5

// DecideChargingAmps decides the charging amps based on the current power metric, charge setting, and vehicle charge state.
// 充電中の場合:
//   - 余剰電力が正の値で、ChargeCurrentRequest が ChargeCurrentRequestMax に達している場合は何もしない
//   - 余剰電力が PowerUsageIncreaseThreshold より大きい場合、アンペアを増やす
//     A = max(ChargeAmps + (余剰電力 + PowerUsageDecreaseThreshold) / ChargerVoltage, ChargeCurrentRequestMax)
//   - 余剰電力が PowerUsageDecreaseThreshold より小さい場合、
//     1. まずはアンペアを減らした結果が PowerUsageDecreaseThreshold より大きいかどうかを確認する
//     A = ChargeAmps - (PowerUsageDecreaseThreshold - 余剰電力) / ChargerVoltage
//     2. 1の結果が 最小アンペア数 (5A) より小さい場合、最小アンペア数に設定したときの余剰電力が ChargeStopThreshold より大きいかどうかを確認する
//     condition: 5 * ChargerVoltage > ChargeStartThreshold
//     上記の結果が false の場合、充電を停止する。 true の場合、最小アンペア数を返す
//
// 充電中でない場合:
//   - 余剰電力が ChargeStartThreshold より大きい場合、充電を開始する
func DecideChargingAmps(metric *model.PowerMetric, setting *model.ChargeSetting, state *model.VehicleChargeState) int {
	if !state.IsCharging() {
		if metric.Watt >= setting.ChargeStartThreshold {
			return MinimumChargeAmps
		}
		return 0
	}
	if state.ChargeCurrentRequest != state.ChargeAmps {
		return state.ChargeCurrentRequest
	}
	if metric.Watt >= setting.PowerUsageIncreaseThreshold {
		return min(state.ChargeAmps+(metric.Watt+setting.PowerUsageIncreaseThreshold)/state.ChargerVoltage, state.ChargeCurrentRequestMax)
	}
	if metric.Watt <= setting.PowerUsageDecreaseThreshold {
		amps := state.ChargeAmps - (setting.PowerUsageDecreaseThreshold-metric.Watt)/state.ChargerVoltage
		if amps < MinimumChargeAmps {
			if MinimumChargeAmps*state.ChargerVoltage+metric.Watt < setting.ChargeStartThreshold {
				return 0
			}
		}
		return max(amps, MinimumChargeAmps)
	}
	return state.ChargeAmps
}
