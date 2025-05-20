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
	brokerStock, err := service.bSService.FindByBrokerAndStock(dto.BrokerName, dto.StockID)
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
	ratingHistoric, err := mapper.RatingHistoricFromFullResponse(&dto, brokerStock, action, fromRating, toRating)
	if err != nil {
		return fmt.Errorf("the dto mapping failed")
	}
	return service.ratingHistoricRepository.Update(&ratingHistoric)
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

// func (service RatingHistoricService) SaveMultipleResponseRatingHistoric(dtos *[]dto.FullResponseRatingHistoricDTO) error {
// 	numberOfNewEntries := len(*dtos)
// 	actionNames, brokerNames, ratingNames := make([]string, numberOfNewEntries), make([]string, numberOfNewEntries), make([]string, 2*numberOfNewEntries)
// 	var ratingHistoric domain.RatingHistoric

// 	for i, dto := range *dtos {
// 		actionNames[i] = dto.ActionName
// 		brokerNames[i] = dto.BrokerName
// 		ratingNames[2*i] = dto.FromRating
// 		ratingNames[2*i+1] = dto.ToRating
// 	}
// 	actions, err := service.actService.FindByNames(&actionNames)
// 	if err != nil {
// 		return err
// 	}
// 	ratings, err := service.ratService.FindByNames(&ratingNames)
// 	if err != nil {
// 		return err
// 	}
// 	brokerStocks, err := service.bSService.FindByBrokersAndStock(brokerNames,)
// 	if err != nil {
// 		return err
// 	}

// 	for _, dto := range *dtos {
// 		ratingHistoric = mapper.RatingHistoricFromFullResponse(dto,brokerStocks[],actions[dto.ActionName],ratings[dto.FromRating],ratings[dto.ToRating])
// 	}

// 	return nil
// }
