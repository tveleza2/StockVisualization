package domain_test

import (
	"stock-app/internal/core/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBrokerFields(t *testing.T) {
	id := uuid.New()
	name := "Goldman Sachs"
	broker := domain.Broker{
		ID:   id,
		Name: name,
	}

	assert.Equal(t, id, broker.ID)
	assert.Equal(t, name, broker.Name)
}
