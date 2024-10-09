package gateway

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/entity/domain"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/infra/mysql/entity"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/shared/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type (
	mockListOrdersMySQL struct {
		mock.Mock
	}
)

func (m *mockListOrdersMySQL) List(ctx context.Context) ([]*entity.Order, *dto.Error) {
	args := m.Called(ctx)

	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}

		return nil, args.Get(1).(*dto.Error)
	}

	return args.Get(0).([]*entity.Order), nil
}

func TestList_ReturnsOrders_WhenFoundOrders(t *testing.T) {
	listOrdersMySQLMock := new(mockListOrdersMySQL)
	gateway := NewListOrdersGateway(listOrdersMySQLMock)

	ctx := context.Background()
	orderID := uuid.NewString()

	ordersResultMock := []*entity.Order{
		{
			ID:         orderID,
			Price:      100,
			Tax:        10,
			FinalPrice: 110,
		},
	}
	listOrdersMySQLMock.On("List", ctx).Return(ordersResultMock, nil)

	orders, err := gateway.List(ctx)

	assert.Nil(t, err)
	assert.NotNil(t, orders)
	ordersExpectedResult := []*domain.Order{
		{
			ID:         orderID,
			Price:      100,
			Tax:        10,
			FinalPrice: 110,
		},
	}
	assert.Equal(t, ordersExpectedResult, orders)

	listOrdersMySQLMock.AssertExpectations(t)
}

func TestList_ReturnsNil_WhenNotFoundOrders(t *testing.T) {
	listOrdersMySQLMock := new(mockListOrdersMySQL)
	gateway := NewListOrdersGateway(listOrdersMySQLMock)

	ctx := context.Background()

	listOrdersMySQLMock.On("List", ctx).Return(nil, nil)

	orders, err := gateway.List(ctx)

	assert.Nil(t, err)
	assert.Nil(t, orders)

	listOrdersMySQLMock.AssertExpectations(t)
}

func TestList_ReturnsError_WhenListFailed(t *testing.T) {
	listOrdersMySQLMock := new(mockListOrdersMySQL)
	gateway := NewListOrdersGateway(listOrdersMySQLMock)

	ctx := context.Background()

	errorResult := dto.NewError("list failed", "error")
	listOrdersMySQLMock.On("List", ctx).Return(nil, errorResult)

	orders, err := gateway.List(ctx)

	assert.Nil(t, orders)
	assert.NotNil(t, err)
	assert.Equal(t, errorResult.Title, err.Title)
	assert.Equal(t, errorResult.Detail, err.Detail)

	listOrdersMySQLMock.AssertExpectations(t)
}
