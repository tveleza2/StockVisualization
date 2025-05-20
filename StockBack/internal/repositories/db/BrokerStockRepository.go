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

func (repository BrokerStockRepository) FindByBrokerAndStock(brokerId uuid.UUID, stockId string) (domain.BrokerStock, error) {
	var brokerStock domain.BrokerStock
	err := repository.db.Where("broker_id = ? AND stock_id = ?", brokerId, stockId).First(&brokerStock).Error
	return brokerStock, err
}
func (repository BrokerStockRepository) FindByBrokersAndStock(brokerIds []uuid.UUID, stockIds []string) (*[]domain.BrokerStock, error) {
	var brokerStock []domain.BrokerStock
	err := repository.db.Where("broker_id IN ? AND stock_id IN ?", brokerIds, stockIds).First(&brokerStock).Error
	return &brokerStock, err
}

func (repository BrokerStockRepository) FindAllByStock(stockId string) ([]domain.BrokerStock, error) {
	var brokerStocks []domain.BrokerStock
	err := repository.db.Where("stock_id = ?", stockId).Find(&brokerStocks).Error
	return brokerStocks, err
}

func (repository BrokerStockRepository) Update(brokerStock *domain.BrokerStock) error {
	return repository.db.Save(brokerStock).Error
}

func (repository BrokerStockRepository) Delete(id uuid.UUID) error {
	return repository.db.Delete(&domain.BrokerStock{}, id).Error
}
