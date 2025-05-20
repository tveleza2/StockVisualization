package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"stock-app/internal/core/ports"
	"stock-app/internal/handlers/dto"
	"stock-app/internal/handlers/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func validateRatingHistoricDTOForCreate(dto *dto.RatingHistoricDTO) error {
	if dto.BrokerStockID == uuid.Nil {
		return errors.New("rating historic has no broker stock ID")
	}
	if dto.ActionID == uuid.Nil {
		return errors.New("rating historic has no action ID")
	}
	return nil
}

func validateRatingHistoricDTOForUpdateOrDelete(dto *dto.RatingHistoricDTO) error {
	if dto.ID == uuid.Nil {
		return errors.New("rating historic has no ID")
	}
	return nil
}

type RatingHistoricService struct {
	ratingHistoricRepository ports.RatingHistoricPort
	bSService                BrokerStockService
	actService               ActionService
	ratService               RatingService
}

func NewRatingHistoricService(repository ports.RatingHistoricPort, brokerStockService BrokerStockService, actionService ActionService, ratingService RatingService) *RatingHistoricService {
	return &RatingHistoricService{
		ratingHistoricRepository: repository,
		bSService:                brokerStockService,
		actService:               actionService,
		ratService:               ratingService,
	}
}

func (service RatingHistoricService) CreateRatingHistoric(ratingHistoricDTO dto.RatingHistoricDTO) (dto.RatingHistoricDTO, error) {
	err := validateRatingHistoricDTOForCreate(&ratingHistoricDTO)
	if err != nil {
		return ratingHistoricDTO, fmt.Errorf("validation error: %w", err)
	}
	newRatingHistoric := mapper.ToRatingHistoric(&ratingHistoricDTO)
	err = service.ratingHistoricRepository.Create(&newRatingHistoric)
	if err != nil {
		return ratingHistoricDTO, err
	}
	return mapper.FromRatingHistoric(&newRatingHistoric), nil
}

func (service RatingHistoricService) ReadRatingHistoric(id uuid.UUID) (dto.RatingHistoricDTO, error) {
	ratingHistoric, err := service.ratingHistoricRepository.Find(id)
	if err != nil {
		return dto.RatingHistoricDTO{}, err
	}
	return mapper.FromRatingHistoric(ratingHistoric), nil
}

func (service RatingHistoricService) ReadRatingHistorics() ([]dto.RatingHistoricDTO, error) {
	ratingHistorics, err := service.ratingHistoricRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return mapper.FromRatingHistorics(ratingHistorics), nil
}

func (service RatingHistoricService) UpdateRatingHistoric(ratingHistoricDTO dto.RatingHistoricDTO) error {
	err := validateRatingHistoricDTOForUpdateOrDelete(&ratingHistoricDTO)
	if err != nil {
		return err
	}
	_, err = service.ratingHistoricRepository.Find(ratingHistoricDTO.ID)
	if err != nil {
		return fmt.Errorf("rating historic with ID %s not found", ratingHistoricDTO.ID)
	}
	updatedRatingHistoric := mapper.ToRatingHistoric(&ratingHistoricDTO)
	return service.ratingHistoricRepository.Update(&updatedRatingHistoric)
}

func (service RatingHistoricService) DeleteRatingHistoric(ratingHistoricDTO dto.RatingHistoricDTO) error {
	err := validateRatingHistoricDTOForUpdateOrDelete(&ratingHistoricDTO)
	if err != nil {
		return err
	}
	return service.ratingHistoricRepository.Delete(ratingHistoricDTO.ID)
}

func (service RatingHistoricService) ReadRatingHistoricByStock(stock string) ([]dto.FullResponseRatingHistoricDTO, error) {
	brokerStockIds, err := service.bSService.IdsByStock(stock)
	if err != nil {
		return []dto.FullResponseRatingHistoricDTO{}, err
	}
	ratingHistoric, err := service.ratingHistoricRepository.FindAllByStock(brokerStockIds)
	if err != nil {
		return []dto.FullResponseRatingHistoricDTO{}, err
	}
	return mapper.FullResponseFromRatingHistorics(ratingHistoric), nil
}

func (service RatingHistoricService) SaveResponseRatingHistoric(dto dto.FullResponseRatingHistoricDTO) error {

	action, err := service.actService.FindByName(dto.ActionName)
	if err != nil {
		return err
	}
	brokerStock, err := service.bSService.FindByBrokerAndStock(dto.BrokerName, dto.StockID, dto.StockName)
	if err != nil {
		return err
	}
	fromRating, err := service.ratService.FindByName(dto.FromRating)
	if err != nil {
		return err
	}
	toRating, err := service.ratService.FindByName(dto.ToRating)
	if err != nil {
		return err
	}
	newRatingHistoric, err := mapper.RatingHistoricFromFullResponse(&dto, brokerStock, action, fromRating, toRating)
	if err != nil {
		return fmt.Errorf("the dto mapping failed")
	}
	oldRatingHistoric, err := service.ratingHistoricRepository.FindExistence(brokerStock.ID, dto.Time)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return service.ratingHistoricRepository.Update(&newRatingHistoric)
		}
		return fmt.Errorf("there was an error in the connection with database")
	}
	newRatingHistoric.ID = oldRatingHistoric.ID
	return service.ratingHistoricRepository.Update(&newRatingHistoric)
}

func (service RatingHistoricService) SaveMultipleResponsesRatingHistoric(dtos []dto.FullResponseRatingHistoricDTO) error {
	for _, dto := range dtos {
		err := service.SaveResponseRatingHistoric(dto)
		if err != nil {
			return err
		}
	}
	return nil
}

func (service RatingHistoricService) FetchRatingsFromSource() (*[]dto.FullResponseRatingHistoricDTO, error) {
	var dtos []dto.FullResponseRatingHistoricDTO
	var responseFormat dto.ApiResponseFromSource
	endpoint := os.Getenv("DATA_SOURCE")
	authToken := os.Getenv("AUTH_TOKEN")
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println("Error making request:", err)
		return &dtos, err
	}
	request.Header.Set("Authorization", "Bearer "+authToken)
	request.Header.Set("Accept", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status: %w", err)
	}

	if err := json.NewDecoder(response.Body).Decode(&responseFormat); err != nil {
		return nil, fmt.Errorf("JSON decoding failed: %w", err)
	}

	return &responseFormat.Items, nil
}

func (service RatingHistoricService) GetRatingsFromDB() (*[]dto.FullResponseRatingHistoricDTO, error) {
	ratingHistorics, err := service.ratingHistoricRepository.FindAll()
	if err != nil {
		return &[]dto.FullResponseRatingHistoricDTO{}, err
	}
	fmt.Println("ENTITIES: ", ratingHistorics)
	dtos := mapper.FullResponseFromRatingHistorics(ratingHistorics)
	fmt.Println("DTOS:", dtos)
	return &dtos, nil
}
