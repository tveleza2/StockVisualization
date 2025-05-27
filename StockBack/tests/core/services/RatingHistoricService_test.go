package services_test

import (
	"errors"
	"stock-app/internal/core/domain"
	"stock-app/internal/core/services"
	"stock-app/internal/handlers/dto"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type mockRatingHistoricRepo struct {
	historics []domain.RatingHistoric
}

func (m *mockRatingHistoricRepo) Create(r *domain.RatingHistoric) error {
	m.historics = append(m.historics, *r)
	return nil
}
func (m *mockRatingHistoricRepo) Find(id uuid.UUID) (*domain.RatingHistoric, error) {
	for _, rh := range m.historics {
		if rh.ID == id {
			return &rh, nil
		}
	}
	return nil, errors.New("not found")
}
func (m *mockRatingHistoricRepo) FindOneByBrokerStock(id uuid.UUID) (*domain.RatingHistoric, error) {
	for _, rh := range m.historics {
		if rh.BrokerStockID == id {
			return &rh, nil
		}
	}
	return nil, errors.New("not found")
}
func (m *mockRatingHistoricRepo) FindExistence(brokerStockId uuid.UUID, t time.Time) (*domain.RatingHistoric, error) {
	for _, rh := range m.historics {
		if rh.BrokerStockID == brokerStockId && rh.Time.Equal(t) {
			return &rh, nil
		}
	}
	return nil, errors.New("not found")
}
func (m *mockRatingHistoricRepo) FindAllByStock(brokerStockIds []uuid.UUID) ([]domain.RatingHistoric, error) {
	var result []domain.RatingHistoric
	for _, rh := range m.historics {
		for _, id := range brokerStockIds {
			if rh.BrokerStockID == id {
				result = append(result, rh)
			}
		}
	}
	return result, nil
}
func (m *mockRatingHistoricRepo) FindRecent(date time.Time) ([]domain.RatingHistoric, error) {
	var result []domain.RatingHistoric
	for _, rh := range m.historics {
		if rh.Time.After(date) {
			result = append(result, rh)
		}
	}
	return result, nil
}
func (m *mockRatingHistoricRepo) FindAll() ([]domain.RatingHistoric, error) { return m.historics, nil }
func (m *mockRatingHistoricRepo) Update(r *domain.RatingHistoric) error     { return nil }
func (m *mockRatingHistoricRepo) Delete(id uuid.UUID) error                 { return nil }

func TestRatingHistoricService_SaveRatingHistoric(t *testing.T) {
	repo := &mockRatingHistoricRepo{}
	service := services.NewRatingHistoricService(repo, nil, nil, nil)
	dto := dto.RatingHistoricDTO{
		ID:            uuid.New(),
		BrokerStockID: uuid.New(),
		Time:          time.Now(),
	}
	err := service.CreateRatingHistoric(dto)
	assert.NoError(t, err)
}

func TestRatingHistoricService_SaveRatingHistoric_ValidationError(t *testing.T) {
	repo := &mockRatingHistoricRepo{}
	service := services.NewRatingHistoricService(repo, nil, nil, nil)
	dto := dto.RatingHistoricDTO{} // Missing required fields
	err := service.SaveRatingHistoric(dto)
	assert.Error(t, err)
}

func TestRatingHistoricService_ReadRatingHistoric_Success(t *testing.T) {
	repo := &mockRatingHistoricRepo{}
	service := services.NewRatingHistoricService(repo, nil, nil, nil)
	id := uuid.New()
	dto := dto.RatingHistoricDTO{
		ID:            id,
		BrokerStockID: uuid.New(),
		Time:          time.Now(),
	}
	service.SaveRatingHistoric(dto)
	found, err := service.ReadRatingHistoric(id)
	assert.NoError(t, err)
	assert.Equal(t, id, found.ID)
}

func TestRatingHistoricService_ReadRatingHistoric_NotFound(t *testing.T) {
	repo := &mockRatingHistoricRepo{}
	service := services.NewRatingHistoricService(repo, nil, nil, nil)
	_, err := service.ReadRatingHistoric(uuid.New())
	assert.Error(t, err)
}

func TestRatingHistoricService_ReadRatingHistorics(t *testing.T) {
	repo := &mockRatingHistoricRepo{}
	service := services.NewRatingHistoricService(repo, nil, nil, nil)
	service.SaveRatingHistoric(dto.RatingHistoricDTO{ID: uuid.New(), BrokerStockID: uuid.New(), Time: time.Now()})
	service.SaveRatingHistoric(dto.RatingHistoricDTO{ID: uuid.New(), BrokerStockID: uuid.New(), Time: time.Now()})
	historics, err := service.ReadRatingHistorics()
	assert.NoError(t, err)
	assert.Len(t, historics, 2)
}

func TestRatingHistoricService_UpdateRatingHistoric(t *testing.T) {
	repo := &mockRatingHistoricRepo{}
	service := services.NewRatingHistoricService(repo, nil, nil, nil)
	id := uuid.New()
	dto := dto.RatingHistoricDTO{ID: id, BrokerStockID: uuid.New(), Time: time.Now()}
	service.SaveRatingHistoric(dto)
	updateDTO := dto.RatingHistoricDTO{ID: id, BrokerStockID: uuid.New(), Time: time.Now()}
	err := service.UpdateRatingHistoric(updateDTO)
	assert.NoError(t, err)
}

func TestRatingHistoricService_UpdateRatingHistoric_NotFound(t *testing.T) {
	repo := &mockRatingHistoricRepo{}
	service := services.NewRatingHistoricService(repo, nil, nil, nil)
	updateDTO := dto.RatingHistoricDTO{ID: uuid.New(), BrokerStockID: uuid.New(), Time: time.Now()}
	err := service.UpdateRatingHistoric(updateDTO)
	assert.Error(t, err)
}

func TestRatingHistoricService_DeleteRatingHistoric(t *testing.T) {
	repo := &mockRatingHistoricRepo{}
	service := services.NewRatingHistoricService(repo, nil, nil, nil)
	id := uuid.New()
	dto := dto.RatingHistoricDTO{ID: id, BrokerStockID: uuid.New(), Time: time.Now()}
	service.SaveRatingHistoric(dto)
	err := service.DeleteRatingHistoric(id)
	assert.NoError(t, err)
}

func TestRatingHistoricService_DeleteRatingHistoric_NotFound(t *testing.T) {
	repo := &mockRatingHistoricRepo{}
	service := services.NewRatingHistoricService(repo, nil, nil, nil)
	err := service.DeleteRatingHistoric(uuid.New())
	assert.Error(t, err)
}
