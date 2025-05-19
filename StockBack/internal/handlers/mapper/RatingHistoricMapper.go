package mapper

import (
	"stock-app/internal/core/domain"
	"stock-app/internal/handlers/dto"
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

func FullResponseFromRatingHistoric(ratingHistoric *domain.RatingHistoric) dto.FullResponseRatingHistoricDTO {
	return dto.FullResponseRatingHistoricDTO{
		ID:         ratingHistoric.ID,
		StockName:  ratingHistoric.BrokerStock.Stock.Name,
		FromTarget: ratingHistoric.FromTarget,
		ToTarget:   ratingHistoric.ToTarget,
		Time:       ratingHistoric.Time,
		FromRating: ratingHistoric.FromRating.Name, // Optional if loaded
		ToRating:   ratingHistoric.ToRating.Name,   // Optional if loaded
		ActionName: ratingHistoric.Action.Name,     // Optional if loaded
	}
}

func RatingHistoricFromFullResponse(FullDto *dto.FullResponseRatingHistoricDTO) domain.RatingHistoric {
	return domain.RatingHistoric{
			ID:FullDto.ID,
			BrokerStockID: FullDto.BrokerName,
			ActionID:      FullDto.ActionName,
			FromRatingID:  FullDto.FromRating,
			ToRatingID:    FullDto.ToRating,
			FromTarget:    FullDto.FromTarget,
			ToTarget:      FullDto.ToTarget,
			Time:          FullDto.Time,

			FromRating:  ,
			ToRating:    ,
			Action:     ,
			BrokerStock: ,






		ID:         ratingHistoric.ID,
		StockName:  ratingHistoric.BrokerStock.Stock.Name,
		FromTarget: ratingHistoric.FromTarget,
		ToTarget:   ratingHistoric.ToTarget,
		Time:       ratingHistoric.Time,
		FromRating: ratingHistoric.FromRating.Name, // Optional if loaded
		ToRating:   ratingHistoric.ToRating.Name,   // Optional if loaded
		ActionName: ratingHistoric.Action.Name,     // Optional if loaded
	}
}

func FromRatingHistorics(ratingHistorics []domain.RatingHistoric) []dto.RatingHistoricDTO {
	ratingHistoricsDTO := make([]dto.RatingHistoricDTO, len(ratingHistorics))
	for i, rh := range ratingHistorics {
		ratingHistoricsDTO[i] = FromRatingHistoric(&rh)
	}
	return ratingHistoricsDTO
}
