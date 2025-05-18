package repository

import (
	"context"
	"query-compensation-data/internal/compensation/domain/entity"
	"query-compensation-data/internal/compensation/enum"
	"time"
)

type RequestData struct {
	ID                 *int
	Timestamp          *time.Time
	Company            *string
	TitleContains      *string
	City               *string
	State              *string
	TotalComp          *float64
	AnnualBasePay      *float64
	AnnualBonus        *float64
	AnnualStockValue   *float64
	YearsExp           *float64
	AdditionalComments *string
	Gender             *enum.Gender
	SigningBonus       *float64
	YearsAtEmployer    *float64
}

type CompensationRepository interface {
	GetByID(ctx context.Context, id int) (*entity.Compensation, error)
	GetBulkByFields(ctx context.Context, data RequestData) ([]*entity.Compensation, error)
}
