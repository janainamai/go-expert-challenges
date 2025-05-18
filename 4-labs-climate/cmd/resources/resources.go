package resources

import (
	"github.com/janainamai/go-expert-challenges/4-labs-climate/internal/core/usecase"
	"github.com/janainamai/go-expert-challenges/4-labs-climate/internal/entrypoint"
	"github.com/janainamai/go-expert-challenges/4-labs-climate/internal/outpoint/gateway"
	"github.com/janainamai/go-expert-challenges/4-labs-climate/internal/outpoint/infra/rest"
	"github.com/janainamai/go-expert-challenges/4-labs-climate/internal/outpoint/infra/rest/viacep"
	"github.com/janainamai/go-expert-challenges/4-labs-climate/internal/outpoint/infra/rest/weather"
)

type (
	Resources struct {
		GetTemperatureHandler entrypoint.GetTemperature
	}
)

func LoadResources(key string) *Resources {

	restConnector := rest.New()

	viacepRest := viacep.New(restConnector)
	weatherRest := weather.New(key, restConnector)

	getTemperatureGateway := gateway.NewGetTemperature(viacepRest, weatherRest)

	getTemperatureUsecase := usecase.NewGetTemperature(getTemperatureGateway)

	return &Resources{
		GetTemperatureHandler: entrypoint.NewGetTemperature(getTemperatureUsecase),
	}
}
