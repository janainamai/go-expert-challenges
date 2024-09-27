package gateway

import (
	"context"

	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/entity/domain"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/entity/usecase"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/infra/mysql"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/shared/dto"
)

type (
	gatewayList struct {
		mysql mysql.ListOrdersMySQL
	}
)

func NewListOrdersGateway(mysql mysql.ListOrdersMySQL) usecase.ListOrdersInfraInterface {
	return &gatewayList{
		mysql: mysql,
	}
}

func (i *gatewayList) List(ctx context.Context) ([]*domain.Order, *dto.Error) {

	entities, err := i.mysql.List(ctx)
	if err != nil {
		return nil, err
	}

	if len(entities) == 0 {
		return nil, nil
	}

	var orders []*domain.Order
	for _, entity := range entities {
		order := domain.Order{
			ID:         entity.ID,
			Price:      entity.Price,
			Tax:        entity.Tax,
			FinalPrice: entity.FinalPrice,
		}

		orders = append(orders, &order)
	}

	return orders, nil
}
