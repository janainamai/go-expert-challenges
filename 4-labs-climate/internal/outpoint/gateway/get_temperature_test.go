package gateway

import (
	"testing"

	"github.com/janainamai/go-expert-challenges/4-labs-climate/internal/core/domain"
	"github.com/janainamai/go-expert-challenges/4-labs-climate/internal/outpoint/infra/rest/viacep"
	"github.com/janainamai/go-expert-challenges/4-labs-climate/internal/outpoint/infra/rest/weather"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type (
	mockViaCep struct {
		mock.Mock
	}

	mockWeather struct {
		mock.Mock
	}
)

func (m *mockViaCep) GetAddress(cep string) (*viacep.Response, error) {
	args := m.Called(cep)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	response := args.Get(0).(*viacep.Response)
	return response, args.Error(1)
}

func (m *mockWeather) GetTemperature(address *domain.Address) (*weather.Response, error) {
	args := m.Called(address)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	response := args.Get(0).(*weather.Response)
	return response, args.Error(1)
}

func TestGetAddress(t *testing.T) {
	cases := []struct {
		name       string
		cep        string
		setup      func(mockViaCep *mockViaCep)
		validation func(t *testing.T, address *domain.Address, err error)
	}{
		{
			name: "valid address",
			cep:  "12345678",
			setup: func(mockViaCep *mockViaCep) {
				mockViaCep.On("GetAddress", "12345678").Return(&viacep.Response{
					Cep:         "12345678",
					Logradouro:  "Rua Teste",
					Complemento: "Apto 101",
					Unidade:     "Unidade Teste",
					Bairro:      "Bairro Teste",
					Localidade:  "Cidade Teste",
					Uf:          "UF Teste",
					Estado:      "Estado Teste",
					Regiao:      "Regiao Teste",
					Ibge:        "IBGE Teste",
					Gia:         "GIA Teste",
					Ddd:         "DDD Teste",
					Siafi:       "SIAFI Teste",
				}, nil)
			},
			validation: func(t *testing.T, address *domain.Address, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, address)
				assert.Equal(t, "12345678", address.Cep)
				assert.Equal(t, "Rua Teste", address.Logradouro)
				assert.Equal(t, "Apto 101", address.Complemento)
				assert.Equal(t, "Unidade Teste", address.Unidade)
				assert.Equal(t, "Bairro Teste", address.Bairro)
				assert.Equal(t, "Cidade Teste", address.Localidade)
				assert.Equal(t, "UF Teste", address.Uf)
				assert.Equal(t, "Estado Teste", address.Estado)
				assert.Equal(t, "Regiao Teste", address.Regiao)
				assert.Equal(t, "IBGE Teste", address.Ibge)
				assert.Equal(t, "GIA Teste", address.Gia)
				assert.Equal(t, "DDD Teste", address.Ddd)
				assert.Equal(t, "SIAFI Teste", address.Siafi)
			},
		},
		{
			name: "error getting address",
			cep:  "12345678",
			setup: func(mockViaCep *mockViaCep) {
				mockError := errors.New("error")
				mockViaCep.On("GetAddress", "12345678").Return(nil, mockError)
			},
			validation: func(t *testing.T, address *domain.Address, err error) {
				assert.Error(t, err)
				assert.Nil(t, address)
				assert.Equal(t, "error", err.Error())
			},
		},
		{
			name: "null response from viacep",
			cep:  "12345678",
			setup: func(mockViaCep *mockViaCep) {
				mockViaCep.On("GetAddress", "12345678").Return(nil, nil)
			},
			validation: func(t *testing.T, address *domain.Address, err error) {
				assert.NoError(t, err)
				assert.Nil(t, address)
			},
		},
		{
			name: "empty address from viacep",
			cep:  "12345678",
			setup: func(mockViaCep *mockViaCep) {
				mockViaCep.On("GetAddress", "12345678").Return(&viacep.Response{}, nil)
			},
			validation: func(t *testing.T, address *domain.Address, err error) {
				assert.NoError(t, err)
				assert.Nil(t, address)
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockViaCep := new(mockViaCep)
			c.setup(mockViaCep)

			gateway := NewGetTemperature(mockViaCep, nil)
			address, err := gateway.GetAddress(c.cep)

			c.validation(t, address, err)

			mockViaCep.AssertExpectations(t)
		})
	}
}

func TestGetTemperature(t *testing.T) {
	cases := []struct {
		name       string
		address    *domain.Address
		setup      func(mockWeather *mockWeather)
		validation func(t *testing.T, temperature *float64, err error)
	}{
		{
			name: "valid temperature",
			address: &domain.Address{
				Cep: "12345678",
			},
			setup: func(mockWeather *mockWeather) {
				mockWeather.On("GetTemperature", mock.Anything).Return(&weather.Response{
					Location: weather.Location{
						Name: "City Teste",
					},
					Current: weather.Current{
						TempC: 25.0,
					},
				}, nil)
			},
			validation: func(t *testing.T, temperature *float64, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, temperature)
				assert.Equal(t, 25.0, *temperature)
			},
		},
		{
			name: "error getting temperature",
			address: &domain.Address{
				Cep: "12345678",
			},
			setup: func(mockWeather *mockWeather) {
				mockError := errors.New("error")
				mockWeather.On("GetTemperature", mock.Anything).Return(nil, mockError)
			},
			validation: func(t *testing.T, temperature *float64, err error) {
				assert.Error(t, err)
				assert.Nil(t, temperature)
				assert.Equal(t, "error", err.Error())
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockWeather := new(mockWeather)
			c.setup(mockWeather)

			gateway := NewGetTemperature(nil, mockWeather)
			temperature, err := gateway.GetTemperature(c.address)

			c.validation(t, temperature, err)

			mockWeather.AssertExpectations(t)
		})
	}
}
