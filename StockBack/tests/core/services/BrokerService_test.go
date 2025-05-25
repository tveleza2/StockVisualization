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

type mockBrokerRepo struct {
	brokers []domain.Broker
}

func (m *mockBrokerRepo) Create(broker *domain.Broker) error {
	m.brokers = append(m.brokers, *broker)
	return nil
}
func (m *mockBrokerRepo) Find(id uuid.UUID) (*domain.Broker, error) {
	for _, b := range m.brokers {
		if b.ID == id {
			return &b, nil
		}
	}
	return nil, errors.New("not found")
}
func (m *mockBrokerRepo) FindAll() ([]domain.Broker, error) { return m.brokers, nil }
func (m *mockBrokerRepo) FindByName(name string) (domain.Broker, error) {
	for _, b := range m.brokers {
		if b.Name == name {
			return b, nil
		}
	}
	return domain.Broker{}, errors.New("not found")
}
func (m *mockBrokerRepo) FindByNames(names *[]string) (*[]domain.Broker, error) {
	var result []domain.Broker
	for _, n := range *names {
		for _, b := range m.brokers {
			if b.Name == n {
				result = append(result, b)
			}
		}
	}
	return &result, nil
}
func (m *mockBrokerRepo) Update(broker *domain.Broker) error {
	for i, b := range m.brokers {
		if b.ID == broker.ID {
			m.brokers[i] = *broker
			return nil
		}
	}
	return errors.New("not found")
}
func (m *mockBrokerRepo) Delete(id uuid.UUID) error {
	for i, b := range m.brokers {
		if b.ID == id {
			m.brokers = append(m.brokers[:i], m.brokers[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}

func TestBrokerService_CreateBroker(t *testing.T) {
	repo := &mockBrokerRepo{}
	service := services.NewBrokerService(repo)
	dto := dto.BrokerDTO{Name: "Goldman Sachs"}
	result, err := service.CreateBroker(dto)
	assert.NoError(t, err)
	assert.Equal(t, "Goldman Sachs", result.Name)
}

func TestBrokerService_CreateBroker_ValidationError(t *testing.T) {
	repo := &mockBrokerRepo{}
	service := services.NewBrokerService(repo)
	dto := dto.BrokerDTO{Name: ""}
	_, err := service.CreateBroker(dto)
	assert.Error(t, err)
}

func TestBrokerService_ReadBroker_Success(t *testing.T) {
	repo := &mockBrokerRepo{}
	service := services.NewBrokerService(repo)
	dto := dto.BrokerDTO{Name: "Morgan Stanley"}
	created, _ := service.CreateBroker(dto)
	found, err := service.ReadBroker(created.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Morgan Stanley", found.Name)
}

func TestBrokerService_ReadBroker_NotFound(t *testing.T) {
	repo := &mockBrokerRepo{}
	service := services.NewBrokerService(repo)
	_, err := service.ReadBroker(uuid.New())
	assert.Error(t, err)
}

func TestBrokerService_ReadBrokers(t *testing.T) {
	repo := &mockBrokerRepo{}
	service := services.NewBrokerService(repo)
	service.CreateBroker(dto.BrokerDTO{Name: "A"})
	service.CreateBroker(dto.BrokerDTO{Name: "B"})
	brokers, err := service.ReadBrokers()
	assert.NoError(t, err)
	assert.Len(t, brokers, 2)
}

func TestBrokerService_UpdateBroker_Success(t *testing.T) {
	repo := &mockBrokerRepo{}
	service := services.NewBrokerService(repo)
	created, _ := service.CreateBroker(dto.BrokerDTO{Name: "A"})
	updateDTO := dto.BrokerDTO{ID: created.ID, Name: "B"}
	err := service.UpdateBroker(updateDTO)
	assert.NoError(t, err)
	updated, _ := service.ReadBroker(created.ID)
	assert.Equal(t, "B", updated.Name)
}

func TestBrokerService_UpdateBroker_ValidationError(t *testing.T) {
	repo := &mockBrokerRepo{}
	service := services.NewBrokerService(repo)
	updateDTO := dto.BrokerDTO{ID: uuid.New(), Name: ""}
	err := service.UpdateBroker(updateDTO)
	assert.Error(t, err)
}

func TestBrokerService_UpdateBroker_NotFound(t *testing.T) {
	repo := &mockBrokerRepo{}
	service := services.NewBrokerService(repo)
	updateDTO := dto.BrokerDTO{ID: uuid.New(), Name: "Update"}
	err := service.UpdateBroker(updateDTO)
	assert.Error(t, err)
}

func TestBrokerService_DeleteBroker_Success(t *testing.T) {
	repo := &mockBrokerRepo{}
	service := services.NewBrokerService(repo)
	created, _ := service.CreateBroker(dto.BrokerDTO{Name: "A"})
	deleteDTO := dto.BrokerDTO{ID: created.ID, Name: created.Name}
	err := service.DeleteBroker(deleteDTO)
	assert.NoError(t, err)
	_, err = service.ReadBroker(created.ID)
	assert.Error(t, err)
}

func TestBrokerService_DeleteBroker_NotFound(t *testing.T) {
	repo := &mockBrokerRepo{}
	service := services.NewBrokerService(repo)
	deleteDTO := dto.BrokerDTO{ID: uuid.New(), Name: "Nonexistent"}
	err := service.DeleteBroker(deleteDTO)
	assert.Error(t, err)
}

func TestBrokerService_FindByName_Existing(t *testing.T) {
	repo := &mockBrokerRepo{}
	service := services.NewBrokerService(repo)
	service.CreateBroker(dto.BrokerDTO{Name: "Broker1"})
	broker, err := service.FindByName("Broker1")
	assert.NoError(t, err)
	assert.Equal(t, "Broker1", broker.Name)
}

func TestBrokerService_FindByName_NotExisting(t *testing.T) {
	repo := &mockBrokerRepo{}
	service := services.NewBrokerService(repo)
	broker, err := service.FindByName("CreateMe")
	assert.NoError(t, err)
	assert.Equal(t, "CreateMe", broker.Name)
}

func TestBrokerService_FindByMapOfNames(t *testing.T) {
	repo := &mockBrokerRepo{}
	service := services.NewBrokerService(repo)
	service.CreateBroker(dto.BrokerDTO{Name: "A"})
	service.CreateBroker(dto.BrokerDTO{Name: "B"})
	names := []string{"A", "B"}
	result, err := service.FindByMapOfNames(&names)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(*result))
}
