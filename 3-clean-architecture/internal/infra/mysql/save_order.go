package mysql

import (
	"context"

	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/infra/mysql/entity"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/infra/mysql/setup"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/shared/dto"
)

type (
	SaveOrderMySQL interface {
		Save(ctx context.Context, order *entity.Order) *dto.Error
	}

	saveOrderMySQL struct {
		mysql *setup.MySQL
	}
)

func NewSaveOrderMySQL(mysql *setup.MySQL) SaveOrderMySQL {
	return &saveOrderMySQL{
		mysql: mysql,
	}
}

func (i *saveOrderMySQL) Save(ctx context.Context, order *entity.Order) *dto.Error {
	stmt, err := i.mysql.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return dto.InitError().WithDetail(err.Error())
	}

	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return dto.InitError().WithDetail(err.Error())
	}

	return nil
}
