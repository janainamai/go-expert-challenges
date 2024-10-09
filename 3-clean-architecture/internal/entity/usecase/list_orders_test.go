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
	mockListOrdersInfraInterface struct {
		mock.Mock
	}
)

func (m *mockListOrdersInfraInterface) List(ctx context.Context) ([]*domain.Order, *dto.Error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}

		return nil, args.Get(1).(*dto.Error)
	}

	return args.Get(0).([]*domain.Order), nil
}

func TestList_ReturnsOrders_WhenSuccess(t *testing.T) {
	listOrdersInfraInterfaceMock := new(mockListOrdersInfraInterface)
	usecase := NewListOrdersUseCase(listOrdersInfraInterfaceMock)

	ctx := context.Background()
	orderID := uuid.NewString()

	ordersResultMock := []*domain.Order{
		{
			ID:         orderID,
			Price:      100,
			Tax:        10,
			FinalPrice: 110,
		},
	}
	listOrdersInfraInterfaceMock.On("List", ctx).Return(ordersResultMock, nil)

	orders, err := usecase.List(ctx)

	assert.Nil(t, err)
	assert.NotNil(t, orders)
	assert.Equal(t, ordersResultMock, orders)

	listOrdersInfraInterfaceMock.AssertExpectations(t)
}

func TestList_ReturnsNil_WhenNotFoundOrders(t *testing.T) {
	listOrdersInfraInterfaceMock := new(mockListOrdersInfraInterface)
	usecase := NewListOrdersUseCase(listOrdersInfraInterfaceMock)

	ctx := context.Background()

	listOrdersInfraInterfaceMock.On("List", ctx).Return(nil, nil)

	orders, err := usecase.List(ctx)

	assert.Nil(t, err)
	assert.Empty(t, orders)

	listOrdersInfraInterfaceMock.AssertExpectations(t)
}

func TestList_ReturnsError_WhenListFailed(t *testing.T) {
	listOrdersInfraInterfaceMock := new(mockListOrdersInfraInterface)
	usecase := NewListOrdersUseCase(listOrdersInfraInterfaceMock)

	ctx := context.Background()

	errorResult := dto.NewError("list failed", "error")
	listOrdersInfraInterfaceMock.On("List", ctx).Return(nil, errorResult)

	orders, err := usecase.List(ctx)

	assert.Nil(t, orders)
	assert.NotNil(t, err)
	assert.Equal(t, "unexpected_error", err.Title)
	assert.Equal(t, "error listing orders", err.Detail)

	listOrdersInfraInterfaceMock.AssertExpectations(t)
}
