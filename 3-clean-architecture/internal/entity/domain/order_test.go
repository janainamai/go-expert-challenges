package domain

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidate_ReturnsSuccess(t *testing.T) {
	order := Order{
		ID:         uuid.NewString(),
		Price:      100.50,
		Tax:        45.20,
		FinalPrice: 145.70,
	}

	err := order.Validate()

	assert.Nil(t, err)
}

func TestValidate_ReturnsError_WhenInvalidOrder(t *testing.T) {
	mapInvalidOrders := map[string]Order{
		"id is required": {
			ID:         "",
			Price:      100.50,
			Tax:        45.20,
			FinalPrice: 145.70,
		},
		"price must be greater then zero": {
			ID:         uuid.NewString(),
			Price:      0,
			Tax:        45.20,
			FinalPrice: 145.70,
		},
		"tax must be greater then zero": {
			ID:         uuid.New().String(),
			Price:      100.50,
			Tax:        0,
			FinalPrice: 145.70,
		},
	}

	for expectedError, order := range mapInvalidOrders {
		err := order.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "", err.Title)
		assert.Equal(t, expectedError, err.Detail)
	}
}
