package domain_navigation_test

import (
	service "dns/internal/domain_navigation"
	"testing"
)

// Format Format the value to 2 decimals
func TestDataBankService_Format(t *testing.T) {
	tests := []struct {
		input    service.Location
		expected service.Location
	}{
		{
			input:    777.56777,
			expected: 777.57,
		},
		{
			input:    777.53222,
			expected: 777.53,
		},
		{
			input:    777.547899,
			expected: 777.55,
		},
		{
			input:    777.501999,
			expected: 777.50,
		},
	}

	for _, test := range tests {
		loc := service.Location(test.input)
		result := loc.Format()

		if result != test.expected {
			t.Errorf("expected %f but got %f", test.expected, result)
		}
	}

}
