package services

import (
	"errors"
	"fmt"
	"stock-app/internal/core/domain"
	"stock-app/internal/core/ports"
	"stock-app/internal/handlers/dto"
	"stock-app/internal/handlers/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
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
	brokerService         BrokerService
	stockService          StockService
}

func NewBrokerStockService(repository ports.BrokerStockPort, brokerService BrokerService, stockService StockService) *BrokerStockService {
	return &BrokerStockService{
		brokerStockRepository: repository,
		brokerService:         brokerService,
		stockService:          stockService,
	}
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

func (service BrokerStockService) IdsByStock(stockId string) ([]uuid.UUID, error) {
	brokerStockIds := []uuid.UUID{}
	brokerStocks, err := service.brokerStockRepository.FindAllByStock(stockId)
	if err != nil {
		return nil, err
	}
	for _, brokerStock := range brokerStocks {
		brokerStockIds = append(brokerStockIds, brokerStock.ID)
	}
	return brokerStockIds, nil
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

func (service BrokerStockService) FindByBrokerAndStock(brokerName string, stockId string, stockName string) (domain.BrokerStock, error) {
	broker, err := service.brokerService.FindByName(brokerName)
	if err != nil {
		return domain.BrokerStock{}, err
	}
	stock, err := service.stockService.FindById(stockId, stockName)
	if err != nil {
		return domain.BrokerStock{}, err
	}
	brokerStock, err := service.brokerStockRepository.FindByBrokerAndStock(broker.ID, stock.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			brokerStockDTO, err := service.CreateBrokerStock(dto.BrokerStockDTO{BrokerID: broker.ID, StockID: stockId})
			return mapper.ToBrokerStock(&brokerStockDTO), err
		}
		return brokerStock, err
	}
	return brokerStock, nil
}
func (service BrokerStockService) FindByBrokersAndStock(brokerIds []uuid.UUID, stockIds []string) (*[]domain.BrokerStock, error) {
	brokerStocks, err := service.brokerStockRepository.FindByBrokersAndStock(brokerIds, stockIds)
	if err != nil {
		return brokerStocks, err
	}
	return brokerStocks, nil
}
