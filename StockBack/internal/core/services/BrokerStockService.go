package services

import (
	"errors"
	"fmt"
	"stock-app/internal/core/ports"
	"stock-app/internal/handlers/dto"
	"stock-app/internal/handlers/mapper"

	"github.com/google/uuid"
)

func validateBrokerStockDTOForCreate(dto *dto.BrokerStockDTO) error {
	if dto.BrokerID == uuid.Nil {
		return errors.New("broker stock has no broker ID")
	}
	if dto.StockID == "" {
		return errors.New("broker stock has no stock ID")
	}
	return nil
}

func validateBrokerStockDTOForUpdateOrDelete(dto *dto.BrokerStockDTO) error {
	if dto.ID == uuid.Nil {
		return errors.New("broker stock has no ID")
	}
	return nil
}

type BrokerStockService struct {
	brokerStockRepository ports.BrokerStockPort
}

func NewBrokerStockService(repository ports.BrokerStockPort) *BrokerStockService {
	return &BrokerStockService{repository}
}

func (service BrokerStockService) CreateBrokerStock(brokerStockDTO dto.BrokerStockDTO) (dto.BrokerStockDTO, error) {
	err := validateBrokerStockDTOForCreate(&brokerStockDTO)
	if err != nil {
		return brokerStockDTO, fmt.Errorf("validation error: %w", err)
	}
	newBrokerStock := mapper.ToBrokerStock(&brokerStockDTO)
	err = service.brokerStockRepository.Create(&newBrokerStock)
	if err != nil {
		return brokerStockDTO, err
	}
	return mapper.FromBrokerStock(&newBrokerStock), nil
}

func (service BrokerStockService) ReadBrokerStock(id uuid.UUID) (dto.BrokerStockDTO, error) {
	brokerStock, err := service.brokerStockRepository.Find(id)
	if err != nil {
		return dto.BrokerStockDTO{}, err
	}
	return mapper.FromBrokerStock(brokerStock), nil
}

func (service BrokerStockService) ReadBrokerStocks() ([]dto.BrokerStockDTO, error) {
	brokerStocks, err := service.brokerStockRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return mapper.FromBrokerStocks(brokerStocks), nil
}

func (service BrokerStockService) UpdateBrokerStock(brokerStockDTO dto.BrokerStockDTO) error {
	err := validateBrokerStockDTOForUpdateOrDelete(&brokerStockDTO)
	if err != nil {
		return err
	}
	_, err = service.brokerStockRepository.Find(brokerStockDTO.ID)
	if err != nil {
		return fmt.Errorf("broker stock with ID %s not found", brokerStockDTO.ID)
	}
	updatedBrokerStock := mapper.ToBrokerStock(&brokerStockDTO)
	return service.brokerStockRepository.Update(&updatedBrokerStock)
}

func (service BrokerStockService) DeleteBrokerStock(brokerStockDTO dto.BrokerStockDTO) error {
	err := validateBrokerStockDTOForUpdateOrDelete(&brokerStockDTO)
	if err != nil {
		return err
	}
	return service.brokerStockRepository.Delete(brokerStockDTO.ID)
}
