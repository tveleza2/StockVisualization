package ports

import (
	"stock-app/internal/core/domain"

	"github.com/google/uuid"
)

type StockPricePort interface {
	Create(stockPrice *domain.StockPrice) error
	Find(id uuid.UUID) (*domain.StockPrice, error)
	FindAll() ([]domain.StockPrice, error)
	Update(stockPrice *domain.StockPrice) error
	Delete(id uuid.UUID) error
}
