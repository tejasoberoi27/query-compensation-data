package env

import (
	"fmt"
	"query-compensation-data/internal/compensation/config"
	"query-compensation-data/internal/compensation/env/postgres"
)

type Env struct {
	Postgres *postgres.Postgres
	Host     string
}

func New(cfg config.Config) (*Env, error) {
	postgresEnv, err := postgres.New(cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("failed to create a postgresql env: %w", err)
	}

	return &Env{
		Postgres: postgresEnv,
		Host:     cfg.APIServer.Port,
	}, nil
}
