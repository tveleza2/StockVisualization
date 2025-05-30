package mapper

import (
	"stock-app/internal/core/domain"
	"stock-app/internal/handlers/dto"
)

func ToStockPrice(stockPriceDTO *dto.StockPriceDTO) domain.StockPrice {
	return domain.StockPrice{
		StockID: stockPriceDTO.StockID,
		Price:   stockPriceDTO.Price,
		Time:    stockPriceDTO.Time,
	}
}

func FromStockPrice(stockPrice *domain.StockPrice) dto.StockPriceDTO {
	return dto.StockPriceDTO{
		StockID:   stockPrice.StockID,
		Price:     stockPrice.Price,
		Time:      stockPrice.Time,
		StockName: stockPrice.Stock.Name, // Optional if loaded
	}
}

func FromStockPrices(stockPriceList []domain.StockPrice) []dto.StockPriceDTO {
	stockPriceDTO := make([]dto.StockPriceDTO, len(stockPriceList))
	for i, sp := range stockPriceList {
		stockPriceDTO[i] = FromStockPrice(&sp)
	}
	return stockPriceDTO
}
