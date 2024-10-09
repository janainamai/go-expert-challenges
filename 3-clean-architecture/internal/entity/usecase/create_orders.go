package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/entity/domain"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/shared/dto"
)

type (
	CreateOrderUseCaseInterface interface {
		Create(ctx context.Context, order *domain.Order) (*domain.Order, *dto.Error)
	}

	CreateOrderInfraInterface interface {
		Create(ctx context.Context, order *domain.Order) *dto.Error
	}

	useCaseCreate struct {
		infraInterface CreateOrderInfraInterface
	}
)

func NewCreateOrderUseCase(infra CreateOrderInfraInterface) CreateOrderUseCaseInterface {
	return &useCaseCreate{
		infraInterface: infra,
	}
}

func (u *useCaseCreate) Create(ctx context.Context, order *domain.Order) (*domain.Order, *dto.Error) {
	err := u.validateRequest(order)
	if err != nil {
		return nil, dto.NewError("invalid_request", err.GetDetail())
	}

	order.ID = uuid.NewString()
	order.FinalPrice = order.Price + order.Tax

	err = u.infraInterface.Create(ctx, order)
	if err != nil {
		return nil, dto.NewError("unexpected_error", "error creating order")
	}

	return order, nil
}

func (u *useCaseCreate) validateRequest(order *domain.Order) *dto.Error {
	if order.Price <= 0 {
		return dto.InitError().WithDetail("price must be greater then zero")
	}

	if order.Tax <= 0 {
		return dto.InitError().WithDetail("tax must be greater then zero")
	}

	return nil
}
