package domain

import "math"

type Temperature struct {
	fahrenheit float64
	celsius    float64
	kelvin     float64
}

func NewTemperature(celsius float64) *Temperature {
	f := math.Round(((celsius*1.8)+32)*10) / 10
	k := math.Round((celsius+273)*10) / 10
	c := math.Round(celsius*10) / 10

	return &Temperature{
		celsius:    c,
		fahrenheit: f,
		kelvin:     k,
	}
}

func GetTemperature(celsius float64) *Temperature {
	return &Temperature{
		celsius: celsius,
	}
}

func (t *Temperature) Celsius() float64 {
	return t.celsius
}

func (t *Temperature) Fahrenheit() float64 {
	return t.fahrenheit
}

func (t *Temperature) Kelvin() float64 {
	return t.kelvin
}

func (t *Temperature) SetCelsius(celsius float64) {
	t.celsius = celsius
	t.fahrenheit = math.Round(((celsius*1.8)+32)*10) / 10
	t.kelvin = math.Round((celsius+273)*10) / 10
}
