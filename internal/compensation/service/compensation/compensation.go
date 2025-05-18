package compensation

import (
	"context"
	"log"
	"query-compensation-data/internal/compensation/domain/entity"
	"query-compensation-data/internal/compensation/domain/repository"
	"sort"
)

// Filter represents the filtering, sorting, and pagination parameters for fetching compensations.
type Filter struct {
	TitleContains *string  // Filter by title containing this string
	City          *string  // Filter by city
	State         *string  // Filter by state
	Company       *string  // Filter by company
	MinTotalComp  *float64 // Minimum total compensation
	MaxTotalComp  *float64 // Maximum total compensation
	SortBy        *SortKey // Sorting criteria (e.g., "totalCompAsc", "timestampDesc")
	Limit         *int     // Maximum number of records to return
	Offset        *int     // Number of records to skip
}

// GetFilteredCompensations retrieves a list of compensations based on the provided filter criteria.
// It applies filtering, sorting, and pagination logic.
func (s Service) GetFilteredCompensations(ctx context.Context, filter Filter) ([]*entity.Compensation, error) {
	// Build the request data for filtering
	requestData := repository.RequestData{
		TitleContains: filter.TitleContains,
		City:          filter.City,
		State:         filter.State,
		Company:       filter.Company,
	}

	// Fetch entities using the repository
	compensationEntities, err := s.repository.GetBulkByFields(ctx, requestData)
	if err != nil {
		return nil, err
	}

	// Filter by min and max total compensation
	filteredCompensations := []*entity.Compensation{}
	for _, comp := range compensationEntities {
		if comp == nil {
			continue
		}
		if (filter.MinTotalComp == nil || comp.TotalComp >= *filter.MinTotalComp) &&
			(filter.MaxTotalComp == nil || comp.TotalComp <= *filter.MaxTotalComp) {
			filteredCompensations = append(filteredCompensations, comp)
		}
	}

	// Apply sorting
	if filter.SortBy != nil {
		sort.Slice(filteredCompensations, func(i, j int) bool {
			// Use the SortKey enum for sorting logic
			switch *filter.SortBy {
			case SortKeyTotalCompAsc:
				return filteredCompensations[i].TotalComp < filteredCompensations[j].TotalComp
			case SortKeyTotalCompDesc:
				return filteredCompensations[i].TotalComp > filteredCompensations[j].TotalComp
			case SortKeyTimestampAsc:
				return filteredCompensations[i].Timestamp.Before(filteredCompensations[j].Timestamp)
			case SortKeyTimestampDesc:
				return filteredCompensations[i].Timestamp.After(filteredCompensations[j].Timestamp)
			default:
				log.Printf("unknown sort key: %s", *filter.SortBy)
				return false
			}
		})
	}

	// Apply pagination
	start := 0
	if filter.Offset != nil {
		start = *filter.Offset
	}
	if start >= len(filteredCompensations) {
		return []*entity.Compensation{}, nil
	}

	end := len(filteredCompensations)
	if filter.Limit != nil && start+*filter.Limit < end {
		end = start + *filter.Limit
	}
	return filteredCompensations[start:end], nil
}

// GetByID retrieves a single compensation record by its ID.
// It calls the repository method to fetch the compensation entity.
func (s Service) GetByID(ctx context.Context, id int) (*entity.Compensation, error) {
	// Call the repository method to fetch the compensation by ID
	compensation, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err // Return the error if the repository call fails
	}

	// Return the fetched compensation entity
	return compensation, nil
}
