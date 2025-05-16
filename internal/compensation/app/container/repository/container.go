package repository

import (
	"go.uber.org/fx"
	"query-compensation-data/internal/compensation/env"
	"query-compensation-data/pkg/repository/postgres"
)

func New() fx.Option {
	postgresDB := func(env *env.Env) *postgres.DB {
		return env.Postgres.DB
	}

	return fx.Provide(
		postgresDB,
	)
}
