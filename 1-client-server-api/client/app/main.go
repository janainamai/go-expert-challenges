package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type CotacaoResponse struct {
	Bid string `json:"bid"`
}

func main() {
	cotacao, err := obterCotacao()
	if err != nil {
		panic(err)
	}

	err = registrarCotacao(cotacao)
	if err != nil {
		panic(err)
	}
}

func obterCotacao() (*CotacaoResponse, error) {
	//ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Printf("Erro ao montar requisição para obter a cotação: %s", err.Error())
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			log.Printf("Erro na comunicação com servidor, tempo de resposta excedido")
			return nil, err
		}

		fmt.Printf("Erro ao obter a cotação: %s", err.Error())
		return nil, err
	}
	defer closeBody(resp)

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Erro ao obter a cotação, status code: %d, body: %s\n", resp.StatusCode, string(body))
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Erro ao converter resultado da cotação: %s", err.Error())
		return nil, err
	}

	var cotacao CotacaoResponse
	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		fmt.Printf("Erro ao ler resultado da cotação: %s", err.Error())
		return nil, err
	}

	return &cotacao, nil
}

func closeBody(resp *http.Response) {
	err := resp.Body.Close()
	if err != nil {
		fmt.Printf("Erro ao fechar conexão com response.body: %s", err.Error())
		return
	}
}

func registrarCotacao(cotacao *CotacaoResponse) error {
	pathFile := "../cotacao.txt"

	if _, err := os.Stat(pathFile); os.IsNotExist(err) {
		pathFile = "./cotacao.txt"
	}

	arquivo, err := os.Open(pathFile)
	if err != nil {
		fmt.Printf("Erro abrir arquivo cotacao.txt para registro: %s", err.Error())
		return err
	}
	defer closeFile(arquivo)

	err = os.WriteFile(pathFile, []byte(fmt.Sprintf("Dólar: %s\n", cotacao.Bid)), os.ModePerm)
	if err != nil {
		fmt.Printf("Erro ao registrar cotação no arquivo cotacao.txt: %s", err.Error())
		return err
	}

	fmt.Printf("Cotação do dólar registrada com sucesso no arquivo cotacao.txt")
	return nil
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Printf("Erro ao fechar conexão com arquivo.txt: %s", err.Error())
		return
	}
}
