package rest

import (
	"encoding/json"
	"net/http"

	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/entity/domain"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/entity/usecase"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/infra/rest/shared"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/shared/dto"
)

type (
	CreateOrderRestHandler struct {
		usecase usecase.CreateOrderUseCaseInterface
	}

	CreateOrderInput struct {
		Price float64
		Tax   float64
	}

	CreateOrderOutput struct {
		ID         string
		Price      float64
		Tax        float64
		FinalPrice float64
	}
)

func NewCreateOrderRestHandler(usecase usecase.CreateOrderUseCaseInterface) *CreateOrderRestHandler {
	return &CreateOrderRestHandler{
		usecase: usecase,
	}
}

func (h *CreateOrderRestHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input CreateOrderInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		response := dto.NewError("error decoding json", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&response)
		return
	}

	domain := domain.Order{
		Price: input.Price,
		Tax:   input.Tax,
	}

	order, errDto := h.usecase.Create(r.Context(), &domain)
	if errDto != nil {
		if errDto.GetTitle() == "invalid_request" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&errDto)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errDto)
		return
	}

	output := CreateOrderOutput{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}

	jsonResult, errDto := shared.Encode(output)
	if errDto != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errDto)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResult)
}
