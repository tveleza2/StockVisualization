package web

import (
	"stock-app/internal/core/services"
	"stock-app/internal/handlers"
	"stock-app/internal/handlers/http/routes"
	"stock-app/internal/repositories/db"

	"github.com/gin-gonic/gin"
)

func NewTestRouter(actionService services.ActionService, ratingHistoricService services.RatingHistoricService) *gin.Engine {
	router := gin.Default()
	router.Use(handlers.CORSMiddleware())
	routes.RegisterActionRoutes(router, actionService)
	routes.RegisterRatingHistoricRoutes(router, ratingHistoricService)
	return router
}

func NewRouter(schema *db.Schema) *gin.Engine {
	router := gin.Default()
	router.Use(handlers.CORSMiddleware())
	routes.RegisterActionRoutes(router, schema.ActionService)
	routes.RegisterRatingHistoricRoutes(router, schema.RHService)
	return router
}
