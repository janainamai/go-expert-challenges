package weather

import (
	"errors"
	"testing"

	"github.com/janainamai/go-expert-challenges/4-labs-climate/internal/core/domain"
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

func TestGetTemperature(t *testing.T) {
	cases := []struct {
		name       string
		address    *domain.Address
		key        string
		setup      func(mockConn *mockConnector)
		validation func(t *testing.T, response *Response, err error)
	}{
		{
			name: "valid address",
			address: &domain.Address{
				Localidade: "São Paulo",
			},
			key: "123",
			setup: func(mockConn *mockConnector) {
				mockResponse := `{
					"location": {
						"name": "São Paulo"
					},
					"current": {
						"temp_c": 25.0
					}
				}`
				mockConn.On("GetWithTimeout", "https://api.weatherapi.com/v1/current.json?key=123&q=São Paulo").Return([]byte(mockResponse), nil)
			},
			validation: func(t *testing.T, response *Response, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, "São Paulo", response.Location.Name)
				assert.Equal(t, 25.0, response.Current.TempC)
			},
		},
		{
			name: "error getting temperature",
			address: &domain.Address{
				Localidade: "São Paulo",
			},
			key: "123",
			setup: func(mockConn *mockConnector) {
				mockError := errors.New("error")
				mockConn.On("GetWithTimeout", "https://api.weatherapi.com/v1/current.json?key=123&q=São Paulo").Return(nil, mockError)
			},
			validation: func(t *testing.T, response *Response, err error) {
				assert.Error(t, err)
				assert.Nil(t, response)
				assert.EqualError(t, err, "error")
			},
		},
		{
			name: "error unmarshalling response",
			address: &domain.Address{
				Localidade: "São Paulo",
			},
			key: "123",
			setup: func(mockConn *mockConnector) {
				mockResponse := `\`
				mockConn.On("GetWithTimeout", "https://api.weatherapi.com/v1/current.json?key=123&q=São Paulo").Return([]byte(mockResponse), nil)
			},
			validation: func(t *testing.T, response *Response, err error) {
				assert.Error(t, err)
				assert.Nil(t, response)
				assert.Contains(t, err.Error(), "error unmarshalling weather response")
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockConn := new(mockConnector)
			c.setup(mockConn)

			w := New(c.key, mockConn)
			response, err := w.GetTemperature(c.address)

			c.validation(t, response, err)

			mockConn.AssertExpectations(t)
		})
	}
}
