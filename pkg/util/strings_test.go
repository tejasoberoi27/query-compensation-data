package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDigit(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"String with digits", "abc123", true},
		{"String without digits", "abcdef", false},
		{"Empty string", "", false},
		{"String with special characters", "!@#$%^&*", false},
		{"String with only digits", "12345", true},
		{"String with spaces and digits", " 123 ", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ContainsDigit(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
