package entrypoint

import (
	"encoding/json"
	"net/http"

	"github.com/janainamai/go-expert-challenges/4-labs-climate/internal/core/usecase"
	"github.com/janainamai/go-expert-challenges/4-labs-climate/internal/entrypoint/response"
)

type (
	GetTemperature struct {
		usecase usecase.GetTemperature
	}

	GetTemperatureResponse struct {
		TempC float64 `json:"temp_C" required:"true"`
		TempF float64 `json:"temp_F" required:"true"`
		TempK float64 `json:"temp_K" required:"true"`
	}
)

func NewGetTemperature(usecase usecase.GetTemperature) GetTemperature {
	return GetTemperature{
		usecase: usecase,
	}
}

func (h *GetTemperature) Get(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("zipcode")
	if cep == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.NewError("zipcode is required"))
		return
	}

	temperature, err := h.usecase.GetTemperature(cep)
	if err != nil {
		if err == usecase.ErrInvalidCEP {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(response.NewError("invalid zipcode"))
			return
		}

		if err == usecase.ErrCepNotFound {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response.NewError("can not find zipcode"))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response.NewError("error getting temperature"))
		return
	}

	temperatureResponse := GetTemperatureResponse{
		TempC: temperature.Celsius(),
		TempF: temperature.Fahrenheit(),
		TempK: temperature.Kelvin(),
	}

	jsonResult, err := json.Marshal(temperatureResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response.NewError("error marshalling response"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResult)
}
