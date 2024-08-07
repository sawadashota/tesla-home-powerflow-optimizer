package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/service"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/collector"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/repository"
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

func (w *Worker) Run(ctx context.Context) error {
	interval := w.r.AppConfig().CollectorIntervalDuration()
	w.r.Logger().Info(fmt.Sprintf("collector interval: %s", interval))
	ticker := time.NewTicker(interval)
	var perform = func() {
		w.r.Logger().Info("collecting surplus power...")
		if err := w.perform(ctx); err != nil {
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

func (w *Worker) perform(ctx context.Context) error {
	if err := w.collect(ctx); err != nil {
		return err
	}
	return w.r.ChargeService().Adjust(ctx)
}

func (w *Worker) collect(ctx context.Context) error {
	metric, err := w.r.Collector().GetSurplusPower(ctx)
	if err != nil {
		return err
	}
	return w.r.PowerMetricRepository().CreateOne(ctx, metric)
}
