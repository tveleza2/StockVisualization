package domain_test

import (
	"stock-app/internal/core/domain"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRatingHistoricFields(t *testing.T) {
	id := uuid.New()
	brokerStockID := uuid.New()
	fromRatingID := uuid.New()
	toRatingID := uuid.New()
	actionID := uuid.New()
	timestamp := time.Now()

	rh := domain.RatingHistoric{
		ID:            id,
		BrokerStockID: brokerStockID,
		FromRatingID:  fromRatingID,
		ToRatingID:    toRatingID,
		ActionID:      actionID,
		Time:          timestamp,
	}

	assert.Equal(t, id, rh.ID)
	assert.Equal(t, brokerStockID, rh.BrokerStockID)
	assert.Equal(t, fromRatingID, rh.FromRatingID)
	assert.Equal(t, toRatingID, rh.ToRatingID)
	assert.Equal(t, actionID, rh.ActionID)
	assert.Equal(t, timestamp, rh.Time)
}
