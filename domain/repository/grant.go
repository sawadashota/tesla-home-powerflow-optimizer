package repository

import (
	"context"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/model"
)

type (
	GrantRepository interface {
		// FindLatestOne finds the latest grant
		// If there is no grant, it returns *model.ErrGrantNotFound
		FindLatestOne(ctx context.Context) (*model.Grant, error)
		SaveOne(ctx context.Context, grant *model.Grant) error
	}

	GrantRepositoryProvider interface {
		GrantRepository() GrantRepository
	}
)
