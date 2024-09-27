package domain

import (
	"fmt"

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
		err := dto.InitError().WithDetail("invalid id")
		fmt.Println(err)

		return err
	}

	if o.Price <= 0 {
		return dto.InitError().WithDetail("invalid price")
	}

	if o.Tax <= 0 {
		return dto.InitError().WithDetail("invalid tax")
	}

	return nil
}
