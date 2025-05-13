package domain

import (
	"time"

	"github.com/google/uuid"
)

type RatingHistoric struct {
	ID            uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	BrokerStockID uuid.UUID
	ActionID      uuid.UUID
	FromRatingID  uuid.UUID
	ToRatingID    uuid.UUID
	FromTarget    float64
	ToTarget      float64
	Time          time.Time

	FromRating  Rating      `gorm:"foreignKey:FromRatingID;references:ID"`
	ToRating    Rating      `gorm:"foreignKey:ToRatingID;references:ID"`
	Action      Action      `gorm:"foreignKey:ActionID;references:ID"`
	BrokerStock BrokerStock `gorm:"foreignKey:BrokerStockID;references:ID"`
}
