package routes

import (
	"stock-app/internal/core/services"
	handlers "stock-app/internal/handlers/http/http_handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRatingHistoricRoutes(router *gin.Engine, service services.RatingHistoricService) {
	handler := handlers.NewRatingHistoricHandler(service)
	ratingHistoric := router.Group("/rating-historics")
	{
		ratingHistoric.POST("/", handler.CreateRatingHistoric)
		ratingHistoric.GET("/:id", handler.GetRatingHistoric)
		ratingHistoric.GET("/", handler.GetRatingHistorics)
		ratingHistoric.PUT("/", handler.UpdateRatingHistoric)
		ratingHistoric.DELETE("/", handler.DeleteRatingHistoric)
	}
}
