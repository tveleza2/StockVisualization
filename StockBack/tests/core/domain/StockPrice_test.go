package domain_test

import (
	"stock-app/internal/core/domain"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestStockPriceFields(t *testing.T) {
	id := uuid.New()
	stockID := "AAPL"
	price := 189.99
	timestamp := time.Now()

	sp := domain.StockPrice{
		ID:      id,
		StockID: stockID,
		Price:   price,
		Time:    timestamp,
	}

	assert.Equal(t, id, sp.ID)
	assert.Equal(t, stockID, sp.StockID)
	assert.Equal(t, price, sp.Price)
	assert.Equal(t, timestamp, sp.Time)
}
