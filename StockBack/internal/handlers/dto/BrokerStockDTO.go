package dto

import "github.com/google/uuid"

type BrokerStockDTO struct {
	ID         uuid.UUID `json:"id"`
	BrokerID   uuid.UUID `json:"broker_id"`
	StockID    string    `json:"stock_id"`
	BrokerName string    `json:"broker_name,omitempty"` // Optional for additional context
	StockName  string    `json:"stock_name,omitempty"`  // Optional for additional context
}
