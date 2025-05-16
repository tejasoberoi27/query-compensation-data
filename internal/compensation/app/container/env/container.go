package env

import (
	"go.uber.org/fx"
	"query-compensation-data/internal/compensation/env"
)

func New() fx.Option {
	return fx.Provide(
		env.New,
	)
}
