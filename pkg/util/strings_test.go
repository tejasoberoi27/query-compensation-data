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

func TestStringToStringPtr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *string
	}{
		{"Non-empty string", "hello", StringToStringPtr("hello")},
		{"Empty string", "", StringToStringPtr("")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StringToStringPtr(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFloat64ToFloat64Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected *float64
	}{
		{"Positive float", 123.45, Float64ToFloat64Ptr(123.45)},
		{"Zero value", 0.0, Float64ToFloat64Ptr(0.0)},
		{"Negative float", -987.65, Float64ToFloat64Ptr(-987.65)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Float64ToFloat64Ptr(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
