package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateCardExpirationDate(t *testing.T) {
	testCases := []struct {
		name  string
		month string
		year  string
		err   error
	}{
		{
			name:  "uncorrect month length",
			month: "1",
			year:  "2025",
			err:   ErrInvalidCardExpMonth,
		},
		{
			name:  "uncorrect month symbol",
			month: "1*",
			year:  "2025",
			err:   ErrInvalidCardExpMonth,
		},
		{
			name:  "not existing month number",
			month: "15",
			year:  "2025",
			err:   ErrInvalidCardExpMonth,
		},
		{
			name:  "uncorrect year length",
			month: "10",
			year:  "20250",
			err:   ErrInvalidCardExpYear,
		},
		{
			name:  "uncorrect year symbol",
			month: "10",
			year:  "20*5",
			err:   ErrInvalidCardExpYear,
		},
		{
			name:  "expired date",
			month: "10",
			year:  "2022",
			err:   ErrExpiredCard,
		},
		{
			name:  "valid date",
			month: "10",
			year:  "2025",
			err:   nil,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := ValidateCardExpirationDate(tC.month, tC.year)

			require.Equal(t, tC.err, actual)
		})
	}
}

func TestValidateCardNumber(t *testing.T) {
	testCases := []struct {
		name   string
		number string
		err    error
	}{
		{
			name:   "empty number",
			number: "",
			err:    ErrInvalidCardNumberLength,
		},
		{
			name:   "invalid number length (too long)",
			number: "41111111111111111111111",
			err:    ErrInvalidCardNumberLength,
		},
		{
			name:   "invalid number length (too short)",
			number: "411111111111",
			err:    ErrInvalidCardNumberLength,
		},
		{
			name:   "invalid number symbol",
			number: "411111111111111@",
			err:    ErrInvalidCardNumberSymbol,
		},
		{
			name:   "invalid number according to luna algorithm",
			number: "4111111111111121",
			err:    ErrInvalidCardNumber,
		},
		{
			name:   "invalid number according to luna algorithm",
			number: "4111111111111121",
			err:    ErrInvalidCardNumber,
		},
		{
			name:   "valid number",
			number: "4111111111111111",
			err:    nil,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := ValidateCardNumber(tC.number)

			require.Equal(t, tC.err, actual)
		})
	}
}

func TestLunaCheck(t *testing.T) {
	testCases := []struct {
		name     string
		number   string
		expected bool
	}{
		{
			name:     "valid number 1",
			number:   "4111111111111111",
			expected: true,
		},
		{
			name:     "valid number 2",
			number:   "4012888888881881",
			expected: true,
		},
		{
			name:     "valid number 3",
			number:   "378282246310005",
			expected: true,
		},
		{
			name:     "invalid number 1",
			number:   "4111111111111113",
			expected: false,
		},
		{
			name:     "invalid number 2",
			number:   "4111111111111111111",
			expected: false,
		},
		{
			name:     "invalid number 3",
			number:   "123",
			expected: false,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := lunaCheck(tC.number)

			require.Equal(t, tC.expected, actual)
		})
	}
}
