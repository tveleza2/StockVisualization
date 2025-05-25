package domain_test

import (
	"stock-app/internal/core/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestActionFields(t *testing.T) {
	id := uuid.New()
	name := "Buy"
	action := domain.Action{
		ID:   id,
		Name: name,
	}

	assert.Equal(t, id, action.ID)
	assert.Equal(t, name, action.Name)
}
