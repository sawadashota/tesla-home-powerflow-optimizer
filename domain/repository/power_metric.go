package repository

import (
	"context"
	"time"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/model"
)

type (
	PowerMetricRepository interface {
		// FindLatestN finds the latest power metrics
		//
		// If there is no power metric, it returns an error with *model.ErrCodeNotFound
		FindLatestN(ctx context.Context, n int) (model.PowerMetricList, error)

		// CreateOne creates a power metric
		CreateOne(ctx context.Context, powerMetric *model.PowerMetric) error

		// DeleteOlderThan deletes power metrics older than the specified time
		DeleteOlderThan(ctx context.Context, t time.Time) error
	}
	PowerMetricRepositoryProvider interface {
		PowerMetricRepository() PowerMetricRepository
	}
)
