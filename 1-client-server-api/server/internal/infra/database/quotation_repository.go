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
	QuotationRepository interface {
		SaveQuotation(quotation *entities.QuotationEntity) error
		GetRegisteredQuotations() ([]*entities.QuotationEntity, error)
	}

	quotationRepository struct {
	}
)

func NewQuotationRepository() QuotationRepository {
	return &quotationRepository{}
}

func (r *quotationRepository) connect() (*gorm.DB, error) {
	databasePath := "../internal/infra/database/quotation.db"

	if _, err := os.Stat(databasePath); os.IsNotExist(err) {
		databasePath = "./internal/infra/database/quotation.db"
	}

	db, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Erro ao iniciar conexão com database: %s", err.Error()))
	}

	err = db.AutoMigrate(&entities.QuotationEntity{})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Erro ao sincronizar tabelas do database: %s", err.Error()))
	}

	return db, nil
}

func (r *quotationRepository) SaveQuotation(quotation *entities.QuotationEntity) error {
	db, err := r.connect()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	err = db.WithContext(ctx).Save(quotation).Error
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			log.Printf("Erro ao salvar cotação no database, tempo de resposta excedido")

			return errors.New(fmt.Sprint("Tempo de resposta excedido"))
		}

		return err
	}

	return nil
}

func (r *quotationRepository) GetRegisteredQuotations() ([]*entities.QuotationEntity, error) {
	db, err := r.connect()
	if err != nil {
		return nil, err
	}

	var cotacoes []*entities.QuotationEntity
	err = db.Find(&cotacoes).Error
	if err != nil {
		return nil, err
	}

	return cotacoes, nil
}
