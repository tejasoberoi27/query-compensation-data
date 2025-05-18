package compensation

import "query-compensation-data/internal/compensation/handler/graph/compensation/model"

// SortKey represents the sorting criteria for compensations.
type SortKey string

const (
	SortKeyTotalCompAsc  SortKey = "TOTALCOMP_ASC"  // Sort by total compensation in ascending order
	SortKeyTotalCompDesc SortKey = "TOTALCOMP_DESC" // Sort by total compensation in descending order
	SortKeyTimestampAsc  SortKey = "TIMESTAMP_ASC"  // Sort by timestamp in ascending order
	SortKeyTimestampDesc SortKey = "TIMESTAMP_DESC" // Sort by timestamp in descending order
)

// NewSortKey creates a new SortKey from a model.SortKey.
// If the provided model.SortKey is invalid, it returns nil.
func NewSortKey(modelKey model.CompensationSort) *SortKey {
	v := SortKeyTotalCompAsc
	switch modelKey {
	case model.CompensationSortTimestampAsc:
		v = SortKeyTimestampAsc
		return &v
	case model.CompensationSortTimestampDesc:
		v = SortKeyTimestampDesc
		return &v
	case model.CompensationSortTotalcompAsc:
		v = SortKeyTotalCompAsc
		return &v
	case model.CompensationSortTotalcompDesc:
		v = SortKeyTotalCompDesc
		return &v
	default:
		return nil
	}
}

// IsValid checks if the SortKey is valid.
func (s SortKey) IsValid() bool {
	switch s {
	case SortKeyTotalCompAsc, SortKeyTotalCompDesc, SortKeyTimestampAsc, SortKeyTimestampDesc:
		return true
	}
	return false
}
