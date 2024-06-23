package mapper

import (
	"server/internal/core/domain"
	"server/internal/infra/database/entities"
	"server/internal/infra/rest/dto"
)

func RestToDomain(quotationDTO *dto.QuotationDTO) *domain.Quotation {
	quotationDomain := domain.Quotation{
		Code:       quotationDTO.ContentDTO.Code,
		Codein:     quotationDTO.ContentDTO.Codein,
		Name:       quotationDTO.ContentDTO.Name,
		High:       quotationDTO.ContentDTO.High,
		Low:        quotationDTO.ContentDTO.Low,
		VarBid:     quotationDTO.ContentDTO.VarBid,
		PctChange:  quotationDTO.ContentDTO.PctChange,
		Bid:        quotationDTO.ContentDTO.Bid,
		Ask:        quotationDTO.ContentDTO.Ask,
		Timestamp:  quotationDTO.ContentDTO.Timestamp,
		CreateDate: quotationDTO.ContentDTO.CreateDate,
	}

	return &quotationDomain
}

func DomainToEntity(quotation *domain.Quotation) *entities.QuotationEntity {
	quotationEntity := entities.QuotationEntity{
		Code:       quotation.Code,
		Codein:     quotation.Codein,
		Name:       quotation.Name,
		High:       quotation.High,
		Low:        quotation.Low,
		VarBid:     quotation.VarBid,
		PctChange:  quotation.PctChange,
		Bid:        quotation.Bid,
		Ask:        quotation.Ask,
		Timestamp:  quotation.Timestamp,
		CreateDate: quotation.CreateDate,
	}

	return &quotationEntity
}

func EntitiesToDomains(cotacoesEntity []*entities.QuotationEntity) []*domain.Quotation {
	var cotacoes []*domain.Quotation
	for _, quotationEntity := range cotacoesEntity {
		quotationDomain := EntityToDomain(quotationEntity)
		cotacoes = append(cotacoes, quotationDomain)
	}

	return cotacoes
}

func EntityToDomain(quotationEntity *entities.QuotationEntity) *domain.Quotation {
	quotationDomain := domain.Quotation{
		ID:         quotationEntity.ID,
		Code:       quotationEntity.Code,
		Codein:     quotationEntity.Codein,
		Name:       quotationEntity.Name,
		High:       quotationEntity.High,
		Low:        quotationEntity.Low,
		VarBid:     quotationEntity.VarBid,
		PctChange:  quotationEntity.PctChange,
		Bid:        quotationEntity.Bid,
		Ask:        quotationEntity.Ask,
		Timestamp:  quotationEntity.Timestamp,
		CreateDate: quotationEntity.CreateDate,
	}

	return &quotationDomain
}
