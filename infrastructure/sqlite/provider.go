package sqlite

import (
	"context"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/repository"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/driver/configuration"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/ent"
)

type provider struct {
	client          *ent.Client
	grantRepository repository.GrantRepository
}

var _ repository.Provider = new(provider)

func newClient(dsn string) (*ent.Client, error) {
	client, err := ent.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewProvider(config configuration.AppConfigProvider) (repository.Provider, error) {
	client, err := newClient(config.AppConfig().SqliteDSN)
	if err != nil {
		return nil, err
	}
	return &provider{
		client: client,
	}, nil
}

func (p *provider) Migrate(ctx context.Context) error {
	return p.client.Schema.Create(ctx)
}

func (p *provider) GrantRepository() repository.GrantRepository {
	if p.grantRepository == nil {
		p.grantRepository = newGrantRepository(p.client)
	}
	return p.grantRepository
}
