package domain

import "github.com/google/uuid"

type Action struct {
	ID   uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name string
}
