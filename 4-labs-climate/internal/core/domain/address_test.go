package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAddress(t *testing.T) {
	cases := []struct {
		name    string
		cep     string
		wantCep string
	}{
		{
			name:    "valid cep",
			cep:     "12345678",
			wantCep: "12345678",
		},
		{
			name:    "cep with spaces",
			cep:     " 12345678 ",
			wantCep: "12345678",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			address := NewAddress(c.cep)

			assert.NotNil(t, address)
			assert.Equal(t, c.wantCep, address.Cep)
		})
	}
}

func TestIsValidCEP(t *testing.T) {
	cases := []struct {
		name     string
		cep      string
		expected bool
	}{
		{
			name:     "valid cep",
			cep:      "12345678",
			expected: true,
		},
		{
			name:     "invalid cep with letters",
			cep:      "1234a678",
			expected: false,
		},
		{
			name:     "empty cep",
			cep:      "",
			expected: false,
		},
		{
			name:     "invalid cep less than 8 digits",
			cep:      "1234567",
			expected: false,
		},
		{
			name:     "invalid cep more than 8 digits",
			cep:      "123456789",
			expected: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			address := NewAddress(c.cep)
			assert.Equal(t, c.expected, address.IsValidCEP())
		})
	}
}
