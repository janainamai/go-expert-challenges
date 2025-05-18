package gateway

import (
	"github.com/janainamai/go-expert-challenges/4-labs-climate/internal/core/domain"
	"github.com/janainamai/go-expert-challenges/4-labs-climate/internal/core/usecase"
	"github.com/janainamai/go-expert-challenges/4-labs-climate/internal/outpoint/infra/rest/viacep"
	"github.com/janainamai/go-expert-challenges/4-labs-climate/internal/outpoint/infra/rest/weather"
)

type (
	getTemperatureGateway struct {
		viacepRest  viacep.ViaCEPRest
		weatherRest weather.WeatherRest
	}
)

func NewGetTemperature(
	viacepRest viacep.ViaCEPRest,
	weatherRest weather.WeatherRest,
) usecase.GetTemperatureGateway {
	return &getTemperatureGateway{
		viacepRest:  viacepRest,
		weatherRest: weatherRest,
	}
}

func (g *getTemperatureGateway) GetAddress(cep string) (*domain.Address, error) {
	response, err := g.viacepRest.GetAddress(cep)
	if err != nil {
		return nil, err
	}

	if response == nil || response.Cep == "" {
		return nil, nil
	}

	address := domain.Address{
		Cep:         response.Cep,
		Logradouro:  response.Logradouro,
		Complemento: response.Complemento,
		Unidade:     response.Unidade,
		Bairro:      response.Bairro,
		Localidade:  response.Localidade,
		Uf:          response.Uf,
		Estado:      response.Estado,
		Regiao:      response.Regiao,
		Ibge:        response.Ibge,
		Gia:         response.Gia,
		Ddd:         response.Ddd,
		Siafi:       response.Siafi,
	}

	return &address, nil
}

func (g *getTemperatureGateway) GetTemperature(address *domain.Address) (*float64, error) {
	response, err := g.weatherRest.GetTemperature(address)
	if err != nil {
		return nil, err
	}

	return &response.Current.TempC, nil
}
