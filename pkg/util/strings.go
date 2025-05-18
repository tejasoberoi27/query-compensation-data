package util

import "unicode"

func ContainsDigit(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func StringToStringPtr(s string) *string {
	return &s
}

func Float64ToFloat64Ptr(f float64) *float64 {
	return &f
}
