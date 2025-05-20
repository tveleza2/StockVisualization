package db

import (
	"stock-app/internal/core/domain"
	"stock-app/internal/core/ports"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RatingHistoricRepository struct {
	db *gorm.DB
}

func NewRatingHistoricRepository(db *gorm.DB) ports.RatingHistoricPort {
	return &RatingHistoricRepository{db}
}

func (repository RatingHistoricRepository) Create(ratingHistoric *domain.RatingHistoric) error {
	return repository.db.Create(ratingHistoric).Error
}

func (repository RatingHistoricRepository) Find(id uuid.UUID) (*domain.RatingHistoric, error) {
	ratingHistoric := domain.RatingHistoric{}
	return &ratingHistoric, repository.db.First(&ratingHistoric, id).Error
}

func (repository RatingHistoricRepository) FindOneByBrokerStock(id uuid.UUID) (*domain.RatingHistoric, error) {
	ratingHistoric := domain.RatingHistoric{}
	return &ratingHistoric, repository.db.Where("broker_stock_id = ?", id).First(&ratingHistoric).Error
}

func (repository RatingHistoricRepository) FindAllByStock(brokerStockIds *[]uuid.UUID) ([]domain.RatingHistoric, error) {
	ratingHistoric := []domain.RatingHistoric{}
	return ratingHistoric, repository.db.Where("stock_id IN ?", brokerStockIds).First(&ratingHistoric).Error
}

func (repository RatingHistoricRepository) FindAll() ([]domain.RatingHistoric, error) {
	var ratingHistorics []domain.RatingHistoric
	err := repository.db.Find(&ratingHistorics).Error
	return ratingHistorics, err
}

func (repository RatingHistoricRepository) Update(ratingHistoric *domain.RatingHistoric) error {
	return repository.db.Save(ratingHistoric).Error
}

func (repository RatingHistoricRepository) Delete(id uuid.UUID) error {
	return repository.db.Delete(&domain.RatingHistoric{}, id).Error
}
