package service

import (
	"context"
	"query-compensation-data/internal/compensation/domain/entity"
	compensation_service "query-compensation-data/internal/compensation/service/compensation"
)

type CompensationService interface {
	GetByID(ctx context.Context, id int) (*entity.Compensation, error)
	GetFilteredCompensations(ctx context.Context, filter compensation_service.Filter) ([]*entity.Compensation, error)
}
