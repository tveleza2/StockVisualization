package mapper

import (
	"stock-app/internal/core/domain"
	"stock-app/internal/handlers/dto"
)

func ToRating(ratingDTO *dto.RatingDTO) domain.Rating {
	return domain.Rating{
		ID:   ratingDTO.ID,
		Name: ratingDTO.Name,
	}
}

func FromRating(rating *domain.Rating) dto.RatingDTO {
	return dto.RatingDTO{
		ID:   rating.ID,
		Name: rating.Name,
	}
}

func FromRatings(ratings []domain.Rating) []dto.RatingDTO {
	ratingsDTO := make([]dto.RatingDTO, len(ratings))
	for i, r := range ratings {
		ratingsDTO[i] = FromRating(&r)
	}
	return ratingsDTO
}
