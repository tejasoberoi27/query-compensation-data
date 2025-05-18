package compensation

import (
	domain_service "query-compensation-data/internal/compensation/domain/service/compensation"
	"query-compensation-data/internal/compensation/handler/graph/compensation/generated"
)

type Handler struct {
	compensationService domain_service.CompensationService
}

type Resolver struct {
	compensationService domain_service.CompensationService
}

func NewResolver(compensationService domain_service.CompensationService) *Resolver {
	return &Resolver{compensationService: compensationService}
}

func (r *Resolver) Query() generated.QueryResolver {
	return Handler{compensationService: r.compensationService}
}
