package entrypoint

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/janainamai/go-expert-challenges/4-labs-climate/internal/core/domain"
	"github.com/janainamai/go-expert-challenges/4-labs-climate/internal/core/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUsecase struct {
	mock.Mock
}

func (m *mockUsecase) GetTemperature(zipcode string) (*domain.Temperature, error) {
	args := m.Called(zipcode)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Temperature), args.Error(1)
}

func TestGetTemperature(t *testing.T) {
	cases := []struct {
		name       string
		zipcode    string
		setup      func(mockGateway *mockUsecase)
		validation func(t *testing.T, body string, statusCode int)
	}{
		{
			name:    "valid zipcode",
			zipcode: "12345678",
			setup: func(mockUsecase *mockUsecase) {
				mockUsecase.On("GetTemperature", "12345678").Return(domain.NewTemperature(25), nil)
			},
			validation: func(t *testing.T, body string, statusCode int) {
				assert.Equal(t, http.StatusOK, statusCode)
				assert.Equal(t, "{\"temp_C\":25,\"temp_F\":77,\"temp_K\":298}", body)
			},
		},
		{
			name:    "empty zipcode",
			zipcode: "",
			setup:   func(mockUsecase *mockUsecase) {},
			validation: func(t *testing.T, body string, statusCode int) {
				assert.Equal(t, http.StatusBadRequest, statusCode)
				assert.Equal(t, "{\"message\":\"zipcode is required\"}\n", body)
			},
		},
		{
			name:    "invalid zipcode",
			zipcode: "1234567",
			setup: func(mockUsecase *mockUsecase) {
				mockUsecase.On("GetTemperature", "1234567").Return(nil, usecase.ErrInvalidCEP)
			},
			validation: func(t *testing.T, body string, statusCode int) {
				assert.Equal(t, http.StatusUnprocessableEntity, statusCode)
				assert.Equal(t, "{\"message\":\"invalid zipcode\"}\n", body)
			},
		},
		{
			name:    "zipcode not found",
			zipcode: "12345678",
			setup: func(mockUsecase *mockUsecase) {
				mockUsecase.On("GetTemperature", "12345678").Return(nil, usecase.ErrCepNotFound)
			},
			validation: func(t *testing.T, body string, statusCode int) {
				assert.Equal(t, http.StatusNotFound, statusCode)
				assert.Equal(t, "{\"message\":\"can not find zipcode\"}\n", body)
			},
		},
		{
			name:    "internal server error",
			zipcode: "12345678",
			setup: func(mockUsecase *mockUsecase) {
				mockError := errors.New("error getting temperature")
				mockUsecase.On("GetTemperature", "12345678").Return(nil, mockError)
			},
			validation: func(t *testing.T, body string, statusCode int) {
				assert.Equal(t, http.StatusInternalServerError, statusCode)
				assert.Equal(t, "{\"message\":\"error getting temperature\"}\n", body)
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockUsecase := new(mockUsecase)
			c.setup(mockUsecase)

			handler := NewGetTemperature(mockUsecase)
			req, err := http.NewRequest(http.MethodGet, "/?zipcode="+c.zipcode, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler.Get(rr, req)

			c.validation(t, rr.Body.String(), rr.Code)
		})
	}
}
