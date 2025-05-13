package dto

import (
	"time"

	"github.com/google/uuid"
)

type RatingHistoricDTO struct {
	ID             uuid.UUID `json:"id"`
	BrokerStockID  uuid.UUID `json:"broker_stock_id"`
	ActionID       uuid.UUID `json:"action_id"`
	FromRatingID   uuid.UUID `json:"from_rating_id"`
	ToRatingID     uuid.UUID `json:"to_rating_id"`
	FromTarget     float64   `json:"from_target"`
	ToTarget       float64   `json:"to_target"`
	Time           time.Time `json:"time"`
	FromRatingName string    `json:"from_rating_name,omitempty"` // Optional for additional context
	ToRatingName   string    `json:"to_rating_name,omitempty"`   // Optional for additional context
	ActionName     string    `json:"action_name,omitempty"`      // Optional for additional context
}
