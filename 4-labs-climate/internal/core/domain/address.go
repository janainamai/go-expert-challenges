package domain

import (
	"regexp"
	"strings"
)

var onlyNumbersRegex = regexp.MustCompile(`[^0-9]`)

type Address struct {
	Cep         string
	Logradouro  string
	Complemento string
	Unidade     string
	Bairro      string
	Localidade  string
	Uf          string
	Estado      string
	Regiao      string
	Ibge        string
	Gia         string
	Ddd         string
	Siafi       string
}

func NewAddress(cep string) *Address {
	return &Address{
		Cep: strings.TrimSpace(cep),
	}
}

func (c *Address) IsValidCEP() bool {
	c.Cep = onlyNumbersRegex.ReplaceAllString(c.Cep, "")
	if c.Cep == "" {
		return false
	}

	if len(c.Cep) != 8 {
		return false
	}

	return true
}
