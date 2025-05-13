package dto

import "github.com/google/uuid"

type RatingDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
