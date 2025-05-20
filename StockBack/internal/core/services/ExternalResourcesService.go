package services

import (
	"fmt"
	"os"
	"stock-app/internal/core/ports"
)

type ExternalResourcesService struct {
	ratingsAPI            ports.ExternalRatingsPort
	ratingHistoricService RatingHistoricService
}

func NewExternalResourcesService(ratingsApi ports.ExternalRatingsPort, ratingHistoricService RatingHistoricService) *ExternalResourcesService {
	return &ExternalResourcesService{
		ratingsAPI:            ratingsApi,
		ratingHistoricService: ratingHistoricService,
	}
}

func (service ExternalResourcesService) SaveIncomingRatings() error {
	endpoint := os.Getenv("DATA_SOURCE")
	authToken := os.Getenv("AUTH_TOKEN")
	incomingRatings, err := service.ratingsAPI.FetchRatingsFromSource(endpoint, authToken)
	fmt.Println("Fetched Data", incomingRatings)
	if err != nil {
		return err
	}
	return service.ratingHistoricService.SaveMultipleResponsesRatingHistoric(incomingRatings)
}
