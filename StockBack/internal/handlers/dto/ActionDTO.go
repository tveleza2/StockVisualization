package dto

import "github.com/google/uuid"

type ActionDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
