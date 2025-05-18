package usecase

import (
	"errors"

	"github.com/janainamai/go-expert-challenges/4-labs-climate/internal/core/domain"
)

var ErrInvalidCEP = errors.New("invalid CEP")
var ErrCepNotFound = errors.New("cep not found")

type (
	GetTemperature interface {
		GetTemperature(cep string) (*domain.Temperature, error)
	}

	GetTemperatureGateway interface {
		GetAddress(cep string) (*domain.Address, error)
		GetTemperature(address *domain.Address) (*float64, error)
	}

	usecase struct {
		gateway GetTemperatureGateway
	}
)

func NewGetTemperature(gateway GetTemperatureGateway) GetTemperature {
	return &usecase{
		gateway: gateway,
	}
}

func (u *usecase) GetTemperature(cep string) (*domain.Temperature, error) {
	address := domain.NewAddress(cep)
	if !address.IsValidCEP() {
		return nil, ErrInvalidCEP
	}

	address, err := u.gateway.GetAddress(cep)
	if err != nil {
		return nil, err
	}

	if address == nil {
		return nil, ErrCepNotFound
	}

	celsius, err := u.gateway.GetTemperature(address)
	if err != nil {
		return nil, err
	}

	temperature := domain.NewTemperature(*celsius)
	return temperature, nil
}
