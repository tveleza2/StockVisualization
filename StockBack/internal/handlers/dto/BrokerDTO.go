package dto

import "github.com/google/uuid"

type BrokerDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
