package domain

import "github.com/google/uuid"

type BrokerStock struct {
	ID       uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	BrokerID uuid.UUID
	StockID  string `gorm:"not null"`

	Broker Broker `gorm:"foreignKey:BrokerID;references:ID"`
	Stock  Stock  `gorm:"foreignKey:StockID;references:ID"`
}
