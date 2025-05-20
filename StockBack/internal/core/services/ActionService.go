package services

import (
	"errors"
	"fmt"
	"stock-app/internal/core/domain"
	"stock-app/internal/core/ports"
	"stock-app/internal/handlers/dto"
	"stock-app/internal/handlers/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func validateDTOForCreate(dto *dto.ActionDTO) error {
	if dto.Name == "" {
		return errors.New("action has no name")
	}
	return nil
}

func validateDTOForUpdateOrDelete(dto *dto.ActionDTO) error {
	if dto.Name == "" {
		return errors.New("action has no name")
	}

	if dto.ID == uuid.Nil {
		return errors.New("action has no ID")
	}

	return nil
}

type ActionService struct {
	actionRepository ports.ActionPort
}

func NewActionService(repository ports.ActionPort) *ActionService {
	return &ActionService{repository}
}

func (service ActionService) CreateAction(actionDTO dto.ActionDTO) (dto.ActionDTO, error) {
	err := validateDTOForCreate(&actionDTO)
	if err != nil {
		return actionDTO, fmt.Errorf("validation error: %w", err)
	}
	newAction := mapper.ToAction(&actionDTO)
	err = service.actionRepository.Create(&newAction)
	if err != nil {
		return actionDTO, err
	}
	return mapper.FromAction(&newAction), nil
}

func (service ActionService) ReadAction(id uuid.UUID) (dto.ActionDTO, error) {
	newAction := &domain.Action{}
	newAction, err := service.actionRepository.Find(id)
	if err != nil {
		return mapper.FromAction(&domain.Action{}), err
	}
	return mapper.FromAction(newAction), nil
}

func (service ActionService) ReadActions() ([]dto.ActionDTO, error) {
	actions, err := service.actionRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return mapper.FromActions(actions), nil
}

func (service ActionService) UpdateAction(actionDTO dto.ActionDTO) error {
	err := validateDTOForUpdateOrDelete(&actionDTO)
	if err != nil {
		panic(err.Error())
	}
	_, err = service.actionRepository.Find(actionDTO.ID)
	if err != nil {
		return fmt.Errorf("action with ID %s not found", actionDTO.ID)
	}
	updatedAction := mapper.ToAction(&actionDTO)
	err = service.actionRepository.Update(&updatedAction)
	if err != nil {
		return err
	}
	return nil
}

func (service ActionService) DeleteAction(actionDTO dto.ActionDTO) error {
	validateDTOForUpdateOrDelete(&actionDTO)
	_, err := service.actionRepository.Find(actionDTO.ID)
	if err != nil {
		return err
	}
	err = service.actionRepository.Delete(actionDTO.ID)
	if err != nil {
		return err
	}
	return nil
}

func (service ActionService) FindByName(name string) (domain.Action, error) {
	action, err := service.actionRepository.FindByName(name)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			actionDTO, err := service.CreateAction(dto.ActionDTO{Name: name})
			return mapper.ToAction(&actionDTO), err
		}
		return action, err
	}
	return action, nil
}

func (service ActionService) FindByNames(names *[]string) (*map[string]uuid.UUID, error) {
	actions, err := service.actionRepository.FindByNames(names)
	if err != nil {
		return actions, nil
	}
	return actions, err
}
