package services_test

import (
	"errors"
	"stock-app/internal/core/domain"
	"stock-app/internal/core/services"
	"stock-app/internal/handlers/dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockStockRepo struct {
	stocks []domain.Stock
}

func (m *mockStockRepo) Create(stock *domain.Stock) error {
	m.stocks = append(m.stocks, *stock)
	return nil
}
func (m *mockStockRepo) Find(id string) (*domain.Stock, error) {
	for _, s := range m.stocks {
		if s.ID == id {
			return &s, nil
		}
	}
	return nil, errors.New("not found")
}
func (m *mockStockRepo) FindAll() ([]domain.Stock, error) { return m.stocks, nil }
func (m *mockStockRepo) Update(stock *domain.Stock) error {
	for i, s := range m.stocks {
		if s.ID == stock.ID {
			m.stocks[i] = *stock
			return nil
		}
	}
	return errors.New("not found")
}
func (m *mockStockRepo) Delete(id string) error {
	for i, s := range m.stocks {
		if s.ID == id {
			m.stocks = append(m.stocks[:i], m.stocks[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}

func TestStockService_CreateStock(t *testing.T) {
	repo := &mockStockRepo{}
	service := services.NewStockService(repo)
	dto := dto.StockDTO{ID: "AAPL", Name: "Apple"}
	result, err := service.CreateStock(dto)
	assert.NoError(t, err)
	assert.Equal(t, "AAPL", result.ID)
}

func TestStockService_CreateStock_ValidationError(t *testing.T) {
	repo := &mockStockRepo{}
	service := services.NewStockService(repo)
	dto := dto.StockDTO{ID: "", Name: ""}
	_, err := service.CreateStock(dto)
	assert.Error(t, err)
}

func TestStockService_ReadStock_Success(t *testing.T) {
	repo := &mockStockRepo{}
	service := services.NewStockService(repo)
	dto := dto.StockDTO{ID: "AAPL", Name: "Apple"}
	created, _ := service.CreateStock(dto)
	found, err := service.ReadStock(created.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Apple", found.Name)
}

func TestStockService_ReadStock_NotFound(t *testing.T) {
	repo := &mockStockRepo{}
	service := services.NewStockService(repo)
	_, err := service.ReadStock("GOOG")
	assert.Error(t, err)
}

func TestStockService_ReadStocks(t *testing.T) {
	repo := &mockStockRepo{}
	service := services.NewStockService(repo)
	service.CreateStock(dto.StockDTO{ID: "A", Name: "A"})
	service.CreateStock(dto.StockDTO{ID: "B", Name: "B"})
	stocks, err := service.ReadStocks()
	assert.NoError(t, err)
	assert.Len(t, stocks, 2)
}

func TestStockService_UpdateStock_Success(t *testing.T) {
	repo := &mockStockRepo{}
	service := services.NewStockService(repo)
	created, _ := service.CreateStock(dto.StockDTO{ID: "A", Name: "A"})
	updateDTO := dto.StockDTO{ID: created.ID, Name: "B"}
	err := service.UpdateStock(updateDTO)
	assert.NoError(t, err)
	updated, _ := service.ReadStock(created.ID)
	assert.Equal(t, "B", updated.Name)
}

func TestStockService_UpdateStock_ValidationError(t *testing.T) {
	repo := &mockStockRepo{}
	service := services.NewStockService(repo)
	updateDTO := dto.StockDTO{ID: "", Name: ""}
	err := service.UpdateStock(updateDTO)
	assert.Error(t, err)
}

func TestStockService_UpdateStock_NotFound(t *testing.T) {
	repo := &mockStockRepo{}
	service := services.NewStockService(repo)
	updateDTO := dto.StockDTO{ID: "X", Name: "Update"}
	err := service.UpdateStock(updateDTO)
	assert.Error(t, err)
}

func TestStockService_DeleteStock_Success(t *testing.T) {
	repo := &mockStockRepo{}
	service := services.NewStockService(repo)
	created, _ := service.CreateStock(dto.StockDTO{ID: "A", Name: "A"})
	deleteDTO := dto.StockDTO{ID: created.ID, Name: created.Name}
	err := service.DeleteStock(deleteDTO)
	assert.NoError(t, err)
	_, err = service.ReadStock(created.ID)
	assert.Error(t, err)
}

func TestStockService_DeleteStock_NotFound(t *testing.T) {
	repo := &mockStockRepo{}
	service := services.NewStockService(repo)
	deleteDTO := dto.StockDTO{ID: "X", Name: "Nonexistent"}
	err := service.DeleteStock(deleteDTO)
	assert.Error(t, err)
}
