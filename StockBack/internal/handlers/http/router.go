package web

import (
	"stock-app/internal/core/services"
	"stock-app/internal/handlers"
	"stock-app/internal/handlers/http/routes"

	"github.com/gin-gonic/gin"
)

func NewRouter(actionService services.ActionService, ratingHistoricService services.RatingHistoricService) *gin.Engine {
	router := gin.Default()
	router.Use(handlers.CORSMiddleware())
	routes.RegisterActionRoutes(router, actionService)
	routes.RegisterRatingHistoricRoutes(router, ratingHistoricService)
	return router
}
