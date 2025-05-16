package cleaner

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormalizeComp(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
	}{
		{"USD with K suffix", "$10K", 10000},
		{"INR with L suffix", "₹5L", 6000},
		{"GBP with M suffix", "£2M", 2500000},
		{"EUR without suffix", "€1000", 1080},
		{"Invalid input", "abc", 0},
		{"Empty input", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := normalizeComp(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNormalizeLocation(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedCity  string
		expectedState string
	}{
		{"Comma-separated location", "San Francisco, CA", "San Francisco", "CA"},
		{"Space-separated location", "Los Angeles California", "Los Angeles", "California"},
		{"Single-word location", "Seattle", "Seattle", ""},
		{"Invalid location with digits", "12345", "", ""},
		{"Empty input", "", "", ""},
		{"Trailing spaces", "  Boston , MA  ", "Boston", "MA"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			city, state := normalizeLocation(tt.input)
			assert.Equal(t, tt.expectedCity, city)
			assert.Equal(t, tt.expectedState, state)
		})
	}
}

func TestCleanNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Number with dollar sign", "$1000", "1000"},
		{"Number with commas", "1,000,000", "1000000"},
		{"Number with spaces", "  500  ", "500"},
		{"Empty input", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := cleanNumber(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
