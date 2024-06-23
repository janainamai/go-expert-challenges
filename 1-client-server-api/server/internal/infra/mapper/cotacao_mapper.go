package mapper

import (
	"server/internal/core/domain"
	"server/internal/infra/database/entities"
	"server/internal/infra/rest/dto"
)

func RestToDomain(cotacaoDTO *dto.CotacaoDTO) *domain.Cotacao {
	cotacaoDomain := domain.Cotacao{
		Code:       cotacaoDTO.ContentDTO.Code,
		Codein:     cotacaoDTO.ContentDTO.Codein,
		Name:       cotacaoDTO.ContentDTO.Name,
		High:       cotacaoDTO.ContentDTO.High,
		Low:        cotacaoDTO.ContentDTO.Low,
		VarBid:     cotacaoDTO.ContentDTO.VarBid,
		PctChange:  cotacaoDTO.ContentDTO.PctChange,
		Bid:        cotacaoDTO.ContentDTO.Bid,
		Ask:        cotacaoDTO.ContentDTO.Ask,
		Timestamp:  cotacaoDTO.ContentDTO.Timestamp,
		CreateDate: cotacaoDTO.ContentDTO.CreateDate,
	}

	return &cotacaoDomain
}

func DomainToEntity(cotacao *domain.Cotacao) *entities.CotacaoEntity {
	cotacaoEntity := entities.CotacaoEntity{
		Code:       cotacao.Code,
		Codein:     cotacao.Codein,
		Name:       cotacao.Name,
		High:       cotacao.High,
		Low:        cotacao.Low,
		VarBid:     cotacao.VarBid,
		PctChange:  cotacao.PctChange,
		Bid:        cotacao.Bid,
		Ask:        cotacao.Ask,
		Timestamp:  cotacao.Timestamp,
		CreateDate: cotacao.CreateDate,
	}

	return &cotacaoEntity
}

func EntitiesToDomains(cotacoesEntity []*entities.CotacaoEntity) []*domain.Cotacao {
	var cotacoes []*domain.Cotacao
	for _, cotacaoEntity := range cotacoesEntity {
		cotacaoDomain := EntityToDomain(cotacaoEntity)
		cotacoes = append(cotacoes, cotacaoDomain)
	}

	return cotacoes
}

func EntityToDomain(cotacaoEntity *entities.CotacaoEntity) *domain.Cotacao {
	cotacaoDomain := domain.Cotacao{
		ID:         cotacaoEntity.ID,
		Code:       cotacaoEntity.Code,
		Codein:     cotacaoEntity.Codein,
		Name:       cotacaoEntity.Name,
		High:       cotacaoEntity.High,
		Low:        cotacaoEntity.Low,
		VarBid:     cotacaoEntity.VarBid,
		PctChange:  cotacaoEntity.PctChange,
		Bid:        cotacaoEntity.Bid,
		Ask:        cotacaoEntity.Ask,
		Timestamp:  cotacaoEntity.Timestamp,
		CreateDate: cotacaoEntity.CreateDate,
	}

	return &cotacaoDomain
}
