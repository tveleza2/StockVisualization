package domain

import "github.com/google/uuid"

type BrokerStock struct {
	ID       uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	BrokerId uuid.UUID
	StockId  string `gorm:"not null"`

	Broker Broker `gorm:"foreignKey:BrokerId;references:ID"`
	Stock  Stock  `gorm:"foreignKey:StockId;references:ID"`
}
