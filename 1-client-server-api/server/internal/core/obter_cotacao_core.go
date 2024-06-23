package core

import (
	"errors"
	"fmt"
	"server/internal/core/domain"
	"server/internal/infra"
)

type (
	CotacaoCore interface {
		ObterCotacaoAtual() (*domain.Cotacao, error)
		ObterCotacoesRegistradas() ([]*domain.Cotacao, error)
	}

	cotacaoCore struct {
		orchestrator infra.CotacaoOrchestrator
	}
)

func NewCotacaoCore(orchestrator infra.CotacaoOrchestrator) CotacaoCore {
	return &cotacaoCore{
		orchestrator: orchestrator,
	}
}

func (c *cotacaoCore) ObterCotacaoAtual() (*domain.Cotacao, error) {
	cotacao, err := c.orchestrator.ObterCotacaoAtual()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Erro ao obter cotação do dolar: %s", err.Error()))
	}

	err = c.orchestrator.SalvarCotacao(cotacao)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Erro ao salvar cotação: %s", err.Error()))
	}

	return cotacao, nil
}

func (c *cotacaoCore) ObterCotacoesRegistradas() ([]*domain.Cotacao, error) {
	cotacoes, err := c.orchestrator.ObterCotacoesRegistradas()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Erro ao obter cotações: %s", err.Error()))
	}

	return cotacoes, nil
}
