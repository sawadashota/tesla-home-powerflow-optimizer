package repository

import (
	"context"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/model"
)

type (
	ChargeSettingRepository interface {
		// FindOne finds the latest charge setting
		// If there is no charge setting, it returns an error with *model.ErrCodeNotFound
		FindOne(ctx context.Context) (*model.ChargeSetting, error)

		// SaveOne saves the charge setting
		SaveOne(ctx context.Context, setting *model.ChargeSetting) error

		// SetEnabled sets the charge setting enabled
		//
		// If there is no charge setting, it returns an error with *model.ErrCodeNotFound
		SetEnabled(ctx context.Context, enabled bool) error
	}
	ChargeSettingRepositoryProvider interface {
		ChargeSettingRepository() ChargeSettingRepository
	}
)
