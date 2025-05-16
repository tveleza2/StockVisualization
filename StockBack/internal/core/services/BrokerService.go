package services

import (
	"errors"
	"fmt"
	"stock-app/internal/core/ports"
	"stock-app/internal/handlers/dto"
	"stock-app/internal/handlers/mapper"

	"github.com/google/uuid"
)

func validateBrokerDTOForCreate(dto *dto.BrokerDTO) error {
	if dto.Name == "" {
		return errors.New("broker has no name")
	}
	return nil
}

func validateBrokerDTOForUpdateOrDelete(dto *dto.BrokerDTO) error {
	if dto.Name == "" {
		return errors.New("broker has no name")
	}
	if dto.ID == uuid.Nil {
		return errors.New("broker has no ID")
	}
	return nil
}

type BrokerService struct {
	brokerRepository ports.BrokerPort
}

func NewBrokerService(repository ports.BrokerPort) *BrokerService {
	return &BrokerService{repository}
}

func (service BrokerService) CreateBroker(brokerDTO dto.BrokerDTO) (dto.BrokerDTO, error) {
	err := validateBrokerDTOForCreate(&brokerDTO)
	if err != nil {
		return brokerDTO, fmt.Errorf("validation error: %w", err)
	}
	newBroker := mapper.ToBroker(&brokerDTO)
	err = service.brokerRepository.Create(&newBroker)
	if err != nil {
		return brokerDTO, err
	}
	return mapper.FromBroker(&newBroker), nil
}

func (service BrokerService) ReadBroker(id uuid.UUID) (dto.BrokerDTO, error) {
	broker, err := service.brokerRepository.Find(id)
	if err != nil {
		return dto.BrokerDTO{}, err
	}
	return mapper.FromBroker(broker), nil
}

func (service BrokerService) ReadBrokers() ([]dto.BrokerDTO, error) {
	brokers, err := service.brokerRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return mapper.FromBrokers(brokers), nil
}

func (service BrokerService) UpdateBroker(brokerDTO dto.BrokerDTO) error {
	err := validateBrokerDTOForUpdateOrDelete(&brokerDTO)
	if err != nil {
		return err
	}
	_, err = service.brokerRepository.Find(brokerDTO.ID)
	if err != nil {
		return fmt.Errorf("broker with ID %s not found", brokerDTO.ID)
	}
	updatedBroker := mapper.ToBroker(&brokerDTO)
	return service.brokerRepository.Update(&updatedBroker)
}

func (service BrokerService) DeleteBroker(brokerDTO dto.BrokerDTO) error {
	err := validateBrokerDTOForUpdateOrDelete(&brokerDTO)
	if err != nil {
		return err
	}
	return service.brokerRepository.Delete(brokerDTO.ID)
}
