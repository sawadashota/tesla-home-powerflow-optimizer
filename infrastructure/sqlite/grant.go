package sqlite

import (
	"context"

	"entgo.io/ent/dialect/sql"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/model"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/repository"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/ent"
	entgrant "github.com/sawadashota/tesla-home-powerflow-optimizer/ent/grant"
)

type grantRepository struct {
	client *ent.Client
}

var _ repository.GrantRepository = new(grantRepository)

func newGrantRepository(client *ent.Client) *grantRepository {
	return &grantRepository{
		client: client,
	}
}

func (r *grantRepository) FindLatestOne(ctx context.Context) (*model.Grant, error) {
	found, err := r.client.Grant.Query().Order(entgrant.ByID(sql.OrderDesc())).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.ErrGrantNotFound
		}
		return nil, err
	}
	return &model.Grant{
		Subject:      found.Subject,
		AccessToken:  found.AccessToken,
		RefreshToken: found.RefreshToken,
		Scope:        found.Scope,
		Expiry:       found.Expiry,
	}, nil
}

func (r *grantRepository) SaveOne(ctx context.Context, grant *model.Grant) error {
	found, err := r.client.Grant.Query().Order(entgrant.ByID(sql.OrderDesc())).First(ctx)
	if err == nil {
		_, err = found.Update().
			SetSubject(grant.Subject).
			SetAccessToken(grant.AccessToken).
			SetRefreshToken(grant.RefreshToken).
			SetScope(grant.Scope).
			SetExpiry(grant.Expiry).
			Save(ctx)
		return err
	}
	if !ent.IsNotFound(err) {
		return err
	}
	_, err = r.client.Grant.Create().
		SetSubject(grant.Subject).
		SetAccessToken(grant.AccessToken).
		SetRefreshToken(grant.RefreshToken).
		SetScope(grant.Scope).
		SetExpiry(grant.Expiry).
		Save(ctx)
	return err
}
