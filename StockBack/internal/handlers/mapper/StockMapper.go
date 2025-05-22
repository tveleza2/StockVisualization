package mapper

import (
	"stock-app/internal/core/domain"
	"stock-app/internal/handlers/dto"
)

func ToStock(stockDTO *dto.StockDTO) domain.Stock {
	return domain.Stock{
		ID:    stockDTO.ID,
		Name:  stockDTO.Name,
		Score: stockDTO.Score,
	}
}

func FromStock(stock *domain.Stock) dto.StockDTO {
	return dto.StockDTO{
		ID:    stock.ID,
		Name:  stock.Name,
		Score: stock.Score,
	}
}

func FromStocks(stocks []domain.Stock) []dto.StockDTO {
	stocksDTO := make([]dto.StockDTO, len(stocks))
	for i, s := range stocks {
		stocksDTO[i] = FromStock(&s)
	}
	return stocksDTO
}
