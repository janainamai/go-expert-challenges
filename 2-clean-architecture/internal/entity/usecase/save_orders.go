package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/janainamai/go-expert-challenges/internal/entity/domain"
	"github.com/janainamai/go-expert-challenges/internal/shared/dto"
)

type (
	SaveOrderUseCaseInterface interface {
		Save(ctx context.Context, order *domain.Order) (*domain.Order, *dto.Error)
	}

	SaveOrderInfraInterface interface {
		Save(ctx context.Context, order *domain.Order) *dto.Error
	}

	usecaseSave struct {
		infraInterface SaveOrderInfraInterface
	}
)

func NewSaveOrderUseCase(infra SaveOrderInfraInterface) SaveOrderUseCaseInterface {
	return &usecaseSave{
		infraInterface: infra,
	}
}

func (u *usecaseSave) Save(ctx context.Context, order *domain.Order) (*domain.Order, *dto.Error) {
	err := u.validateRequest(order)
	if err != nil {
		return nil, dto.NewError("invalid_request", err.GetDetail())
	}

	order.ID = uuid.NewString()
	order.FinalPrice = order.Price + order.Tax

	err = u.infraInterface.Save(ctx, order)
	if err != nil {
		return nil, dto.NewError("unexpected_error", "error saving order")
	}

	return order, nil
}

func (u *usecaseSave) validateRequest(order *domain.Order) *dto.Error {
	if order.Price <= 0 {
		return dto.InitError().WithDetail("invalid price")
	}

	if order.Tax <= 0 {
		return dto.InitError().WithDetail("invalid tax")
	}

	return nil
}
