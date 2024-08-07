package sqlite

import (
	"context"
	"time"

	"github.com/morikuni/failure/v2"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/model"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/repository"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/ent"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/ent/chargecommandhistory"
)

type chargeCommandHistoryRepository struct {
	client *ent.Client
}

var _ repository.ChargeCommandHistoryRepository = new(chargeCommandHistoryRepository)

func newChargeCommandHistoryRepository(client *ent.Client) *chargeCommandHistoryRepository {
	return &chargeCommandHistoryRepository{
		client: client,
	}
}

func (r *chargeCommandHistoryRepository) FindLatestOne(ctx context.Context) (*model.ChargeCommandHistory, error) {
	found, err := r.client.ChargeCommandHistory.Query().Order(ent.Desc("id")).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, failure.New(model.ErrCodeNotFound, failure.Message("charge command history not found"))
		}
		return nil, err
	}
	return &model.ChargeCommandHistory{
		VIN:       found.Vin,
		Operation: model.ChargeCommandOperation(found.Operation),
		Amps:      found.Amps,
		Timestamp: found.Timestamp,
	}, nil
}

func (r *chargeCommandHistoryRepository) CreateOne(ctx context.Context, history *model.ChargeCommandHistory) error {
	_, err := r.client.ChargeCommandHistory.Create().
		SetVin(history.VIN).
		SetOperation(string(history.Operation)).
		SetAmps(history.Amps).
		SetTimestamp(history.Timestamp).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *chargeCommandHistoryRepository) DeleteOlderThan(ctx context.Context, t time.Time) error {
	_, err := r.client.ChargeCommandHistory.Delete().Where(chargecommandhistory.TimestampLT(t)).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
