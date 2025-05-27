package services_test

import (
	"errors"
	"stock-app/internal/core/domain"
	"stock-app/internal/core/services"
	"stock-app/internal/handlers/dto"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type mockRatingRepo struct {
	ratings []domain.Rating
}

func (m *mockRatingRepo) Create(r *domain.Rating) error {
	m.ratings = append(m.ratings, *r)
	return nil
}
func (m *mockRatingRepo) Find(id uuid.UUID) (*domain.Rating, error) {
	for _, r := range m.ratings {
		if r.ID == id {
			return &r, nil
		}
	}
	return nil, errors.New("not found")
}
func (m *mockRatingRepo) FindAll() ([]domain.Rating, error) { return m.ratings, nil }
func (m *mockRatingRepo) FindByName(name string) (domain.Rating, error) {
	for _, r := range m.ratings {
		if r.Name == name {
			return r, nil
		}
	}
	return domain.Rating{}, errors.New("not found")
}
func (m *mockRatingRepo) FindByNames(names *[]string) (*[]domain.Rating, error) {
	var result []domain.Rating
	for _, n := range *names {
		for _, r := range m.ratings {
			if r.Name == n {
				result = append(result, r)
			}
		}
	}
	return &result, nil
}
func (m *mockRatingRepo) Update(r *domain.Rating) error { return nil }
func (m *mockRatingRepo) Delete(id uuid.UUID) error     { return nil }

func TestRatingService_CreateRating(t *testing.T) {
	repo := &mockRatingRepo{}
	service := services.NewRatingService(repo)
	dto := dto.RatingDTO{Name: "Buy"}
	result, err := service.CreateRating(dto)
	assert.NoError(t, err)
	assert.Equal(t, "Buy", result.Name)
}

func TestRatingService_CreateRating_ValidationError(t *testing.T) {
	repo := &mockRatingRepo{}
	service := services.NewRatingService(repo)
	dto := dto.RatingDTO{Name: ""}
	_, err := service.CreateRating(dto)
	assert.Error(t, err)
}

func TestRatingService_ReadRating_Success(t *testing.T) {
	repo := &mockRatingRepo{}
	service := services.NewRatingService(repo)
	dto := dto.RatingDTO{Name: "Buy"}
	created, _ := service.CreateRating(dto)
	found, err := service.ReadRating(created.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Buy", found.Name)
}

func TestRatingService_ReadRating_NotFound(t *testing.T) {
	repo := &mockRatingRepo{}
	service := services.NewRatingService(repo)
	_, err := service.ReadRating(uuid.New())
	assert.Error(t, err)
}

func TestRatingService_ReadRatings(t *testing.T) {
	repo := &mockRatingRepo{}
	service := services.NewRatingService(repo)
	service.CreateRating(dto.RatingDTO{Name: "Buy"})
	service.CreateRating(dto.RatingDTO{Name: "Sell"})
	ratings, err := service.ReadRatings()
	assert.NoError(t, err)
	assert.Len(t, ratings, 2)
}

func TestRatingService_UpdateRating_Success(t *testing.T) {
	repo := &mockRatingRepo{}
	service := services.NewRatingService(repo)
	created, _ := service.CreateRating(dto.RatingDTO{ID: uuid.New(), Name: "Buy"})
	updateDTO := dto.RatingDTO{ID: created.ID, Name: "Strong Buy"}
	err := service.UpdateRating(updateDTO)
	assert.NoError(t, err)
}

func TestRatingService_UpdateRating_ValidationError(t *testing.T) {
	repo := &mockRatingRepo{}
	service := services.NewRatingService(repo)
	updateDTO := dto.RatingDTO{ID: uuid.New(), Name: ""}
	err := service.UpdateRating(updateDTO)
	assert.Error(t, err)
}

func TestRatingService_UpdateRating_NotFound(t *testing.T) {
	repo := &mockRatingRepo{}
	service := services.NewRatingService(repo)
	updateDTO := dto.RatingDTO{ID: uuid.New(), Name: "Update"}
	err := service.UpdateRating(updateDTO)
	assert.Error(t, err)
}

func TestRatingService_DeleteRating_Success(t *testing.T) {
	repo := &mockRatingRepo{}
	service := services.NewRatingService(repo)
	created, _ := service.CreateRating(dto.RatingDTO{ID: uuid.New(), Name: "Buy"})
	deleteDTO := dto.RatingDTO{ID: created.ID, Name: created.Name}
	err := service.DeleteRating(deleteDTO)
	assert.NoError(t, err)
}

func TestRatingService_DeleteRating_NotFound(t *testing.T) {
	repo := &mockRatingRepo{}
	service := services.NewRatingService(repo)
	deleteDTO := dto.RatingDTO{Name: "Nonexistent"}
	err := service.DeleteRating(deleteDTO)
	assert.Error(t, err)
}

func TestRatingService_FindByName_Existing(t *testing.T) {
	repo := &mockRatingRepo{}
	service := services.NewRatingService(repo)
	service.CreateRating(dto.RatingDTO{Name: "Buy"})
	rating, err := service.FindByName("Buy")
	assert.NoError(t, err)
	assert.Equal(t, "Buy", rating.Name)
}

func TestRatingService_FindByName_NotExisting(t *testing.T) {
	repo := &mockRatingRepo{}
	service := services.NewRatingService(repo)
	_, err := service.FindByName("IDontExist")
	assert.Error(t, err)
}

func TestRatingService_FindByNames(t *testing.T) {
	repo := &mockRatingRepo{}
	service := services.NewRatingService(repo)
	service.CreateRating(dto.RatingDTO{Name: "Buy"})
	service.CreateRating(dto.RatingDTO{Name: "Sell"})
	names := []string{"Buy", "Sell"}
	result, err := service.FindByNames(&names)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(*result))
}
