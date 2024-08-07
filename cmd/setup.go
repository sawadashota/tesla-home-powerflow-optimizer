package cmd

import (
	"context"

	"github.com/morikuni/failure/v2"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/model"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/repository"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/internal/logx"
)

type setupDependencies interface {
	logx.Provider
	repository.MigrationProvider
	repository.ChargeSettingRepositoryProvider
}

func setup(ctx context.Context, r setupDependencies) error {
	r.Logger().Info("Migrating database...")
	if err := r.Migrate(ctx); err != nil {
		return err
	}

	_, err := r.ChargeSettingRepository().FindOne(ctx)
	if err != nil {
		if failure.Is(err, model.ErrCodeNotFound) {
			r.Logger().Info("Creating default charge setting...")
			if err := r.ChargeSettingRepository().SaveOne(ctx, model.ChargeSettingDefault); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}
