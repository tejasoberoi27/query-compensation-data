package postgres

import (
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Config struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
	RootCAs  string
}

type DB struct {
	DB *bun.DB
}

func NewDB(cfg Config) *DB {
	connector := pgdriver.NewConnector(
		pgdriver.WithAddr(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)),
		pgdriver.WithUser(cfg.User),
		pgdriver.WithDatabase(cfg.Database),
		pgdriver.WithPassword(cfg.Password),
		pgdriver.WithInsecure(true),
	)

	return &DB{
		DB: bun.NewDB(sql.OpenDB(connector), pgdialect.New(), bun.WithDiscardUnknownColumns()),
	}
}

func (db *DB) GetHandler() (bun.IDB, error) {
	return db.DB, nil
}

type ErrPostgres struct {
	Err error
}

func (e ErrPostgres) Error() string {
	return e.Err.Error()
}

func (e ErrPostgres) Unwrap() error {
	return e.Err
}
