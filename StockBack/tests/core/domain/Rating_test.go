package domain_test

import (
	"stock-app/internal/core/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRatingFields(t *testing.T) {
	id := uuid.New()
	name := "Strong Buy"
	rating := domain.Rating{
		ID:   id,
		Name: name,
	}

	assert.Equal(t, id, rating.ID)
	assert.Equal(t, name, rating.Name)
}
