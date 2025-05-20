package db

import (
	"stock-app/internal/core/domain"
	"stock-app/internal/core/ports"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RatingRepository struct {
	db *gorm.DB
}

func NewRatingRepository(db *gorm.DB) ports.RatingPort {
	return &RatingRepository{db}
}

func (repository RatingRepository) Create(rating *domain.Rating) error {
	return repository.db.Create(rating).Error
}

func (repository RatingRepository) Find(id uuid.UUID) (*domain.Rating, error) {
	rating := domain.Rating{}
	return &rating, repository.db.First(&rating, id).Error
}

func (repository RatingRepository) FindByNames(names *[]string) (*[]domain.Rating, error) {
	var ratings []domain.Rating
	err := repository.db.Where("name IN ", names).Find(&ratings).Error
	return &ratings, err
}

func (repository RatingRepository) FindByName(name string) (domain.Rating, error) {
	var rating domain.Rating
	err := repository.db.Where("name == ", name).First(&rating).Error
	return rating, err
}

func (repository RatingRepository) FindAll() ([]domain.Rating, error) {
	var ratings []domain.Rating
	err := repository.db.Find(&ratings).Error
	return ratings, err
}

func (repository RatingRepository) Update(rating *domain.Rating) error {
	return repository.db.Save(rating).Error
}

func (repository RatingRepository) Delete(id uuid.UUID) error {
	return repository.db.Delete(&domain.Rating{}, id).Error
}
