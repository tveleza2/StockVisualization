package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"
	"stock-app/internal/handlers/dto"
)

type ImportRepository struct {
}

func NewImportRepository() *ImportRepository {
	return &ImportRepository{}
}
func (repository ImportRepository) FetchRatingsFromSource(endpoint string, authToken string) ([]dto.FullResponseRatingHistoricDTO, error) {
	var dtos []dto.FullResponseRatingHistoricDTO
	var responseFormat dto.ApiResponseFromSource
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println("Error making request:", err)
		return dtos, err
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

	return responseFormat.Items, nil
}
