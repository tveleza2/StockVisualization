package ports

import (
	"stock-app/internal/core/domain"

	"github.com/google/uuid"
)

type ActionPort interface {
	Create(action *domain.Action) error
	Find(id uuid.UUID) (*domain.Action, error)
	FindByName(name string) (domain.Action, error)
	FindByNames(names *[]string) (*map[string]uuid.UUID, error)
	FindAll() ([]domain.Action, error)
	Update(action *domain.Action) error
	Delete(id uuid.UUID) error
}
