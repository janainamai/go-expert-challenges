package weather

import (
	"encoding/json"
	"fmt"

	"github.com/janainamai/go-expert-challenges/4-labs-climate/internal/core/domain"
	"github.com/janainamai/go-expert-challenges/4-labs-climate/internal/outpoint/infra/rest"
)

type (
	WeatherRest interface {
		GetTemperature(address *domain.Address) (*Response, error)
	}

	weatherRest struct {
		key       string
		connector rest.Connector
	}

	Response struct {
		Location Location `json:"location"`
		Current  Current  `json:"current"`
	}

	Location struct {
		Name string `json:"name"`
	}

	Current struct {
		TempC float64 `json:"temp_c"`
	}
)

func New(key string, connector rest.Connector) WeatherRest {
	return &weatherRest{
		key:       key,
		connector: connector,
	}
}

func (w *weatherRest) GetTemperature(address *domain.Address) (*Response, error) {
	city := address.Localidade
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", w.key, city)

	body, err := w.connector.GetWithTimeout(url)
	if err != nil {
		return nil, err
	}

	var res Response
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, fmt.Errorf("error unmarshalling weather response: %w", err)
	}

	return &res, nil
}
