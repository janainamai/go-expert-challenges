package viacep

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockConnector struct {
	mock.Mock
}

func (m *mockConnector) GetWithTimeout(url string) ([]byte, error) {
	args := m.Called(url)
	if args.Get(0) != nil {
		return args.Get(0).([]byte), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestGetAddress(t *testing.T) {
	cases := []struct {
		name       string
		cep        string
		setup      func(mockConn *mockConnector)
		validation func(t *testing.T, address *Response, err error)
	}{
		{
			name: "valid address",
			cep:  "12345678",
			setup: func(mockConn *mockConnector) {
				mockResponse := `{
					"cep": "12345678",
					"logradouro": "Rua Teste",
					"complemento": "Apto 101",
					"unidade": "Unidade Teste",
					"bairro": "Bairro Teste",
					"localidade": "Cidade Teste",
					"uf": "UF Teste",
					"estado": "Estado Teste",
					"regiao": "Regiao Teste",
					"ibge": "IBGE Teste",
					"gia": "GIA Teste",
					"ddd": "DDD Teste",
					"siafi": "SIAFI Teste"
				}`
				mockConn.On("GetWithTimeout", "https://viacep.com.br/ws/12345678/json/").Return([]byte(mockResponse), nil)
			},
			validation: func(t *testing.T, address *Response, err error) {
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
			cep:  "87654321",
			setup: func(mockConn *mockConnector) {
				mockError := errors.New("error")
				mockConn.On("GetWithTimeout", "https://viacep.com.br/ws/87654321/json/").Return(nil, mockError)
			},
			validation: func(t *testing.T, address *Response, err error) {
				assert.Error(t, err)
				assert.Nil(t, address)
				assert.Equal(t, "error", err.Error())
			},
		},
		{
			name: "error unmarshalling response",
			cep:  "12345678",
			setup: func(mockConn *mockConnector) {
				mockResponse := `\`
				mockConn.On("GetWithTimeout", "https://viacep.com.br/ws/12345678/json/").Return([]byte(mockResponse), nil)
			},
			validation: func(t *testing.T, address *Response, err error) {
				assert.Error(t, err)
				assert.Nil(t, address)
				assert.Contains(t, err.Error(), "error unmarshalling viacep response")
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockConn := new(mockConnector)
			viaCepRest := New(mockConn)
			c.setup(mockConn)

			address, err := viaCepRest.GetAddress(c.cep)
			c.validation(t, address, err)

			mockConn.AssertExpectations(t)
		})
	}
}
