package mysql

import (
	"context"

	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/infra/mysql/entity"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/infra/mysql/setup"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/shared/dto"
	"github.com/sirupsen/logrus"
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
	stmt, err := i.mysql.DB.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		logrus.Errorf("error while prepare statement: %s", err.Error())
		return dto.InitError().WithDetail(err.Error())
	}

	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		logrus.Errorf("error while exec insert of order: %s", err.Error())
		return dto.InitError().WithDetail(err.Error())
	}

	return nil
}
