package usecase

import (
	"errors"
	"testing"

	"github.com/janainamai/go-expert-challenges/4-labs-climate/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockGateway struct {
	mock.Mock
}

func (m *mockGateway) GetAddress(cep string) (*domain.Address, error) {
	args := m.Called(cep)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Address), args.Error(1)
}

func (m *mockGateway) GetTemperature(address *domain.Address) (*float64, error) {
	args := m.Called(address)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*float64), args.Error(1)
}

func TestGetTemperature(t *testing.T) {
	cases := []struct {
		name       string
		cep        string
		setup      func(mock *mockGateway)
		validation func(t *testing.T, temperature *domain.Temperature, err error)
	}{
		{
			name: "valid cep",
			cep:  "12345678",
			setup: func(mockGateway *mockGateway) {
				mockGateway.On("GetAddress", "12345678").Return(&domain.Address{
					Cep:        "12345678",
					Localidade: "Sao Paulo",
				}, nil)
				mockGateway.On("GetTemperature", mock.Anything).Return(&[]float64{25.0}[0], nil)
			},
			validation: func(t *testing.T, temperature *domain.Temperature, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, temperature)
				assert.Equal(t, 25.0, temperature.Celsius())
				assert.Equal(t, 77.0, temperature.Fahrenheit())
				assert.Equal(t, 298.0, temperature.Kelvin())
			},
		},
		{
			name:  "invalid cep",
			cep:   "1234567",
			setup: func(mockGateway *mockGateway) {},
			validation: func(t *testing.T, temperature *domain.Temperature, err error) {
				assert.Error(t, err, ErrInvalidCEP)
			},
		},
		{
			name: "error getting cep",
			cep:  "12345678",
			setup: func(mockGateway *mockGateway) {
				mockError := errors.New("error getting cep")
				mockGateway.On("GetAddress", "12345678").Return(nil, mockError)
			},
			validation: func(t *testing.T, temperature *domain.Temperature, err error) {
				assert.Error(t, err)
				assert.Equal(t, "error getting cep", err.Error())
				assert.Nil(t, temperature)
			},
		},
		{
			name: "cep not found",
			cep:  "12345678",
			setup: func(mockGateway *mockGateway) {
				mockGateway.On("GetAddress", "12345678").Return(nil, nil)
			},
			validation: func(t *testing.T, temperature *domain.Temperature, err error) {
				assert.Error(t, err, ErrCepNotFound)
				assert.Nil(t, temperature)
			},
		},
		{
			name: "error getting temperature",
			cep:  "12345678",
			setup: func(mockGateway *mockGateway) {
				mockGateway.On("GetAddress", "12345678").Return(&domain.Address{
					Cep:        "12345678",
					Localidade: "Sao Paulo",
				}, nil)
				mockError := errors.New("error getting temperature")
				mockGateway.On("GetTemperature", mock.Anything).Return(nil, mockError)
			},
			validation: func(t *testing.T, temperature *domain.Temperature, err error) {
				assert.Error(t, err)
				assert.Equal(t, "error getting temperature", err.Error())
				assert.Nil(t, temperature)
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockGateway := new(mockGateway)
			c.setup(mockGateway)

			usecase := NewGetTemperature(mockGateway)
			temperature, err := usecase.GetTemperature(c.cep)

			c.validation(t, temperature, err)

			mockGateway.AssertExpectations(t)
		})
	}

}
