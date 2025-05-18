package postgres

import (
	"context"
	"database/sql"
	"errors"
	"query-compensation-data/internal/compensation/domain/entity"
	"query-compensation-data/internal/compensation/domain/repository"
	"query-compensation-data/pkg/repository/postgres"
)

type CompensationRepository struct {
	db *postgres.DB
}

func (r CompensationRepository) GetByID(ctx context.Context, id int) (*entity.Compensation, error) {
	model := new(CompensationModel)

	// Get the database handler
	h, err := r.db.GetHandler()
	if err != nil {
		return nil, err
	}

	// Build the query to fetch the compensation by ID
	err = h.NewSelect().
		Model(model).
		Where("id = ?", id).
		Scan(ctx)

	// Handle the case where no rows are found
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	// Handle other errors
	if err != nil {
		return nil, err
	}

	// Convert the model to the entity and return it
	return model.Convert(), nil
}

func (r CompensationRepository) GetBulkByFields(ctx context.Context, data repository.RequestData) ([]*entity.Compensation, error) {
	models := []*CompensationModel{}

	// Get the database handler
	h, err := r.db.GetHandler()
	if err != nil {
		return nil, err
	}

	// Build the query to fetch compensations based on the provided fields
	query := h.NewSelect().Model(&models)

	if data.ID != nil {
		query.Where("id = ?", *data.ID)
	}
	if data.Timestamp != nil {
		query.Where("timestamp = ?", *data.Timestamp)
	}
	if data.Company != nil {
		query.Where("company ilike ?", *data.Company)
	}
	if data.TitleContains != nil {
		query.Where("title ILIKE ?", "%"+*data.TitleContains+"%")
	}
	if data.City != nil {
		query.Where("city ilike ?", *data.City)
	}
	if data.State != nil {
		query.Where("state = ?", *data.State)
	}
	if data.TotalComp != nil {
		query.Where("total_comp = ?", *data.TotalComp)
	}
	if data.AnnualBasePay != nil {
		query.Where("annual_base_pay = ?", *data.AnnualBasePay)
	}
	if data.AnnualBonus != nil {
		query.Where("annual_bonus = ?", *data.AnnualBonus)
	}
	if data.AnnualStockValue != nil {
		query.Where("annual_stock_value = ?", *data.AnnualStockValue)
	}
	if data.YearsExp != nil {
		query.Where("years_exp = ?", *data.YearsExp)
	}
	if data.AdditionalComments != nil {
		query.Where("additional_comments = ?", *data.AdditionalComments)
	}
	if data.Gender != nil {
		query.Where("gender = ?", *data.Gender)
	}
	if data.SigningBonus != nil {
		query.Where("signing_bonus = ?", *data.SigningBonus)
	}
	if data.YearsAtEmployer != nil {
		query.Where("years_at_employer = ?", *data.YearsAtEmployer)
	}

	// Execute the query
	err = query.Scan(ctx)
	if err != nil {
		return nil, err
	}

	// Convert the models to entities
	compensations := []*entity.Compensation{}
	for _, model := range models {
		compensations = append(compensations, model.Convert())
	}

	return compensations, nil
}

func New(db *postgres.DB) CompensationRepository {
	return CompensationRepository{
		db: db,
	}
}
