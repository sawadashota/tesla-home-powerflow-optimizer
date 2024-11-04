package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/collector"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/repository"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/service"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/driver/configuration"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/internal/logx"
)

type (
	Worker struct {
		r dependencies
	}
	dependencies interface {
		logx.Provider
		configuration.AppConfigProvider
		collector.Provider
		service.ChargeServiceProvider
		repository.PowerMetricRepositoryProvider
	}
	Provider interface {
		Worker() *Worker
	}
)

func New(r dependencies) *Worker {
	return &Worker{r: r}
}

func (w *Worker) RunSurplusPowerCollector(ctx context.Context) error {
	interval := w.r.AppConfig().CollectorIntervalDuration()
	w.r.Logger().InfoContext(ctx, fmt.Sprintf("collector interval: %s", interval))
	ticker := time.NewTicker(interval)
	perform := func() {
		w.r.Logger().Info("collecting surplus power...")
		if err := w.collectMetricAndAdjustCharging(ctx); err != nil {
			w.r.Logger().Error("failed to collect", logx.ErrorAttr(err))
		}

		const metricRetention = 100
		olderThan := time.Now().Add(-w.r.AppConfig().CollectorIntervalDuration() * metricRetention)
		err := w.r.PowerMetricRepository().DeleteOlderThan(ctx, olderThan)
		if err != nil {
			w.r.Logger().Error("failed to delete old power metrics", logx.ErrorAttr(err))
		}
	}

	perform()
	for {
		select {
		case <-ctx.Done():
			w.r.Logger().Info("context done")
			return nil
		case <-ticker.C:
			perform()
		}
	}
}

func (w *Worker) RunPlugInWatcher(ctx context.Context) error {
	interval := w.r.AppConfig().EVPowerWatcherIntervalDuration()
	w.r.Logger().InfoContext(ctx, fmt.Sprintf("collector interval: %s", interval))
	ticker := time.NewTicker(interval)

	perform := func() {
		if !w.occurredPlugInEvent(ctx) {
			return
		}
		w.r.Logger().Info("detected plug-in event")
		if err := w.r.ChargeService().SetEnabled(ctx, true); err != nil {
			w.r.Logger().Error("failed to enable charge", logx.ErrorAttr(err))
			return
		}
		if _, err := w.r.ChargeService().RefreshChargeState(ctx, w.r.AppConfig().TeslaVIN); err != nil {
			w.r.Logger().Error("failed to refresh charge state", logx.ErrorAttr(err))
			return
		}
		if err := w.r.ChargeService().Adjust(ctx); err != nil {
			w.r.Logger().Error("failed to adjust charging", logx.ErrorAttr(err))
			return
		}
	}
	for {
		select {
		case <-ctx.Done():
			w.r.Logger().Info("context done")
			return nil
		case <-ticker.C:
			perform()
		}
	}
}

func (w *Worker) collectMetricAndAdjustCharging(ctx context.Context) error {
	if err := w.collectSurplusPowerMetric(ctx); err != nil {
		return err
	}
	return w.r.ChargeService().Adjust(ctx)
}

func (w *Worker) collectSurplusPowerMetric(ctx context.Context) error {
	metric, err := w.r.Collector().GetSurplusPower(ctx)
	if err != nil {
		return err
	}
	return w.r.PowerMetricRepository().CreateOne(ctx, metric)
}

func (w *Worker) occurredPlugInEvent(ctx context.Context) bool {
	metric, err := w.r.Collector().GetEVUsagePower(ctx)
	if err != nil {
		w.r.Logger().Error("failed to get EV usage power", logx.ErrorAttr(err))
		return false
	}
	if metric.Watt <= 0 {
		return false
	}
	state, err := w.r.ChargeService().DescribeChargeState(ctx, w.r.AppConfig().TeslaVIN)
	if err != nil {
		w.r.Logger().Error("failed to get vehicle state", logx.ErrorAttr(err))
		return false
	}
	return state.ChargingState != "Charging"
}
