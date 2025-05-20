package ports

import (
	"stock-app/internal/core/domain"
	"time"

	"github.com/google/uuid"
)

type RatingHistoricPort interface {
	Create(rating *domain.RatingHistoric) error
	Find(id uuid.UUID) (*domain.RatingHistoric, error)
	FindOneByBrokerStock(id uuid.UUID) (*domain.RatingHistoric, error)
	FindExistence(brokerStockId uuid.UUID, time time.Time) (*domain.RatingHistoric, error)
	FindAllByStock(brokerStockIds []uuid.UUID) ([]domain.RatingHistoric, error)
	FindAll() ([]domain.RatingHistoric, error)
	Update(rating *domain.RatingHistoric) error
	Delete(id uuid.UUID) error
}
