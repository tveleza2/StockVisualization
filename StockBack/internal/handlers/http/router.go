package web

import (
	"stock-app/internal/core/services"
	"stock-app/internal/handlers/http/routes"

	"github.com/gin-gonic/gin"
)

func NewRouter(actionService services.ActionService) *gin.Engine {
	router := gin.Default()
	routes.RegisterActionRoutes(router, actionService)
	return router
}
