package repository

import "context"

type MigrationProvider interface {
	Migrate(ctx context.Context) error
}

type Provider interface {
	MigrationProvider

	GrantRepositoryProvider
}
