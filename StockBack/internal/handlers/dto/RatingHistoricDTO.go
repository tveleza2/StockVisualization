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

type RequestRatingHistoricDTO struct {
	ID         uuid.UUID `json:"id"`
	StockName  string    `json:"stock"`
	BrokerName string    `json:"broker"`
	ActionName string    `json:"action_name"` // Optional for additional context
	FromTarget float64   `json:"from_target"`
	ToTarget   float64   `json:"to_target"`
	Time       time.Time `json:"time"`
	FromRating string    `json:"from_rating_name,omitempty"` // Optional for additional context
	ToRating   string    `json:"to_rating_name,omitempty"`   // Optional for additional context
}
type FullResponseRatingHistoricDTO struct {
	ID         uuid.UUID `json:"id,omitempty"`
	StockID    string    `json:"ticker"`
	StockName  string    `json:"company"`
	BrokerName string    `json:"brokerage"`
	ActionName string    `json:"action"` // Optional for additional context
	FromTarget string    `json:"target_from"`
	ToTarget   string    `json:"target_to"`
	Time       time.Time `json:"time"`
	FromRating string    `json:"rating_from,omitempty"` // Optional for additional context
	ToRating   string    `json:"rating_to,omitempty"`   // Optional for additional context
}

type ApiResponseFromSource struct {
	Items    []FullResponseRatingHistoricDTO `json:"items"`
	NextPage string                          `json:"next_page"`
}

type SummaryResponseRatingHistoricDTO struct {
	ID         uuid.UUID `json:"id"`
	StockName  string    `json:"stock"`
	BrokerName string    `json:"broker"`
	ToTarget   float64   `json:"to_target"`
	ToRating   string    `json:"to_rating_name,omitempty"` // Optional for additional context
	Time       time.Time `json:"time"`
}
