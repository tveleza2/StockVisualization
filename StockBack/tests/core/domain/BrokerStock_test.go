package domain_test

import (
	"stock-app/internal/core/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBrokerStockFields(t *testing.T) {
	brokerID := uuid.New()
	stockID := "AAPL"
	brokerStock := domain.BrokerStock{
		ID:       uuid.New(),
		BrokerID: brokerID,
		StockID:  stockID,
	}

	assert.Equal(t, brokerID, brokerStock.BrokerID)
	assert.Equal(t, stockID, brokerStock.StockID)
}
