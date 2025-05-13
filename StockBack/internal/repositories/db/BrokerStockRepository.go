package db

import (
	"stock-app/internal/core/domain"
	"stock-app/internal/core/ports"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BrokerStockRepository struct {
	db *gorm.DB
}

func NewBrokerStockRepository(db *gorm.DB) ports.BrokerStockPort {
	return &BrokerStockRepository{db}
}

func (repository BrokerStockRepository) Create(brokerStock *domain.BrokerStock) error {
	return repository.db.Create(brokerStock).Error
}

func (repository BrokerStockRepository) Find(id uuid.UUID) (*domain.BrokerStock, error) {
	brokerStock := domain.BrokerStock{}
	return &brokerStock, repository.db.First(&brokerStock, id).Error
}

func (repository BrokerStockRepository) FindAll() ([]domain.BrokerStock, error) {
	var brokerStocks []domain.BrokerStock
	err := repository.db.Find(&brokerStocks).Error
	return brokerStocks, err
}

func (repository BrokerStockRepository) Update(brokerStock *domain.BrokerStock) error {
	return repository.db.Save(brokerStock).Error
}

func (repository BrokerStockRepository) Delete(id uuid.UUID) error {
	return repository.db.Delete(&domain.BrokerStock{}, id).Error
}
