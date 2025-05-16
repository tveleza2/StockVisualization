package services

import (
	"errors"
	"fmt"
	"stock-app/internal/core/ports"
	"stock-app/internal/handlers/dto"
	"stock-app/internal/handlers/mapper"
)

func validateStockPriceDTOForCreate(dto *dto.StockPriceDTO) error {
	if dto.StockID == "" {
		return errors.New("stock price has no stock ID")
	}
	return nil
}

type StockPriceService struct {
	stockPriceRepository ports.StockPricePort
}

func NewStockPriceService(repository ports.StockPricePort) *StockPriceService {
	return &StockPriceService{repository}
}

func (service StockPriceService) CreateStockPrice(stockPriceDTO dto.StockPriceDTO) (dto.StockPriceDTO, error) {
	err := validateStockPriceDTOForCreate(&stockPriceDTO)
	if err != nil {
		return stockPriceDTO, fmt.Errorf("validation error: %w", err)
	}
	newStockPrice := mapper.ToStockPrice(&stockPriceDTO)
	err = service.stockPriceRepository.Create(&newStockPrice)
	if err != nil {
		return stockPriceDTO, err
	}
	return mapper.FromStockPrice(&newStockPrice), nil
}

func (service StockPriceService) ReadStockPrice(stockID string) ([]dto.StockPriceDTO, error) {
	allPrices, err := service.stockPriceRepository.FindAll()
	if err != nil {
		return nil, err
	}
	// Filter by StockID
	var filtered []dto.StockPriceDTO
	for _, price := range allPrices {
		if price.StockID == stockID {
			filtered = append(filtered, mapper.FromStockPrice(&price))
		}
	}
	return filtered, nil
}

func (service StockPriceService) ReadStockPrices() ([]dto.StockPriceDTO, error) {
	stockPrices, err := service.stockPriceRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return mapper.FromStockPrices(stockPrices), nil
}

// Update and Delete methods can be added if your port and domain support them.
