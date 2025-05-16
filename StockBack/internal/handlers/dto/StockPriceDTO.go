package dto

import (
	"time"
)

type StockPriceDTO struct {
	StockID   string    `json:"stock_id"`
	Price     float64   `json:"price"`
	Time      time.Time `json:"time"`
	StockName string    `json:"stock_name,omitempty"` // Optional for additional context
}
