package usecase

import (
	"context"

	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/entity/domain"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/shared/dto"
)

type (
	ListOrdersUseCaseInterface interface {
		List(ctx context.Context) ([]*domain.Order, *dto.Error)
	}

	ListOrdersInfraInterface interface {
		List(ctx context.Context) ([]*domain.Order, *dto.Error)
	}

	usecaseList struct {
		infraInterface ListOrdersInfraInterface
	}
)

func NewListOrdersUseCase(infra ListOrdersInfraInterface) ListOrdersUseCaseInterface {
	return &usecaseList{
		infraInterface: infra,
	}
}

func (u *usecaseList) List(ctx context.Context) ([]*domain.Order, *dto.Error) {
	orders, err := u.infraInterface.List(ctx)
	if err != nil {
		return nil, dto.NewError("unexpected_error", "error listing orders")
	}

	if orders == nil {
		return []*domain.Order{}, nil
	}

	return orders, nil
}
