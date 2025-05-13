package handlers

import (
	"net/http"
	"stock-app/internal/core/services"
	"stock-app/internal/handlers/dto"

	"github.com/gin-gonic/gin"
)

type ActionHandler struct {
	service services.ActionService
}

func NewActionHandler(service services.ActionService) *ActionHandler {
	return &ActionHandler{service}
}

func (handler ActionHandler) CreateAction(context *gin.Context) {
	var req dto.ActionDTO
	if err := context.ShouldBindJSON(&req); err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	actionResponse, err := handler.service.CreateAction(req)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, actionResponse)
}
