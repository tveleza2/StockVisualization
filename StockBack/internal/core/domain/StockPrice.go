package domain

import (
	"time"

	"github.com/google/uuid"
)

type StockPrice struct {
	ID      uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Price   float64
	Time    time.Time
	StockID string
	Stock   Stock `gorm:"foreignKey:StockID;references:ID"`
}
