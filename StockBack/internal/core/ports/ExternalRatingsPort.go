package ports

import "stock-app/internal/handlers/dto"

type ExternalRatingsPort interface {
	FetchRatingsFromSource(endpoint string, authToken string) ([]dto.FullResponseRatingHistoricDTO, error)
}
