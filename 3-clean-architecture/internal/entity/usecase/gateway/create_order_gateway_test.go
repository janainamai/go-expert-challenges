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
	mockSaveOrderMySQL struct {
		mock.Mock
	}
)

func (m *mockSaveOrderMySQL) Save(ctx context.Context, order *entity.Order) *dto.Error {
	args := m.Called(ctx, order)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*dto.Error)
}

func TestSave_ReturnsSuccess(t *testing.T) {
	saveOrderMySQLMock := new(mockSaveOrderMySQL)
	gateway := NewCreateOrderGateway(saveOrderMySQLMock)

	ctx := context.Background()
	order := &domain.Order{
		ID:         uuid.NewString(),
		Price:      100,
		Tax:        10,
		FinalPrice: 110,
	}

	orderEntity := &entity.Order{
		ID:         order.ID,
		Price:      100,
		Tax:        10,
		FinalPrice: 110,
	}
	saveOrderMySQLMock.On("Save", ctx, orderEntity).Return(nil)

	err := gateway.Create(ctx, order)

	assert.Nil(t, err)

	saveOrderMySQLMock.AssertExpectations(t)
}

func TestSave_ReturnsError_WhenSaveFailed(t *testing.T) {
	saveOrderMySQLMock := new(mockSaveOrderMySQL)
	gateway := NewCreateOrderGateway(saveOrderMySQLMock)

	ctx := context.Background()
	order := &domain.Order{
		ID:         uuid.NewString(),
		Price:      100,
		Tax:        10,
		FinalPrice: 110,
	}

	orderEntity := &entity.Order{
		ID:         order.ID,
		Price:      100,
		Tax:        10,
		FinalPrice: 110,
	}
	errorResult := dto.NewError("save failed", "error")
	saveOrderMySQLMock.On("Save", ctx, orderEntity).Return(errorResult)

	err := gateway.Create(ctx, order)

	assert.NotNil(t, err)
	assert.Equal(t, errorResult.Title, err.Title)
	assert.Equal(t, errorResult.Detail, err.Detail)

	saveOrderMySQLMock.AssertExpectations(t)
}
