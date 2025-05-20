package ports

import (
	"stock-app/internal/core/domain"

	"github.com/google/uuid"
)

type BrokerStockPort interface {
	Create(brokerStock *domain.BrokerStock) error
	Find(id uuid.UUID) (*domain.BrokerStock, error)
	FindAll() ([]domain.BrokerStock, error)
	FindByBrokerAndStock(brokerId uuid.UUID, stockId string) (domain.BrokerStock, error)
	FindByBrokersAndStock(brokerIds []uuid.UUID, stockIds []string) (*[]domain.BrokerStock, error)
	FindAllByStock(stockId string) ([]domain.BrokerStock, error)
	Update(brokerStock *domain.BrokerStock) error
	Delete(id uuid.UUID) error
}
