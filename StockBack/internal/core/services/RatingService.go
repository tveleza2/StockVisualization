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

func validateRatingDTOForCreate(dto *dto.RatingDTO) error {
	if dto.Name == "" {
		return errors.New("rating has no name")
	}
	return nil
}

func validateRatingDTOForUpdateOrDelete(dto *dto.RatingDTO) error {
	if dto.Name == "" {
		return errors.New("rating has no name")
	}
	if dto.ID == uuid.Nil {
		return errors.New("rating has no ID")
	}
	return nil
}

type RatingService struct {
	ratingRepository ports.RatingPort
}

func NewRatingService(repository ports.RatingPort) *RatingService {
	return &RatingService{repository}
}

func (service RatingService) CreateRating(ratingDTO dto.RatingDTO) (dto.RatingDTO, error) {
	err := validateRatingDTOForCreate(&ratingDTO)
	if err != nil {
		return ratingDTO, fmt.Errorf("validation error: %w", err)
	}
	newRating := mapper.ToRating(&ratingDTO)
	err = service.ratingRepository.Create(&newRating)
	if err != nil {
		return ratingDTO, err
	}
	return mapper.FromRating(&newRating), nil
}

func (service RatingService) ReadRating(id uuid.UUID) (dto.RatingDTO, error) {
	rating, err := service.ratingRepository.Find(id)
	if err != nil {
		return dto.RatingDTO{}, err
	}
	return mapper.FromRating(rating), nil
}

func (service RatingService) ReadRatings() ([]dto.RatingDTO, error) {
	ratings, err := service.ratingRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return mapper.FromRatings(ratings), nil
}

func (service RatingService) UpdateRating(ratingDTO dto.RatingDTO) error {
	err := validateRatingDTOForUpdateOrDelete(&ratingDTO)
	if err != nil {
		return err
	}
	_, err = service.ratingRepository.Find(ratingDTO.ID)
	if err != nil {
		return fmt.Errorf("rating with ID %s not found", ratingDTO.ID)
	}
	updatedRating := mapper.ToRating(&ratingDTO)
	return service.ratingRepository.Update(&updatedRating)
}

func (service RatingService) DeleteRating(ratingDTO dto.RatingDTO) error {
	err := validateRatingDTOForUpdateOrDelete(&ratingDTO)
	if err != nil {
		return err
	}
	return service.ratingRepository.Delete(ratingDTO.ID)
}

func (service RatingService) FindByNames(names *[]string) (*[]domain.Rating, error) {
	ratings, err := service.ratingRepository.FindByNames(names)
	if err != nil {
		return ratings, nil
	}
	return ratings, err
}
func (service RatingService) FindByName(name string) (domain.Rating, error) {
	rating, err := service.ratingRepository.FindByName(name)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ratingDTO, err := service.CreateRating(dto.RatingDTO{Name: name})
			return mapper.ToRating(&ratingDTO), err
		}
		return rating, nil
	}
	return rating, err
}
