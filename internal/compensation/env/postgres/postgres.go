package postgres

import (
	"github.com/uptrace/bun/extra/bundebug"
	"query-compensation-data/internal/compensation/config"
	"query-compensation-data/pkg/repository/postgres"
)

type Postgres struct {
	DB *postgres.DB
}

func New(conf config.Postgres) (*Postgres, error) {
	pg := &Postgres{
		DB: postgres.NewDB(
			postgres.Config{
				Host:     conf.Host,
				Port:     conf.Port,
				Database: conf.Database,
				User:     conf.User,
				Password: conf.Password,
			},
		),
	}

	pg.DB.DB.SetMaxOpenConns(200)
	pg.DB.DB.SetMaxIdleConns(100)

	pg.DB.DB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	return pg, nil
}
