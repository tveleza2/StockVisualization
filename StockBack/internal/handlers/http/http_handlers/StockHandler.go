package handlers

import (
	"net/http"

	"stock-app/internal/core/services"

	"github.com/gin-gonic/gin"
)

type StockHandler struct {
	stockService services.StockService
}

func NewStockHandler(stockService services.StockService) *StockHandler {
	return &StockHandler{
		stockService: stockService,
	}
}

// ListStocks handles GET /stocks
func (h *StockHandler) ListStocks(context *gin.Context) {
	stocks, err := h.stockService.ReadStocks()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, stocks)
}
