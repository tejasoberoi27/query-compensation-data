package compensation

import (
	"query-compensation-data/internal/compensation/domain/repository"
	"query-compensation-data/internal/compensation/env"
	"query-compensation-data/internal/compensation/repository/compensation"
)

type Service struct {
	repository repository.CompensationRepository
}

func New(env *env.Env) *Service {
	return &Service{
		repository: compensation.New(env),
	}
}
