package domain

import (
	"time"

	"github.com/google/uuid"
)

type RatingHistoric struct {
	ID             uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	BrockerStockId uuid.UUID
	ActionId       uuid.UUID
	FromRatingId   uuid.UUID
	ToRatingId     uuid.UUID
	FromTarget     float64
	ToTarget       float64
	Time           time.Time

	FromRating  Rating      `gorm:"foreignKey:FromRatingId;references:ID"`
	ToRating    Rating      `gorm:"foreignKey:ToRatingId;references:ID"`
	Action      Action      `gorm:"foreignKey:ActionId;references:ID"`
	BrokerStock BrokerStock `gorm:"foreignKey:BrockerStockId;references:ID"`
}
