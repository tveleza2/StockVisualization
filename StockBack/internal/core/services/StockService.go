package services

import (
	"errors"
	"fmt"
	"stock-app/internal/core/domain"
	"stock-app/internal/core/ports"
	"stock-app/internal/handlers/dto"
	"stock-app/internal/handlers/mapper"

	"gorm.io/gorm"
)

func validateStockDTOForCreate(dto *dto.StockDTO) error {
	if dto.ID == "" {
		return errors.New("stock has no ID")
	}
	if dto.Name == "" {
		return errors.New("stock has no name")
	}
	return nil
}

func validateStockDTOForUpdateOrDelete(dto *dto.StockDTO) error {
	if dto.ID == "" {
		return errors.New("stock has no ID")
	}
	if dto.Name == "" {
		return errors.New("stock has no name")
	}
	return nil
}

type StockService struct {
	stockRepository ports.StockPort
}

func NewStockService(repository ports.StockPort) *StockService {
	return &StockService{repository}
}

func (service StockService) CreateStock(stockDTO dto.StockDTO) (dto.StockDTO, error) {
	err := validateStockDTOForCreate(&stockDTO)
	if err != nil {
		return stockDTO, fmt.Errorf("validation error: %w", err)
	}
	newStock := mapper.ToStock(&stockDTO)
	err = service.stockRepository.Create(&newStock)
	if err != nil {
		return stockDTO, err
	}
	return mapper.FromStock(&newStock), nil
}

func (service StockService) ReadStock(id string) (dto.StockDTO, error) {
	stock, err := service.stockRepository.Find(id)
	if err != nil {
		return dto.StockDTO{}, err
	}
	return mapper.FromStock(stock), nil
}

func (service StockService) FindById(id string, name string) (domain.Stock, error) {
	// Softer version of ReadStock, if no value encountered, it creates a new one in the DB
	stock, err := service.stockRepository.Find(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			stockDTO, err := service.CreateStock(dto.StockDTO{ID: id, Name: name})
			return mapper.ToStock(&stockDTO), err
		}
		return domain.Stock{}, err
	}
	return *stock, nil
}

func (service StockService) ReadStocks() ([]dto.StockDTO, error) {
	stocks, err := service.stockRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return mapper.FromStocks(stocks), nil
}

func (service StockService) UpdateStock(stockDTO dto.StockDTO) error {
	err := validateStockDTOForUpdateOrDelete(&stockDTO)
	if err != nil {
		return err
	}
	_, err = service.stockRepository.Find(stockDTO.ID)
	if err != nil {
		return fmt.Errorf("stock with ID %s not found", stockDTO.ID)
	}
	updatedStock := mapper.ToStock(&stockDTO)
	return service.stockRepository.Update(&updatedStock)
}

func (service StockService) DeleteStock(stockDTO dto.StockDTO) error {
	err := validateStockDTOForUpdateOrDelete(&stockDTO)
	if err != nil {
		return err
	}
	return service.stockRepository.Delete(stockDTO.ID)
}

func (service StockService) CalculateStockScores(recentRatings []domain.RatingHistoric) (map[string]dto.StockDTO, error) {
	stockDtos, err := service.ReadStocks()
	for i, stock := range stockDtos {
		stock.Score = 0
		stockDtos[i] = stock
	}
	stockMap := make(map[string]dto.StockDTO)
	for _, dto := range stockDtos {
		stockMap[dto.ID] = dto
	}
	if err != nil {
		return make(map[string]dto.StockDTO), err
	}

	for _, rating := range recentRatings {
		stockID := rating.BrokerStock.StockID
		stock := stockMap[stockID]
		stock.Score = service.CalculateIndividualScore(stock.Score, &rating)
		stockMap[stockID] = stock
	}

	for _, updatedStock := range stockMap {
		err = service.UpdateStock(updatedStock)
		if err != nil {
			return make(map[string]dto.StockDTO), err
		}
	}
	return stockMap, nil
}

func (service StockService) CalculateIndividualScore(score int, rating *domain.RatingHistoric) int {
	if rating.ToRating.Name == "Buy" {
		if rating.FromRating.Name == "Buy" {
			score++
		} else {
			score += 3
		}
	}
	if rating.FromRating.Name == "Buy" && rating.ToRating.Name != "Buy" {
		score -= 2
	}
	if rating.FromTarget < rating.ToTarget {
		score += 2
	}
	return score
}
