package dto

import (
	"time"

	"github.com/google/uuid"
)

type StockPriceDTO struct {
	ID        uuid.UUID `json:"id"`
	StockID   string    `json:"stock_id"`
	Price     float64   `json:"price"`
	Time      time.Time `json:"time"`
	StockName string    `json:"stock_name,omitempty"` // Optional for additional context
}
