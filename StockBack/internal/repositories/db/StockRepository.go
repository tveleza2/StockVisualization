package db

import (
	"stock-app/internal/core/domain"
	"stock-app/internal/core/ports"

	"gorm.io/gorm"
)

type StockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) ports.StockPort {
	return &StockRepository{db}
}

func (repository StockRepository) Create(stock *domain.Stock) error {
	return repository.db.Create(stock).Error
}

func (repository StockRepository) Find(id string) (*domain.Stock, error) {
	stock := domain.Stock{}
	return &stock, repository.db.Where("id = ?", id).First(&stock).Error
}

func (repository StockRepository) FindAll() ([]domain.Stock, error) {
	var stocks []domain.Stock
	err := repository.db.Find(&stocks).Error
	return stocks, err
}

func (repository StockRepository) Update(stock *domain.Stock) error {
	return repository.db.Save(stock).Error
}

func (repository StockRepository) Delete(id string) error {
	return repository.db.Delete(&domain.Stock{}, id).Error
}
