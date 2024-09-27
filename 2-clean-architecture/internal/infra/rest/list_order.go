package rest

import (
	"encoding/json"
	"net/http"

	"github.com/janainamai/go-expert-challenges/internal/entity/usecase"
	"github.com/janainamai/go-expert-challenges/internal/infra/rest/shared"
)

type (
	ListOrdersRestHandler struct {
		usecase usecase.ListOrdersUseCaseInterface
	}

	ListOrdersOutput struct {
		ID         string
		Price      float64
		Tax        float64
		FinalPrice float64
	}
)

func NewListOrdersRestHandler(usecase usecase.ListOrdersUseCaseInterface) *ListOrdersRestHandler {
	return &ListOrdersRestHandler{
		usecase: usecase,
	}
}

func (h *ListOrdersRestHandler) List(w http.ResponseWriter, r *http.Request) {
	orders, errDto := h.usecase.List(r.Context())
	if errDto != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errDto)
		return
	}

	var ordersOutput []ListOrdersOutput
	for _, order := range orders {
		output := ListOrdersOutput{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}

		ordersOutput = append(ordersOutput, output)
	}

	jsonResult, errDto := shared.Encode(ordersOutput)
	if errDto != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errDto)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResult)
}
