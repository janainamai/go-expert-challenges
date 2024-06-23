package infra

import (
	"server/internal/core/domain"
	"server/internal/infra/database"
	"server/internal/infra/mapper"
	"server/internal/infra/rest"
)

type (
	QuotationOrchestrator interface {
		GetCurrentQuotation() (*domain.Quotation, error)
		GetRegisteredQuotations() ([]*domain.Quotation, error)
		SaveQuotation(quotation *domain.Quotation) error
	}

	quotationOrchestrator struct {
		quotationRepository database.QuotationRepository
		quotationRest       rest.QuotationRest
	}
)

func NewQuotationOrchestrator(quotationRepository database.QuotationRepository, quotationRest rest.QuotationRest) QuotationOrchestrator {
	return &quotationOrchestrator{
		quotationRepository: quotationRepository,
		quotationRest:       quotationRest,
	}
}

func (c *quotationOrchestrator) GetCurrentQuotation() (*domain.Quotation, error) {
	quotationDTO, err := c.quotationRest.GetQuotation()
	if err != nil {
		return nil, err
	}

	quotationDomain := mapper.RestToDomain(quotationDTO)

	return quotationDomain, nil
}

func (c *quotationOrchestrator) GetRegisteredQuotations() ([]*domain.Quotation, error) {
	quotationsEntities, err := c.quotationRepository.GetRegisteredQuotations()
	if err != nil {
		return nil, err
	}

	quotationsDomains := mapper.EntitiesToDomains(quotationsEntities)
	return quotationsDomains, nil
}

func (c *quotationOrchestrator) SaveQuotation(quotation *domain.Quotation) error {
	quotationsEntity := mapper.DomainToEntity(quotation)

	err := c.quotationRepository.SaveQuotation(quotationsEntity)
	if err != nil {
		return err
	}

	return nil
}
