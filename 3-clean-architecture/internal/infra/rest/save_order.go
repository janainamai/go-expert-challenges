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
	SaveOrderRestHandler struct {
		usecase usecase.SaveOrderUseCaseInterface
	}

	SaveOrderInput struct {
		Price float64
		Tax   float64
	}

	SaveOrderOutput struct {
		ID         string
		Price      float64
		Tax        float64
		FinalPrice float64
	}
)

func NewSaveOrderRestHandler(usecase usecase.SaveOrderUseCaseInterface) *SaveOrderRestHandler {
	return &SaveOrderRestHandler{
		usecase: usecase,
	}
}

func (h *SaveOrderRestHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input SaveOrderInput
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

	order, errDto := h.usecase.Save(r.Context(), &domain)
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

	output := SaveOrderOutput{
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
