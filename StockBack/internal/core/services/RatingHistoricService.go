package services

import (
	"errors"
	"fmt"
	"stock-app/internal/core/ports"
	"stock-app/internal/handlers/dto"
	"stock-app/internal/handlers/mapper"

	"github.com/google/uuid"
)

func validateRatingHistoricDTOForCreate(dto *dto.RatingHistoricDTO) error {
	if dto.BrokerStockID == uuid.Nil {
		return errors.New("rating historic has no broker stock ID")
	}
	if dto.ActionID == uuid.Nil {
		return errors.New("rating historic has no action ID")
	}
	return nil
}

func validateRatingHistoricDTOForUpdateOrDelete(dto *dto.RatingHistoricDTO) error {
	if dto.ID == uuid.Nil {
		return errors.New("rating historic has no ID")
	}
	return nil
}

type RatingHistoricService struct {
	ratingHistoricRepository ports.RatingHistoricPort
}

func NewRatingHistoricService(repository ports.RatingHistoricPort) *RatingHistoricService {
	return &RatingHistoricService{repository}
}

func (service RatingHistoricService) CreateRatingHistoric(ratingHistoricDTO dto.RatingHistoricDTO) (dto.RatingHistoricDTO, error) {
	err := validateRatingHistoricDTOForCreate(&ratingHistoricDTO)
	if err != nil {
		return ratingHistoricDTO, fmt.Errorf("validation error: %w", err)
	}
	newRatingHistoric := mapper.ToRatingHistoric(&ratingHistoricDTO)
	err = service.ratingHistoricRepository.Create(&newRatingHistoric)
	if err != nil {
		return ratingHistoricDTO, err
	}
	return mapper.FromRatingHistoric(&newRatingHistoric), nil
}

func (service RatingHistoricService) ReadRatingHistoric(id uuid.UUID) (dto.RatingHistoricDTO, error) {
	ratingHistoric, err := service.ratingHistoricRepository.Find(id)
	if err != nil {
		return dto.RatingHistoricDTO{}, err
	}
	return mapper.FromRatingHistoric(ratingHistoric), nil
}

func (service RatingHistoricService) ReadRatingHistorics() ([]dto.RatingHistoricDTO, error) {
	ratingHistorics, err := service.ratingHistoricRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return mapper.FromRatingHistorics(ratingHistorics), nil
}

func (service RatingHistoricService) UpdateRatingHistoric(ratingHistoricDTO dto.RatingHistoricDTO) error {
	err := validateRatingHistoricDTOForUpdateOrDelete(&ratingHistoricDTO)
	if err != nil {
		return err
	}
	_, err = service.ratingHistoricRepository.Find(ratingHistoricDTO.ID)
	if err != nil {
		return fmt.Errorf("rating historic with ID %s not found", ratingHistoricDTO.ID)
	}
	updatedRatingHistoric := mapper.ToRatingHistoric(&ratingHistoricDTO)
	return service.ratingHistoricRepository.Update(&updatedRatingHistoric)
}

func (service RatingHistoricService) DeleteRatingHistoric(ratingHistoricDTO dto.RatingHistoricDTO) error {
	err := validateRatingHistoricDTOForUpdateOrDelete(&ratingHistoricDTO)
	if err != nil {
		return err
	}
	return service.ratingHistoricRepository.Delete(ratingHistoricDTO.ID)
}
