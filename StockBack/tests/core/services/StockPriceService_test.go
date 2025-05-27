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

type mockStockPriceRepo struct {
	prices []domain.StockPrice
}

func (m *mockStockPriceRepo) Create(sp *domain.StockPrice) error {
	m.prices = append(m.prices, *sp)
	return nil
}
func (m *mockStockPriceRepo) Find(id uuid.UUID) (*domain.StockPrice, error) {
	for _, sp := range m.prices {
		if sp.ID == id {
			return &sp, nil
		}
	}
	return nil, errors.New("not found")
}
func (m *mockStockPriceRepo) FindAll() ([]domain.StockPrice, error) { return m.prices, nil }
func (m *mockStockPriceRepo) Update(sp *domain.StockPrice) error    { return nil }
func (m *mockStockPriceRepo) Delete(id uuid.UUID) error             { return nil }

func TestStockPriceService_CreateStockPrice(t *testing.T) {
	repo := &mockStockPriceRepo{}
	service := services.NewStockPriceService(repo)
	dto := dto.StockPriceDTO{StockID: "AAPL", Price: 100.0, Time: time.Now()}
	result, err := service.CreateStockPrice(dto)
	assert.NoError(t, err)
	assert.Equal(t, "AAPL", result.StockID)
}

func TestStockPriceService_CreateStockPrice_ValidationError(t *testing.T) {
	repo := &mockStockPriceRepo{}
	service := services.NewStockPriceService(repo)
	dto := dto.StockPriceDTO{StockID: "", Price: 100.0, Time: time.Now()}
	_, err := service.CreateStockPrice(dto)
	assert.Error(t, err)
}

func TestStockPriceService_ReadStockPrice_Filtered(t *testing.T) {
	repo := &mockStockPriceRepo{}
	service := services.NewStockPriceService(repo)
	now := time.Now()
	repo.Create(&domain.StockPrice{ID: uuid.New(), StockID: "AAPL", Price: 100.0, Time: now})
	repo.Create(&domain.StockPrice{ID: uuid.New(), StockID: "GOOG", Price: 200.0, Time: now})
	prices, err := service.ReadStockPrice("AAPL")
	assert.NoError(t, err)
	assert.Len(t, prices, 1)
	assert.Equal(t, "AAPL", prices[0].StockID)
}

func TestStockPriceService_ReadStockPrices_All(t *testing.T) {
	repo := &mockStockPriceRepo{}
	service := services.NewStockPriceService(repo)
	now := time.Now()
	repo.Create(&domain.StockPrice{ID: uuid.New(), StockID: "AAPL", Price: 100.0, Time: now})
	repo.Create(&domain.StockPrice{ID: uuid.New(), StockID: "GOOG", Price: 200.0, Time: now})
	prices, err := service.ReadStockPrices()
	assert.NoError(t, err)
	assert.Len(t, prices, 2)
}
