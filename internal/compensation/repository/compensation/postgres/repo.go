package postgres

import (
	"context"
	"query-compensation-data/internal/compensation/domain/entity"
	"query-compensation-data/pkg/repository/postgres"
)

type CompensationRepository struct {
	db *postgres.DB
}

func (r CompensationRepository) Create(ctx context.Context, i *entity.Compensation) error {
	model := newModel(i)

	h, err := r.db.GetHandler()
	if err != nil {
		return err
	}

	_, err = h.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return err
	}

	i.ID = model.Convert().ID
	return nil
}

func New(db *postgres.DB) CompensationRepository {
	return CompensationRepository{
		db: db,
	}
}
