package ports

import (
	"stock-app/internal/core/domain"
)

type StockPort interface {
	Create(stock *domain.Stock) error
	Find(id string) (*domain.Stock, error)
	FindAll() ([]domain.Stock, error)
	Update(stock *domain.Stock) error
	Delete(id string) error
}
