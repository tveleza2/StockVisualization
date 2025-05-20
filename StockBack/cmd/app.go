package main

import (
	"fmt"
	"net/http"
	"stock-app/internal/core/services"
	web "stock-app/internal/handlers/http"
	"stock-app/internal/repositories/db"

	"github.com/gin-gonic/gin"
)

func main() {
	database := db.InitDB()
	actionRepo := db.NewActionRepository(database)
	actionService := services.NewActionService(actionRepo)
	ratHisRepo := db.NewRatingHistoricRepository(database)
	brokerRepo := db.NewBrokerRepository(database)
	brokerService := services.NewBrokerService(brokerRepo)
	ratingRepo := db.NewRatingRepository(database)
	ratingService := services.NewRatingService(ratingRepo)
	brokerStockRepo := db.NewBrokerStockRepository(database)
	brokerStockService := services.NewBrokerStockService(brokerStockRepo, *brokerService)
	ratHisService := services.NewRatingHistoricService(ratHisRepo, *brokerStockService, *actionService, *ratingService)

	fetchedData, err := ratHisService.FetchRatingsFromSource()

	if err != nil {
		fmt.Println("Error fetching ratings:", err)
		return
	}
	if fetchedData == nil {
		fmt.Println("No data received.")
		return
	}
	for _, dto := range *fetchedData {
		fmt.Println(dto)
	}
	router := web.NewRouter(*actionService, *ratHisService)

	router.Run(":8080")

}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello"})
}
