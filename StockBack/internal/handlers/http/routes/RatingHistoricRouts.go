package routes

import (
	"stock-app/internal/core/services"
	handlers "stock-app/internal/handlers/http/http_handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRatingHistoricRoutes(router *gin.Engine, service services.RatingHistoricService) {
	handler := handlers.NewRatingHistoricHandler(service)
	ratingHistoric := router.Group("/actions")
	{
		ratingHistoric.POST("/rating-historics", handler.CreateRatingHistoric)
		ratingHistoric.GET("/rating-historics/:id", handler.GetRatingHistoric)
		ratingHistoric.GET("/rating-historics", handler.GetRatingHistorics)
		ratingHistoric.PUT("/rating-historics", handler.UpdateRatingHistoric)
		ratingHistoric.DELETE("/rating-historics", handler.DeleteRatingHistoric)
	}
}
