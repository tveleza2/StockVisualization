package db

import (
	"stock-app/internal/core/domain"
	"stock-app/internal/core/ports"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ActionRepository struct {
	db *gorm.DB
}

func NewActionRepository(db *gorm.DB) ports.ActionPort {
	return &ActionRepository{db}
}

func (repository ActionRepository) Create(action *domain.Action) error {
	return repository.db.Create(action).Error
}

func (repository ActionRepository) Find(id uuid.UUID) (*domain.Action, error) {
	action := domain.Action{}
	return &action, repository.db.First(&action).Error

}

func (repository ActionRepository) FindAll() ([]domain.Action, error) {
	var actions []domain.Action
	err := repository.db.Find(&actions).Error
	return actions, err
}

func (repository ActionRepository) Update(action *domain.Action) error {
	return repository.db.Save(action).Error
}
func (repository ActionRepository) Delete(id uuid.UUID) error {
	return repository.db.Delete(&domain.Action{}, id).Error
}
