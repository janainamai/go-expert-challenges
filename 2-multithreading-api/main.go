package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Address struct {
	Details    string
	ReceiveURL string
}

func main() {
	cep := "01153000"
	timeout := 1 * time.Second

	brasilapiURL := "https://brasilapi.com.br/api/cep/v1/" + cep
	viacepURL := "http://viacep.com.br/ws/" + cep + "/json"

	addressChann := make(chan Address)

	go getWithTimeout(brasilapiURL, addressChann)
	go getWithTimeout(viacepURL, addressChann)

	select {
	case address := <-addressChann:
		fmt.Printf("Address receive from %s", address.ReceiveURL)
		fmt.Printf("\nAddress: %s", address.Details)
		close(addressChann)
	case <-time.After(timeout):
		close(addressChann)
		panic("Address not received within the specified time of 1 second")
	}
}

func getWithTimeout(url string, addressChann chan Address) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	address := Address{
		Details:    string(body),
		ReceiveURL: url,
	}

	addressChann <- address
}
