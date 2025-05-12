package ports

import (
	"stock-app/internal/core/domain"

	"github.com/google/uuid"
)

type StockPort interface {
	Create(stock *domain.Stock) error
	Find(id uuid.UUID) (*domain.Stock, error)
	FindAll() ([]domain.Stock, error)
	Update(stock *domain.Stock) error
	Delete(id uuid.UUID) error
}
