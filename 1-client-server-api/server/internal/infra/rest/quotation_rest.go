package rest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"server/internal/infra/rest/dto"
	"time"
)

type (
	QuotationRest interface {
		GetQuotation() (*dto.QuotationDTO, error)
	}

	quotationRest struct {
	}
)

func NewQuotationRest() QuotationRest {
	return &quotationRest{}
}

func (r *quotationRest) GetQuotation() (*dto.QuotationDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Erro ao montar requisição para obter a cotação: %s", err.Error()))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			log.Printf("Erro na comunicação com site, tempo de resposta excedido")

			return nil, errors.New(fmt.Sprint("Erro na comunicação com site, tempo de resposta excedido"))
		}

		return nil, errors.New(fmt.Sprintf("Erro na comunicação ao obter a cotação: %s", err.Error()))
	}

	defer func(resp *http.Response) {
		err := closeBody(resp)
		if err != nil {

		}
	}(resp)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Erro ao ler a cotação: %s", err.Error()))
	}

	var quotation dto.QuotationDTO
	err = json.Unmarshal(body, &quotation)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Erro de conversão ao ler a cotação: %s", err.Error()))
	}

	return &quotation, nil
}

func closeBody(resp *http.Response) error {
	err := resp.Body.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("Erro ao fechar conexão com response.body: %s", err.Error()))
	}

	return nil
}
