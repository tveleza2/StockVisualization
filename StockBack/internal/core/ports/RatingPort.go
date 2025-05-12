package ports

import (
	"stock-app/internal/core/domain"

	"github.com/google/uuid"
)

type RatingPort interface {
	Create(rating *domain.Rating) error
	Find(id uuid.UUID) (*domain.Rating, error)
	FindAll() ([]domain.Rating, error)
	Update(rating *domain.Rating) error
	Delete(id uuid.UUID) error
}
