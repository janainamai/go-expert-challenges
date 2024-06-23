package database

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"server/internal/infra/database/entities"
	"time"
)

type (
	CotacaoRepository interface {
		SalvarCotacao(cotacao *entities.CotacaoEntity) error
		ObterCotacoesRegistradas() ([]*entities.CotacaoEntity, error)
	}

	cotacaoRepository struct {
	}
)

func NewCotacaoRepository() CotacaoRepository {
	return &cotacaoRepository{}
}

func (r *cotacaoRepository) connect() (*gorm.DB, error) {
	databasePath := "../internal/infra/database/cotacao.db"

	if _, err := os.Stat(databasePath); os.IsNotExist(err) {
		databasePath = "./internal/infra/database/cotacao.db"
	}

	db, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Erro ao iniciar conexão com database: %s", err.Error()))
	}

	err = db.AutoMigrate(&entities.CotacaoEntity{})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Erro ao sincronizar tabelas do database: %s", err.Error()))
	}

	return db, nil
}

func (r *cotacaoRepository) SalvarCotacao(cotacao *entities.CotacaoEntity) error {
	db, err := r.connect()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	err = db.WithContext(ctx).Save(cotacao).Error
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			log.Printf("Erro ao salvar cotação no database, tempo de resposta excedido")

			return errors.New(fmt.Sprint("Tempo de resposta excedido"))
		}

		return err
	}

	return nil
}

func (r *cotacaoRepository) ObterCotacoesRegistradas() ([]*entities.CotacaoEntity, error) {
	db, err := r.connect()
	if err != nil {
		return nil, err
	}

	var cotacoes []*entities.CotacaoEntity
	err = db.Find(&cotacoes).Error
	if err != nil {
		return nil, err
	}

	return cotacoes, nil
}
