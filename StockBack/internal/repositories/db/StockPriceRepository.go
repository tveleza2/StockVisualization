package db

import (
	"stock-app/internal/core/domain"
	"stock-app/internal/core/ports"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StockPriceRepository struct {
	db *gorm.DB
}

func NewStockPriceRepository(db *gorm.DB) ports.StockPricePort {
	return &StockPriceRepository{db}
}

func (repository StockPriceRepository) Create(stockPrice *domain.StockPrice) error {
	return repository.db.Create(stockPrice).Error
}

func (repository StockPriceRepository) Find(id uuid.UUID) (*domain.StockPrice, error) {
	stockPrice := domain.StockPrice{}
	return &stockPrice, repository.db.First(&stockPrice, id).Error
}

func (repository StockPriceRepository) FindAll() ([]domain.StockPrice, error) {
	var stockPrice []domain.StockPrice
	err := repository.db.Find(&stockPrice).Error
	return stockPrice, err
}

func (repository StockPriceRepository) Update(stockPrice *domain.StockPrice) error {
	return repository.db.Save(stockPrice).Error
}

func (repository StockPriceRepository) Delete(id uuid.UUID) error {
	return repository.db.Delete(&domain.StockPrice{}, id).Error
}
