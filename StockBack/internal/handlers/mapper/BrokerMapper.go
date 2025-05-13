package mapper

import (
	"stock-app/internal/core/domain"
	"stock-app/internal/handlers/dto"
)

func ToBroker(brokerDTO *dto.BrokerDTO) domain.Broker {
	return domain.Broker{
		ID:   brokerDTO.ID,
		Name: brokerDTO.Name,
	}
}

func FromBroker(broker *domain.Broker) dto.BrokerDTO {
	return dto.BrokerDTO{
		ID:   broker.ID,
		Name: broker.Name,
	}
}

func FromBrokers(brokers []domain.Broker) []dto.BrokerDTO {
	brokersDTO := make([]dto.BrokerDTO, len(brokers))
	for i, b := range brokers {
		brokersDTO[i] = FromBroker(&b)
	}
	return brokersDTO
}
