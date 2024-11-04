package sqlite

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/morikuni/failure/v2"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/model"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/repository"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/ent"
	entpowermetric "github.com/sawadashota/tesla-home-powerflow-optimizer/ent/powermetric"
)

type powerMetricRepository struct {
	client *ent.Client
}

var _ repository.PowerMetricRepository = new(powerMetricRepository)

func newPowerMetricRepository(client *ent.Client) *powerMetricRepository {
	return &powerMetricRepository{
		client: client,
	}
}

func (r *powerMetricRepository) CreateOne(ctx context.Context, powerMetric *model.PowerMetric) error {
	_, err := r.client.PowerMetric.Create().
		SetSurplusWatt(powerMetric.SurplusWatt).
		SetTimestamp(powerMetric.Timestamp).
		Save(ctx)
	return err
}

func (r *powerMetricRepository) FindLatestN(ctx context.Context, n int) (model.PowerMetricList, error) {
	founds, err := r.client.PowerMetric.Query().Order(entpowermetric.ByID(sql.OrderDesc())).Limit(n).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(founds) == 0 {
		return nil, failure.New(model.ErrCodeNotFound, failure.Message("power metric not found"))
	}
	res := make(model.PowerMetricList, 0, len(founds))
	for _, found := range founds {
		res = append(res, &model.PowerMetric{
			SurplusWatt: found.SurplusWatt,
			Timestamp:   found.Timestamp,
		})
	}
	return res, nil
}

func (r *powerMetricRepository) DeleteOlderThan(ctx context.Context, t time.Time) error {
	_, err := r.client.PowerMetric.Delete().Where(entpowermetric.TimestampLT(t)).Exec(ctx)
	return err
}
