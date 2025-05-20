package handlers

import (
	"net/http"
	"stock-app/internal/core/services"
	"stock-app/internal/handlers/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RatingHistoricHandler struct {
	service services.RatingHistoricService
}

func NewRatingHistoricHandler(service services.RatingHistoricService) *RatingHistoricHandler {
	return &RatingHistoricHandler{service}
}

func (handler RatingHistoricHandler) CreateRatingHistoric(context *gin.Context) {
	var req dto.RatingHistoricDTO
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	resp, err := handler.service.CreateRatingHistoric(req)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, resp)
}

func (handler RatingHistoricHandler) GetRatingHistoric(context *gin.Context) {
	idParam := context.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	resp, err := handler.service.ReadRatingHistoric(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, resp)
}

func (handler RatingHistoricHandler) GetRatingHistorics(context *gin.Context) {
	// resp, err := handler.service.FetchRatingsFromSource()
	resp, err := handler.service.GetRatingsFromDB()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, resp)
}

func (handler RatingHistoricHandler) UpdateRatingHistoric(context *gin.Context) {
	var req dto.RatingHistoricDTO
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := handler.service.UpdateRatingHistoric(req)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.Status(http.StatusNoContent)
}

func (handler RatingHistoricHandler) DeleteRatingHistoric(context *gin.Context) {
	var req dto.RatingHistoricDTO
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := handler.service.DeleteRatingHistoric(req)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.Status(http.StatusNoContent)
}
