package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTemperature(t *testing.T) {
	cases := []struct {
		name    string
		celsius float64
		wantC   float64
		wantF   float64
		wantK   float64
	}{
		{
			name:    "zero celsius",
			celsius: 0,
			wantC:   0,
			wantF:   32,
			wantK:   273,
		},
		{
			name:    "positive celsius",
			celsius: 25,
			wantC:   25,
			wantF:   77,
			wantK:   298,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			temp := NewTemperature(c.celsius)

			assert.Equal(t, c.wantC, temp.Celsius())
			assert.Equal(t, c.wantF, temp.Fahrenheit())
			assert.Equal(t, c.wantK, temp.Kelvin())
		})
	}
}

func TestGetTemperature(t *testing.T) {
	cases := []struct {
		name    string
		celsius float64
		wantC   float64
	}{
		{
			name:    "zero celsius",
			celsius: 0,
			wantC:   0,
		},
		{
			name:    "positive celsius",
			celsius: 25,
			wantC:   25,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			temp := GetTemperature(c.celsius)

			assert.Equal(t, c.wantC, temp.Celsius())
		})
	}
}

func TestSetCelsius(t *testing.T) {
	cases := []struct {
		name    string
		celsius float64
		wantC   float64
		wantF   float64
		wantK   float64
	}{
		{
			name:    "zero celsius",
			celsius: 0,
			wantC:   0,
			wantF:   32,
			wantK:   273,
		},
		{
			name:    "positive celsius",
			celsius: 25,
			wantC:   25,
			wantF:   77,
			wantK:   298,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			temp := NewTemperature(0)
			temp.SetCelsius(c.celsius)

			assert.Equal(t, c.wantC, temp.Celsius())
			assert.Equal(t, c.wantF, temp.Fahrenheit())
			assert.Equal(t, c.wantK, temp.Kelvin())
		})
	}
}
