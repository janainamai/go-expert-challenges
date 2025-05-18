package viacep

import (
	"encoding/json"
	"fmt"

	"github.com/janainamai/go-expert-challenges/4-labs-climate/internal/outpoint/infra/rest"
	"github.com/sirupsen/logrus"
)

type (
	ViaCEPRest interface {
		GetAddress(cep string) (*Response, error)
	}

	viaCepRest struct {
		connector rest.Connector
	}

	Response struct {
		Cep         string `json:"cep"`
		Logradouro  string `json:"logradouro"`
		Complemento string `json:"complemento"`
		Unidade     string `json:"unidade"`
		Bairro      string `json:"bairro"`
		Localidade  string `json:"localidade"`
		Uf          string `json:"uf"`
		Estado      string `json:"estado"`
		Regiao      string `json:"regiao"`
		Ibge        string `json:"ibge"`
		Gia         string `json:"gia"`
		Ddd         string `json:"ddd"`
		Siafi       string `json:"siafi"`
	}
)

func New(connector rest.Connector) ViaCEPRest {
	return &viaCepRest{
		connector: connector,
	}
}

func (v *viaCepRest) GetAddress(cep string) (*Response, error) {
	body, err := v.connector.GetWithTimeout(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		return nil, err
	}

	var res Response
	if err := json.Unmarshal(body, &res); err != nil {
		logrus.Errorf("error unmarshalling viacep response: %v", err)
		return nil, fmt.Errorf("error unmarshalling viacep response: %w", err)
	}

	return &res, nil
}
