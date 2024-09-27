package gateway

import (
	"context"

	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/entity/domain"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/entity/usecase"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/infra/mysql"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/infra/mysql/entity"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/shared/dto"
)

type (
	gatewayCreate struct {
		saveOrderMySQL mysql.SaveOrderMySQL
	}
)

func NewCreateOrderGateway(saveOrderMySQL mysql.SaveOrderMySQL) usecase.CreateOrderInfraInterface {
	return &gatewayCreate{
		saveOrderMySQL: saveOrderMySQL,
	}
}

func (i *gatewayCreate) Create(ctx context.Context, order *domain.Order) *dto.Error {

	entity := entity.Order{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}
	return i.saveOrderMySQL.Save(ctx, &entity)
}
