package app

import (
	"go.uber.org/fx"
	"query-compensation-data/internal/compensation/app/container/env"
	"query-compensation-data/internal/compensation/app/container/repository"
	"query-compensation-data/internal/compensation/config"
)

type App struct {
	app *fx.App
}

func New() App {
	return App{
		app: fx.New(
			BuildGraph(),
			fx.Invoke(),
		),
	}
}

func (a App) Run() {
	a.app.Run()
}

func BuildGraph() fx.Option {
	return fx.Options(
		fx.NopLogger,
		fx.Provide(
			config.New,
		),
		repository.New(),

		env.New(),
	)
}
