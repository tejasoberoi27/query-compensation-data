package db_population

import (
	"context"
	"github.com/uptrace/bun"
	"query-compensation-data/internal/compensation/repository/compensation/postgres"
)

func InsertData(db bun.IDB, entries []postgres.CompensationModel) error {
	_, err := db.NewInsert().Model(&entries).Exec(context.Background())
	return err
}
