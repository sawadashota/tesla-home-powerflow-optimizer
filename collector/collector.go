package collector

import (
	"context"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/model"
)

type (
	Collector interface {
		GetSurplusPower(ctx context.Context) (*model.PowerMetric, error)
	}
	Provider interface {
		Collector() Collector
	}
)
