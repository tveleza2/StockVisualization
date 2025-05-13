package mapper

import (
	"stock-app/internal/core/domain"
	"stock-app/internal/handlers/dto"
)

func ToAction(actionDTO *dto.ActionDTO) domain.Action {
	return domain.Action{
		ID:   actionDTO.ID,
		Name: actionDTO.Name,
	}
}

func FromAction(action *domain.Action) dto.ActionDTO {
	return dto.ActionDTO{
		ID:   action.ID,
		Name: action.Name,
	}
}

func FromActions(actions []domain.Action) []dto.ActionDTO {
	actionsDTO := make([]dto.ActionDTO, len(actions))
	for i, u := range actions {
		actionsDTO[i] = FromAction(&u)
	}
	return actionsDTO
}
