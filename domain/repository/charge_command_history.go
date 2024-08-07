package repository

import (
	"context"
	"time"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/model"
)

type (
	ChargeCommandHistoryRepository interface {
		// FindLatestOne finds the latest charge command history
		// If there is no charge command history, it returns an error with *model.ErrCodeNotFound
		FindLatestOne(ctx context.Context) (*model.ChargeCommandHistory, error)

		// CreateOne saves the charge command history
		CreateOne(ctx context.Context, history *model.ChargeCommandHistory) error

		// DeleteOlderThan deletes charge command history older than t
		DeleteOlderThan(ctx context.Context, t time.Time) error
	}
	ChargeCommandHistoryRepositoryProvider interface {
		ChargeCommandHistoryRepository() ChargeCommandHistoryRepository
	}
)
