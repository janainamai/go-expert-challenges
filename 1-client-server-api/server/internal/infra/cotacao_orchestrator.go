package infra

import (
	"server/internal/core/domain"
	"server/internal/infra/database"
	"server/internal/infra/mapper"
	"server/internal/infra/rest"
)

type (
	CotacaoOrchestrator interface {
		ObterCotacaoAtual() (*domain.Cotacao, error)
		ObterCotacoesRegistradas() ([]*domain.Cotacao, error)
		SalvarCotacao(cotacao *domain.Cotacao) error
	}

	cotacaoOrchestrator struct {
		cotacaoRepository database.CotacaoRepository
		cotacaoRest       rest.CotacaoRest
	}
)

func NewCotacaoOrchestrator(cotacaoRepository database.CotacaoRepository, cotacaoRest rest.CotacaoRest) CotacaoOrchestrator {
	return &cotacaoOrchestrator{
		cotacaoRepository: cotacaoRepository,
		cotacaoRest:       cotacaoRest,
	}
}

func (c *cotacaoOrchestrator) ObterCotacaoAtual() (*domain.Cotacao, error) {
	cotacaoDTO, err := c.cotacaoRest.ObterCotacao()
	if err != nil {
		return nil, err
	}

	cotacaoDomain := mapper.RestToDomain(cotacaoDTO)

	return cotacaoDomain, nil
}

func (c *cotacaoOrchestrator) ObterCotacoesRegistradas() ([]*domain.Cotacao, error) {
	cotacoesEntity, err := c.cotacaoRepository.ObterCotacoesRegistradas()
	if err != nil {
		return nil, err
	}

	cotacoes := mapper.EntitiesToDomains(cotacoesEntity)
	return cotacoes, nil
}

func (c *cotacaoOrchestrator) SalvarCotacao(cotacao *domain.Cotacao) error {
	cotacaoEntity := mapper.DomainToEntity(cotacao)

	err := c.cotacaoRepository.SalvarCotacao(cotacaoEntity)
	if err != nil {
		return err
	}

	return nil
}
