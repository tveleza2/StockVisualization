package routes

import (
	"stock-app/internal/core/services"
	handlers "stock-app/internal/handlers/http/http_handlers"

	"github.com/gin-gonic/gin"
)

func RegisterActionRoutes(router *gin.Engine, service services.ActionService) {
	handler := handlers.NewActionHandler(service)
	actions := router.Group("/actions")
	{
		actions.POST("/", handler.CreateAction)
		// actions.GET("/", handler.GetActions)
		// actions.GET("/:id", handler.GetAction)
		// actions.PUT("/:id", handler.UpdateAction)
		// actions.DELETE("/:id", handler.DeleteAction)
	}
}
