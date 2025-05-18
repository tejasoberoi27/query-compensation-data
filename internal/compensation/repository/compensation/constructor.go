package compensation

import (
	"query-compensation-data/internal/compensation/domain/repository"
	"query-compensation-data/internal/compensation/env"
	"query-compensation-data/internal/compensation/repository/compensation/postgres"
)

func New(env *env.Env) repository.CompensationRepository {
	return postgres.New(env.Postgres.DB)
}
