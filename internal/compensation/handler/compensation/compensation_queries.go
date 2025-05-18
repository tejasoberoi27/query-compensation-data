package compensation

import (
	"context"
	"query-compensation-data/internal/compensation/handler/graph/compensation/model"
	service "query-compensation-data/internal/compensation/service/compensation"
	"query-compensation-data/pkg/util"
	"strconv"
	"time"
)

func (h Handler) Compensation(ctx context.Context, id string) (*model.CompensationResult, error) {
	// Convert the string ID to an integer
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return &model.CompensationResult{
			Error: &model.CompensationError{
				Message: "Invalid ID format",
			},
		}, nil
	}

	// Fetch the compensation entity using the service
	compensationEntity, err := h.compensationService.GetByID(ctx, idInt)
	if err != nil {
		return &model.CompensationResult{
			Error: &model.CompensationError{
				Message: "Failed to fetch compensation",
			},
		}, nil
	}

	// If no compensation is found, return an error result
	if compensationEntity == nil {
		return &model.CompensationResult{
			Error: &model.CompensationError{
				Message: "Compensation not found",
			},
		}, nil
	}

	// Map the entity to the GraphQL model
	return &model.CompensationResult{
		Compensation: &model.Compensation{
			ID:                 strconv.Itoa(compensationEntity.ID),
			Timestamp:          compensationEntity.Timestamp.Format(time.RFC3339),
			Company:            util.StringToStringPtr(compensationEntity.Company),
			Title:              util.StringToStringPtr(compensationEntity.Title),
			City:               util.StringToStringPtr(compensationEntity.City),
			State:              util.StringToStringPtr(compensationEntity.State),
			TotalComp:          util.Float64ToFloat64Ptr(compensationEntity.TotalComp),
			SigningBonus:       util.Float64ToFloat64Ptr(compensationEntity.SigningBonus),
			BaseSalary:         util.Float64ToFloat64Ptr(compensationEntity.AnnualBasePay),
			AnnualBonus:        util.Float64ToFloat64Ptr(compensationEntity.AnnualBonus),
			AnnualStockValue:   util.Float64ToFloat64Ptr(compensationEntity.AnnualStockValue),
			YearsExp:           util.Float64ToFloat64Ptr(compensationEntity.YearsExp),
			AdditionalComments: util.StringToStringPtr(compensationEntity.AdditionalComments),
			Gender:             func() *model.Gender { g := model.Gender(compensationEntity.Gender); return &g }(),
			YearsAtCompany:     util.Float64ToFloat64Ptr(compensationEntity.YearsAtEmployer),
		},
	}, nil
}

func (h Handler) Compensations(
	ctx context.Context,
	titleContains *string,
	city *string,
	state *string,
	company *string,
	minTotalComp *float64,
	maxTotalComp *float64,
	sortBy *model.CompensationSort,
	limit *int,
	offset *int,
) (*model.CompensationsResult, error) {
	// Build the filter for the service layer
	filter := service.Filter{
		TitleContains: titleContains,
		City:          city,
		State:         state,
		Company:       company,
		MinTotalComp:  minTotalComp,
		MaxTotalComp:  maxTotalComp,
		SortBy:        service.NewSortKey(*sortBy),
		Limit:         limit,
		Offset:        offset,
	}

	// Fetch filtered compensations from the service
	compensationEntities, err := h.compensationService.GetFilteredCompensations(ctx, filter)
	if err != nil {
		return &model.CompensationsResult{
			Error: &model.CompensationError{
				Message: "Failed to fetch compensations",
			},
		}, nil
	}

	// Map to GraphQL model
	result := []*model.Compensation{}
	for _, comp := range compensationEntities {
		if comp == nil {
			continue
		}
		gender := model.Gender(comp.Gender.String())
		result = append(result, &model.Compensation{
			ID:                 strconv.Itoa(comp.ID),
			Timestamp:          comp.Timestamp.Format(time.RFC3339),
			Company:            util.StringToStringPtr(comp.Company),
			Title:              util.StringToStringPtr(comp.Title),
			City:               util.StringToStringPtr(comp.City),
			State:              util.StringToStringPtr(comp.State),
			TotalComp:          util.Float64ToFloat64Ptr(comp.TotalComp),
			SigningBonus:       util.Float64ToFloat64Ptr(comp.SigningBonus),
			BaseSalary:         util.Float64ToFloat64Ptr(comp.AnnualBasePay),
			AnnualBonus:        util.Float64ToFloat64Ptr(comp.AnnualBonus),
			AnnualStockValue:   util.Float64ToFloat64Ptr(comp.AnnualStockValue),
			YearsExp:           util.Float64ToFloat64Ptr(comp.YearsExp),
			YearsAtCompany:     util.Float64ToFloat64Ptr(comp.YearsAtEmployer),
			AdditionalComments: util.StringToStringPtr(comp.AdditionalComments),
			Gender:             &gender,
		})
	}

	return &model.CompensationsResult{
		Compensations: result,
		Count:         len(result),
	}, nil
}
