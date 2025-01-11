package sqlite

import (
	"context"
	"time"

	"github.com/morikuni/failure/v2"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/model"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/repository"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/ent"
)

type chargeSettingRepository struct {
	client *ent.Client
}

var _ repository.ChargeSettingRepository = new(chargeSettingRepository)

func newChargeSettingRepository(client *ent.Client) *chargeSettingRepository {
	return &chargeSettingRepository{
		client: client,
	}
}

func (r *chargeSettingRepository) FindOne(ctx context.Context) (*model.ChargeSetting, error) {
	found, err := r.client.ChargeSetting.Query().Order(ent.Desc("id")).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, failure.New(model.ErrCodeNotFound, failure.Message("charge setting not found"))
		}
		return nil, err
	}

	minChargeTimeRangeStart, err := model.ParseTimeOnly(found.MinChargeTimeRangeStart)
	if err != nil {
		return nil, err
	}
	minChargeTimeRangeEnd, err := model.ParseTimeOnly(found.MinChargeTimeRangeEnd)
	if err != nil {
		return nil, err
	}

	return &model.ChargeSetting{
		Enabled:                     found.Enabled,
		ChargeStartThreshold:        found.ChargeStartThreshold,
		PowerUsageIncreaseThreshold: found.PowerUsageIncreaseThreshold,
		PowerUsageDecreaseThreshold: found.PowerUsageDecreaseThreshold,
		UpdateInterval:              time.Duration(found.UpdateInterval) * time.Minute,
		MinCharge: model.MinimumChargeSetting{
			Threshold: found.MinChargeThreshold,
			TimeRange: model.TimeRange{
				Start: minChargeTimeRangeStart,
				End:   minChargeTimeRangeEnd,
			},
			Amperage: found.MinChargeAmperage,
		},
	}, nil
}

func (r *chargeSettingRepository) SetEnabled(ctx context.Context, enabled bool) error {
	found, err := r.client.ChargeSetting.Query().Order(ent.Desc("id")).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return failure.New(model.ErrCodeNotFound, failure.Message("charge setting not found"))
		}
		return err
	}
	_, err = found.Update().SetEnabled(enabled).Save(ctx)
	return err
}

func (r *chargeSettingRepository) SaveOne(ctx context.Context, setting *model.ChargeSetting) error {
	found, err := r.client.ChargeSetting.Query().Order(ent.Desc("id")).First(ctx)
	if err == nil {
		_, err = found.Update().
			SetEnabled(setting.Enabled).
			SetChargeStartThreshold(setting.ChargeStartThreshold).
			SetPowerUsageIncreaseThreshold(setting.PowerUsageIncreaseThreshold).
			SetPowerUsageDecreaseThreshold(setting.PowerUsageDecreaseThreshold).
			SetUpdateInterval(int(setting.UpdateInterval.Minutes())).
			SetMinChargeThreshold(setting.MinCharge.Threshold).
			SetMinChargeTimeRangeStart(setting.MinCharge.TimeRange.Start.HourMinute()).
			SetMinChargeTimeRangeEnd(setting.MinCharge.TimeRange.End.HourMinute()).
			SetMinChargeAmperage(setting.MinCharge.Amperage).
			Save(ctx)
		return err
	}
	if !ent.IsNotFound(err) {
		return err
	}
	_, err = r.client.ChargeSetting.Create().
		SetEnabled(setting.Enabled).
		SetChargeStartThreshold(setting.ChargeStartThreshold).
		SetPowerUsageIncreaseThreshold(setting.PowerUsageIncreaseThreshold).
		SetPowerUsageDecreaseThreshold(setting.PowerUsageDecreaseThreshold).
		SetUpdateInterval(int(setting.UpdateInterval.Minutes())).
		SetMinChargeThreshold(setting.MinCharge.Threshold).
		SetMinChargeTimeRangeStart(setting.MinCharge.TimeRange.Start.HourMinute()).
		SetMinChargeTimeRangeEnd(setting.MinCharge.TimeRange.End.HourMinute()).
		SetMinChargeAmperage(setting.MinCharge.Amperage).
		Save(ctx)
	return err
}
