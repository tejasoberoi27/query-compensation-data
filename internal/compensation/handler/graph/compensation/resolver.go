package compensation

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

import (
	"context"
	"query-compensation-data/internal/compensation/handler/graph/compensation/generated"
	"query-compensation-data/internal/compensation/handler/graph/compensation/model"
)

type Resolver struct{}

// Compensation is the resolver for the compensation field.
func (r *queryResolver) Compensation(ctx context.Context, id string) (*model.CompensationResult, error) {
	//r.compensationService.GetByFields(ctx, service.RequestData)
	//
	//return &model.Compensation{
	//	ID:                 util.StringToStringPtr("3223"),
	//	Timestamp:          util.StringToStringPtr(time.Now().Format(time.RFC3339)),
	//	Company:            util.StringToStringPtr("TechCorp"),
	//	TitleContains:              util.StringToStringPtr("Software Engineer"),
	//	City:               util.StringToStringPtr("San Francisco"),
	//	State:              util.StringToStringPtr("CA"),
	//	TotalComp:          util.Float64ToFloat64Ptr(155000.0),
	//	BaseSalary:         util.Float64ToFloat64Ptr(120000.0),
	//	AnnualBonus:        util.Float64ToFloat64Ptr(15000.0),
	//	AnnualStockValue:   util.Float64ToFloat64Ptr(20000.0),
	//	YearsExp:           util.Float64ToFloat64Ptr(5),
	//	AdditionalComments: util.StringToStringPtr("good work environment"),
	//	Gender:             func() *model.Gender { g := model.GenderMale; return &g }(),
	//	YearsAtCompany:     util.Float64ToFloat64Ptr(2),
	//}, nil
	panic("not implemented")
}

// Compensations is the resolver for the compensations field.
func (r *queryResolver) Compensations(ctx context.Context, titleContains *string, city *string, state *string, company *string, minTotalComp *float64, maxTotalComp *float64, sortBy *model.CompensationSort, limit *int, offset *int) (*model.CompensationsResult, error) {
	panic("not implemented")
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	type Resolver struct{}
*/
