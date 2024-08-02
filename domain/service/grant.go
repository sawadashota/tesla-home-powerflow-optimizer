package service

import (
	"context"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/model"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/repository"
)

type (
	GrantService struct {
		r GrantServiceDependencies
	}
	GrantServiceDependencies interface {
		repository.GrantRepositoryProvider
	}
	GrantServiceProvider interface {
		GrantService() *GrantService
	}
)

func NewGrantService(r GrantServiceDependencies) *GrantService {
	return &GrantService{r: r}
}

func (s *GrantService) FindLatestOne(ctx context.Context) (*model.Grant, error) {
	return s.r.GrantRepository().FindLatestOne(ctx)
}

func (s *GrantService) Save(ctx context.Context, grant *model.Grant) error {
	if err := grant.Validate(); err != nil {
		return err
	}
	return s.r.GrantRepository().SaveOne(ctx, grant)
}
