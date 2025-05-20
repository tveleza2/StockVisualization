package ports

import (
	"stock-app/internal/core/domain"

	"github.com/google/uuid"
)

type BrokerPort interface {
	Create(broker *domain.Broker) error
	Find(id uuid.UUID) (*domain.Broker, error)
	FindAll() ([]domain.Broker, error)
	FindByName(name string) (domain.Broker, error)
	FindByNames(names *[]string) (*[]domain.Broker, error)
	Update(broker *domain.Broker) error
	Delete(id uuid.UUID) error
}
