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
	gatewaySave struct {
		mysql mysql.SaveOrderMySQL
	}
)

func NewSaveOrderGateway(mysql mysql.SaveOrderMySQL) usecase.SaveOrderInfraInterface {
	return &gatewaySave{
		mysql: mysql,
	}
}

func (i *gatewaySave) Save(ctx context.Context, order *domain.Order) *dto.Error {

	entity := entity.Order{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}
	return i.mysql.Save(ctx, &entity)
}
