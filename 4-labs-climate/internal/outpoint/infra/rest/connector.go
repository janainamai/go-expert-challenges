package rest

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type (
	Connector interface {
		GetWithTimeout(url string) ([]byte, error)
	}

	rest struct{}

	response struct {
		Body []byte
	}
)

func New() Connector {
	return &rest{}
}

func (r *rest) GetWithTimeout(url string) ([]byte, error) {
	timeout := 1 * time.Second

	responseChann := make(chan response)

	go r.getWithTimeout(url, responseChann)

	var response response
	select {
	case response = <-responseChann:
	case <-time.After(timeout):
		return nil, fmt.Errorf("response not received within the specified time of 1 second")
	}

	return response.Body, nil
}

func (r *rest) getWithTimeout(url string, responseChann chan response) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	response := response{
		Body: body,
	}

	responseChann <- response
}
