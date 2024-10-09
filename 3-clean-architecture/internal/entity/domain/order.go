package domain

import (
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/shared/dto"
)

type (
	Order struct {
		ID         string
		Price      float64
		Tax        float64
		FinalPrice float64
	}
)

func (o *Order) Validate() *dto.Error {
	if o.ID == "" {
		err := dto.InitError().WithDetail("id is required")

		return err
	}

	if o.Price <= 0 {
		return dto.InitError().WithDetail("price must be greater then zero")
	}

	if o.Tax <= 0 {
		return dto.InitError().WithDetail("tax must be greater then zero")
	}

	return nil
}
