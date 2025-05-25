package services_test

import (
	"errors"
	"stock-app/internal/core/domain"
	"stock-app/internal/core/services"
	"stock-app/internal/handlers/dto"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type mockActionRepo struct {
	actions []domain.Action
}

func (m *mockActionRepo) Create(action *domain.Action) error {
	if action.Name == "fail" {
		return errors.New("create failed")
	}
	m.actions = append(m.actions, *action)
	return nil
}
func (m *mockActionRepo) Find(id uuid.UUID) (*domain.Action, error) {
	for _, a := range m.actions {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("not found")
}
func (m *mockActionRepo) FindByName(name string) (domain.Action, error) {
	for _, a := range m.actions {
		if a.Name == name {
			return a, nil
		}
	}
	return domain.Action{}, errors.New("not found")
}
func (m *mockActionRepo) FindByNames(names *[]string) (*map[string]uuid.UUID, error) {
	result := make(map[string]uuid.UUID)
	for _, n := range *names {
		for _, a := range m.actions {
			if a.Name == n {
				result[n] = a.ID
			}
		}
	}
	return &result, nil
}
func (m *mockActionRepo) FindAll() ([]domain.Action, error) { return m.actions, nil }
func (m *mockActionRepo) Update(action *domain.Action) error {
	for i, a := range m.actions {
		if a.ID == action.ID {
			m.actions[i] = *action
			return nil
		}
	}
	return errors.New("not found")
}
func (m *mockActionRepo) Delete(id uuid.UUID) error {
	for i, a := range m.actions {
		if a.ID == id {
			m.actions = append(m.actions[:i], m.actions[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}

func TestActionService_CreateAction(t *testing.T) {
	repo := &mockActionRepo{}
	service := services.NewActionService(repo)
	dto := dto.ActionDTO{Name: "Buy"}
	result, err := service.CreateAction(dto)
	assert.NoError(t, err)
	assert.Equal(t, "Buy", result.Name)
}

func TestActionService_CreateAction_ValidationError(t *testing.T) {
	repo := &mockActionRepo{}
	service := services.NewActionService(repo)
	dto := dto.ActionDTO{Name: ""}
	_, err := service.CreateAction(dto)
	assert.Error(t, err)
}

func TestActionService_CreateAction_RepoError(t *testing.T) {
	repo := &mockActionRepo{}
	service := services.NewActionService(repo)
	dto := dto.ActionDTO{Name: "fail"}
	_, err := service.CreateAction(dto)
	assert.Error(t, err)
}

func TestActionService_ReadAction_Success(t *testing.T) {
	repo := &mockActionRepo{}
	service := services.NewActionService(repo)
	dto := dto.ActionDTO{Name: "Sell"}
	created, _ := service.CreateAction(dto)
	found, err := service.ReadAction(created.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Sell", found.Name)
}

func TestActionService_ReadAction_NotFound(t *testing.T) {
	repo := &mockActionRepo{}
	service := services.NewActionService(repo)
	_, err := service.ReadAction(uuid.New())
	assert.Error(t, err)
}

func TestActionService_ReadActions(t *testing.T) {
	repo := &mockActionRepo{}
	service := services.NewActionService(repo)
	service.CreateAction(dto.ActionDTO{Name: "Buy"})
	service.CreateAction(dto.ActionDTO{Name: "Sell"})
	actions, err := service.ReadActions()
	assert.NoError(t, err)
	assert.Len(t, actions, 2)
}

func TestActionService_UpdateAction_Success(t *testing.T) {
	repo := &mockActionRepo{}
	service := services.NewActionService(repo)
	created, _ := service.CreateAction(dto.ActionDTO{ID: uuid.New(), Name: "Buy"})
	updateDTO := dto.ActionDTO{ID: created.ID, Name: "Strong Buy"}
	err := service.UpdateAction(updateDTO)
	assert.NoError(t, err)
	updated, _ := service.ReadAction(created.ID)
	assert.Equal(t, "Strong Buy", updated.Name)
}

func TestActionService_UpdateAction_ValidationError(t *testing.T) {
	repo := &mockActionRepo{}
	service := services.NewActionService(repo)
	updateDTO := dto.ActionDTO{ID: uuid.New(), Name: ""}
	assert.Panics(t, func() {
		_ = service.UpdateAction(updateDTO)
	})
}

func TestActionService_UpdateAction_NotFound(t *testing.T) {
	repo := &mockActionRepo{}
	service := services.NewActionService(repo)
	updateDTO := dto.ActionDTO{ID: uuid.New(), Name: "Update"}
	err := service.UpdateAction(updateDTO)
	assert.Error(t, err)
}

func TestActionService_DeleteAction_Success(t *testing.T) {
	repo := &mockActionRepo{}
	service := services.NewActionService(repo)
	created, _ := service.CreateAction(dto.ActionDTO{Name: "Buy"})
	deleteDTO := dto.ActionDTO{ID: created.ID, Name: created.Name}
	err := service.DeleteAction(deleteDTO)
	assert.NoError(t, err)
	_, err = service.ReadAction(created.ID)
	assert.Error(t, err)
}

func TestActionService_DeleteAction_NotFound(t *testing.T) {
	repo := &mockActionRepo{}
	service := services.NewActionService(repo)
	deleteDTO := dto.ActionDTO{ID: uuid.New(), Name: "Nonexistent"}
	err := service.DeleteAction(deleteDTO)
	assert.Error(t, err)
}

func TestActionService_FindByName_Existing(t *testing.T) {
	repo := &mockActionRepo{}
	service := services.NewActionService(repo)
	service.CreateAction(dto.ActionDTO{Name: "Buy"})
	action, err := service.FindByName("Buy")
	assert.NoError(t, err)
	assert.Equal(t, "Buy", action.Name)
}

func TestActionService_FindByName_NotExisting(t *testing.T) {
	repo := &mockActionRepo{}
	service := services.NewActionService(repo)
	_, err := service.FindByName("CreateMe")
	assert.Error(t, err)
}

func TestActionService_FindByNames(t *testing.T) {
	repo := &mockActionRepo{}
	service := services.NewActionService(repo)
	service.CreateAction(dto.ActionDTO{Name: "Buy"})
	service.CreateAction(dto.ActionDTO{Name: "Sell"})
	names := []string{"Buy", "Sell"}
	result, err := service.FindByNames(&names)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(*result))
}
