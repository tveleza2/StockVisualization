package domain_test

import (
	"stock-app/internal/core/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStockFields(t *testing.T) {
	id := "AAPL"
	name := "Apple Inc."
	score := 95
	stock := domain.Stock{
		ID:    id,
		Name:  name,
		Score: score,
	}

	assert.Equal(t, id, stock.ID)
	assert.Equal(t, name, stock.Name)
	assert.Equal(t, score, stock.Score)
}
