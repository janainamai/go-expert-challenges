package mysql

import (
	"context"
	"fmt"

	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/infra/mysql/entity"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/infra/mysql/setup"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/shared/dto"
)

type (
	ListOrdersMySQL interface {
		List(ctx context.Context) ([]*entity.Order, *dto.Error)
	}

	listOrdersMySQL struct {
		mysql *setup.MySQL
	}
)

func NewListOrdersMySQL(mysql *setup.MySQL) ListOrdersMySQL {
	return &listOrdersMySQL{
		mysql: mysql,
	}
}

func (i *listOrdersMySQL) List(ctx context.Context) ([]*entity.Order, *dto.Error) {
	rows, err := i.mysql.DB.Query("SELECT id, price, tax, final_price FROM orders")
	if err != nil {
		return nil, dto.InitError().WithDetail(err.Error())
	}

	var orders []*entity.Order
	for rows.Next() {
		var order entity.Order
		err = rows.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice)
		if err != nil {
			return nil, dto.InitError().WithDetail(fmt.Sprintf("error scanning rows: %s", err.Error()))
		}

		orders = append(orders, &order)
	}

	return orders, nil
}
