package db

import (
	"context"
	"github.com/uptrace/bun"
	repository "query-compensation-data/internal/compensation/repository/compensation/postgres"
)

func InitTable(db bun.IDB) error {
	// Drop the table if it exists
	_, err := db.NewDropTable().Model((*repository.CompensationModel)(nil)).IfExists().Exec(context.Background())
	if err != nil {
		return err
	}

	// Create the table
	_, err = db.NewCreateTable().Model((*repository.CompensationModel)(nil)).IfNotExists().Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}
