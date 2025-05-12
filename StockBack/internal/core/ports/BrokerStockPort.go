package ports

import (
	"stock-app/internal/core/domain"

	"github.com/google/uuid"
)

type BrokerStockPort interface {
	Create(brokerStock *domain.BrokerStock) error
	Find(id uuid.UUID) (*domain.BrokerStock, error)
	FindAll() ([]domain.BrokerStock, error)
	Update(brokerStock *domain.BrokerStock) error
	Delete(id uuid.UUID) error
}
