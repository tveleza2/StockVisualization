package db

import (
	"stock-app/internal/core/domain"
	"stock-app/internal/core/ports"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BrokerRepository struct {
	db *gorm.DB
}

func NewBrokerRepository(db *gorm.DB) ports.BrokerPort {
	return &BrokerRepository{db}
}

func (repository BrokerRepository) Create(broker *domain.Broker) error {
	return repository.db.Create(broker).Error
}

func (repository BrokerRepository) Find(id uuid.UUID) (*domain.Broker, error) {
	broker := domain.Broker{}
	return &broker, repository.db.First(&broker, id).Error
}

func (repository BrokerRepository) FindByName(name string) (domain.Broker, error) {
	var broker domain.Broker
	err := repository.db.Where("name == ?", name).First(&broker).Error
	return broker, err
}
func (repository BrokerRepository) FindByNames(names *[]string) (*[]domain.Broker, error) {
	var brokers []domain.Broker
	err := repository.db.Where("name IN ", names).Find(&brokers).Error
	return &brokers, err
}

func (repository BrokerRepository) FindAll() ([]domain.Broker, error) {
	var brokers []domain.Broker
	err := repository.db.Find(&brokers).Error
	return brokers, err
}

func (repository BrokerRepository) Update(broker *domain.Broker) error {
	return repository.db.Save(broker).Error
}

func (repository BrokerRepository) Delete(id uuid.UUID) error {
	return repository.db.Delete(&domain.Broker{}, id).Error
}
