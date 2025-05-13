package mapper

import (
	"stock-app/internal/core/domain"
	"stock-app/internal/handlers/dto"
)

func ToBrokerStock(brokerStockDTO *dto.BrokerStockDTO) domain.BrokerStock {
	return domain.BrokerStock{
		ID:       brokerStockDTO.ID,
		BrokerID: brokerStockDTO.BrokerID,
		StockID:  brokerStockDTO.StockID,
	}
}

func FromBrokerStock(brokerStock *domain.BrokerStock) dto.BrokerStockDTO {
	return dto.BrokerStockDTO{
		ID:         brokerStock.ID,
		BrokerID:   brokerStock.BrokerID,
		StockID:    brokerStock.StockID,
		BrokerName: brokerStock.Broker.Name, // Optional if Broker is loaded
		StockName:  brokerStock.Stock.Name,  // Optional if Stock is loaded
	}
}

func FromBrokerStocks(brokerStocks []domain.BrokerStock) []dto.BrokerStockDTO {
	brokerStocksDTO := make([]dto.BrokerStockDTO, len(brokerStocks))
	for i, bs := range brokerStocks {
		brokerStocksDTO[i] = FromBrokerStock(&bs)
	}
	return brokerStocksDTO
}
