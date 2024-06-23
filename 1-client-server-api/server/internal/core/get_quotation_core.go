package core

import (
	"errors"
	"fmt"
	"server/internal/core/domain"
	"server/internal/infra"
)

type (
	QuotationCore interface {
		GetCurrentQuotation() (*domain.Quotation, error)
		GetRegisteredQuotations() ([]*domain.Quotation, error)
	}

	quotationCore struct {
		orchestrator infra.QuotationOrchestrator
	}
)

func NewQuotationCore(orchestrator infra.QuotationOrchestrator) QuotationCore {
	return &quotationCore{
		orchestrator: orchestrator,
	}
}

func (c *quotationCore) GetCurrentQuotation() (*domain.Quotation, error) {
	quotation, err := c.orchestrator.GetCurrentQuotation()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Erro ao obter cotação do dolar: %s", err.Error()))
	}

	err = c.orchestrator.SaveQuotation(quotation)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Erro ao salvar cotação: %s", err.Error()))
	}

	return quotation, nil
}

func (c *quotationCore) GetRegisteredQuotations() ([]*domain.Quotation, error) {
	quotations, err := c.orchestrator.GetRegisteredQuotations()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Erro ao obter cotações: %s", err.Error()))
	}

	return quotations, nil
}
