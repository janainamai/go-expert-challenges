package usecase

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/entity/domain"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/shared/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type (
	mockCreateOrderInfraInterface struct {
		mock.Mock
	}
)

func (m *mockCreateOrderInfraInterface) Create(ctx context.Context, order *domain.Order) *dto.Error {
	args := m.Called(ctx, order)
	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(*dto.Error)
}

func TestCreate_WhenSuccess(t *testing.T) {
	createOrderInfraInterfaceMock := new(mockCreateOrderInfraInterface)
	usecase := NewCreateOrderUseCase(createOrderInfraInterfaceMock)

	ctx := context.Background()
	orderInput := domain.Order{
		Price: 1000,
		Tax:   100,
	}

	createOrderInfraInterfaceMock.On("Create", ctx, mock.Anything).Return(nil)

	order, err := usecase.Create(ctx, &orderInput)

	assert.Nil(t, err)
	assert.NotNil(t, order)
	assert.NotNil(t, order.ID)
	assert.Equal(t, 1100.0, order.FinalPrice)

	createOrderInfraInterfaceMock.AssertExpectations(t)
}

func TestCreate_ReturnsError_WhenInvalidRequest(t *testing.T) {
	mapInvalidRequests := map[string]*domain.Order{
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

	for expectedError, orderInput := range mapInvalidRequests {
		createOrderInfraInterfaceMock := new(mockCreateOrderInfraInterface)
		usecase := NewCreateOrderUseCase(createOrderInfraInterfaceMock)

		ctx := context.Background()

		order, err := usecase.Create(ctx, orderInput)

		assert.Nil(t, order)
		assert.NotNil(t, err)
		assert.Equal(t, "invalid_request", err.Title)
		assert.Equal(t, expectedError, err.Detail)

		createOrderInfraInterfaceMock.AssertExpectations(t)
	}
}

func TestCreate_ReturnsError_WhenCreateFailed(t *testing.T) {
	createOrderInfraInterfaceMock := new(mockCreateOrderInfraInterface)
	usecase := NewCreateOrderUseCase(createOrderInfraInterfaceMock)

	ctx := context.Background()
	orderInput := domain.Order{
		Price: 1000,
		Tax:   100,
	}

	errorMock := dto.InitError().WithDetail("timeout")
	createOrderInfraInterfaceMock.On("Create", ctx, mock.Anything).Return(errorMock)

	order, err := usecase.Create(ctx, &orderInput)

	assert.Nil(t, order)
	assert.NotNil(t, err)
	assert.Equal(t, "unexpected_error", err.Title)
	assert.Equal(t, "error creating order", err.Detail)

	createOrderInfraInterfaceMock.AssertExpectations(t)
}
