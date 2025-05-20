package mapper

import (
	"stock-app/internal/core/domain"
	"stock-app/internal/handlers/dto"
	"strconv"
	"strings"
)

func ToRatingHistoric(ratingHistoricDTO *dto.RatingHistoricDTO) domain.RatingHistoric {
	return domain.RatingHistoric{
		ID:            ratingHistoricDTO.ID,
		BrokerStockID: ratingHistoricDTO.BrokerStockID,
		ActionID:      ratingHistoricDTO.ActionID,
		FromRatingID:  ratingHistoricDTO.FromRatingID,
		ToRatingID:    ratingHistoricDTO.ToRatingID,
		FromTarget:    ratingHistoricDTO.FromTarget,
		ToTarget:      ratingHistoricDTO.ToTarget,
		Time:          ratingHistoricDTO.Time,
	}
}

func FromRatingHistoric(ratingHistoric *domain.RatingHistoric) dto.RatingHistoricDTO {
	return dto.RatingHistoricDTO{
		ID:             ratingHistoric.ID,
		BrokerStockID:  ratingHistoric.BrokerStockID,
		ActionID:       ratingHistoric.ActionID,
		FromRatingID:   ratingHistoric.FromRatingID,
		ToRatingID:     ratingHistoric.ToRatingID,
		FromTarget:     ratingHistoric.FromTarget,
		ToTarget:       ratingHistoric.ToTarget,
		Time:           ratingHistoric.Time,
		FromRatingName: ratingHistoric.FromRating.Name, // Optional if loaded
		ToRatingName:   ratingHistoric.ToRating.Name,   // Optional if loaded
		ActionName:     ratingHistoric.Action.Name,     // Optional if loaded
	}
}

func FromRatingHistorics(ratingHistorics []domain.RatingHistoric) []dto.RatingHistoricDTO {
	ratingHistoricsDTO := make([]dto.RatingHistoricDTO, len(ratingHistorics))
	for i, rh := range ratingHistorics {
		ratingHistoricsDTO[i] = FromRatingHistoric(&rh)
	}
	return ratingHistoricsDTO
}

func FullResponseFromRatingHistoric(ratingHistoric *domain.RatingHistoric) dto.FullResponseRatingHistoricDTO {
	fromTarget := strconv.FormatFloat(ratingHistoric.FromTarget, 'f', -1, 64)
	toTarget := strconv.FormatFloat(ratingHistoric.ToTarget, 'f', -1, 64)
	return dto.FullResponseRatingHistoricDTO{
		ID:         ratingHistoric.ID,
		BrokerName: ratingHistoric.BrokerStock.Broker.Name,
		StockID:    ratingHistoric.BrokerStock.Stock.ID,
		StockName:  ratingHistoric.BrokerStock.Stock.Name,
		FromTarget: fromTarget,
		ToTarget:   toTarget,
		Time:       ratingHistoric.Time,
		FromRating: ratingHistoric.FromRating.Name,
		ToRating:   ratingHistoric.ToRating.Name,
		ActionName: ratingHistoric.Action.Name,
	}
}

func FullResponseFromRatingHistorics(ratingHistorics []domain.RatingHistoric) []dto.FullResponseRatingHistoricDTO {
	ratingHistoricsDTO := make([]dto.FullResponseRatingHistoricDTO, len(ratingHistorics))
	for i, rh := range ratingHistorics {
		ratingHistoricsDTO[i] = FullResponseFromRatingHistoric(&rh)
	}
	return ratingHistoricsDTO
}

func RatingHistoricFromFullResponse(fullDto *dto.FullResponseRatingHistoricDTO, brokerStock domain.BrokerStock, action domain.Action, fromRating domain.Rating, toRating domain.Rating) (domain.RatingHistoric, error) {
	fromTarget, err := parseDollarAmount(fullDto.FromTarget)
	if err != nil {
		return domain.RatingHistoric{}, err
	}
	toTarget, err := parseDollarAmount(fullDto.ToTarget)
	if err != nil {
		return domain.RatingHistoric{}, err
	}
	return domain.RatingHistoric{
		BrokerStock:   brokerStock,
		BrokerStockID: brokerStock.ID,
		ActionID:      action.ID,
		FromRatingID:  fromRating.ID,
		ToRatingID:    toRating.ID,
		FromTarget:    fromTarget,
		ToTarget:      toTarget,
		Time:          fullDto.Time,
	}, nil
}
func parseDollarAmount(s string) (float64, error) {
	s = strings.TrimPrefix(s, "$")
	return strconv.ParseFloat(s, 64)
}
