package routes

import (
	"stock-app/internal/core/services"
	handlers "stock-app/internal/handlers/http/http_handlers"

	"github.com/gin-gonic/gin"
)

func RegisterStockRoutes(router *gin.Engine, service services.StockService) {
	handler := handlers.NewStockHandler(service)
	stock := router.Group("/stocks")
	{
		stock.GET("/", handler.ListStocks)
	}
}
