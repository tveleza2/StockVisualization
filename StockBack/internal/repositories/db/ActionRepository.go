package db

import (
	"stock-app/internal/core/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ActionRepository struct {
	db *gorm.DB
}

func NewActionRepository(db *gorm.DB) *ActionRepository {
	return &ActionRepository{db}
}

func (repository ActionRepository) Create(action *domain.Action) error {
	return repository.db.Create(action).Error
}

func (repository ActionRepository) Find(id uuid.UUID) (*domain.Action, error) {
	action := domain.Action{}
	return &action, repository.db.First(&action, "id = ?", id).Error
}

func (repository ActionRepository) FindByName(name string) (domain.Action, error) {
	var action domain.Action
	err := repository.db.Where("name == ?", name).First(&action).Error
	return action, err
}

func (repository ActionRepository) FindByNames(names *[]string) (*map[string]uuid.UUID, error) {
	var actions []domain.Action
	actionMap := make(map[string]uuid.UUID)
	err := repository.db.Where("name IN ?", names).Find(&actions).Error
	if err == nil {
		for _, action := range actions {
			actionMap[action.Name] = action.ID
		}
	}
	return &actionMap, err
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
