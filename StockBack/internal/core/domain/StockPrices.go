package domain

import (
	"time"

	"github.com/google/uuid"
)

type StockPrices struct {
	ID      uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Price   float64
	Time    time.Time
	StockId string
	Stock   Stock `gorm:"foreignKey:StockId;references:ID"`
}
