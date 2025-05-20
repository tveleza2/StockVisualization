package main

import (
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

	brokerRepo := db.NewBrokerRepository(database)
	brokerService := services.NewBrokerService(brokerRepo)

	stockRepo := db.NewStockRepository(database)
	stockService := services.NewStockService(stockRepo)

	ratingRepo := db.NewRatingRepository(database)
	ratingService := services.NewRatingService(ratingRepo)

	brokerStockRepo := db.NewBrokerStockRepository(database)
	brokerStockService := services.NewBrokerStockService(brokerStockRepo, *brokerService, *stockService)

	ratHisRepo := db.NewRatingHistoricRepository(database)
	ratHisService := services.NewRatingHistoricService(ratHisRepo, *brokerStockService, *actionService, *ratingService)

	// importRepository := infrastructure.NewImportRepository()
	// importService := services.NewExternalResourcesService(importRepository, *ratHisService)
	router := web.NewRouter(*actionService, *ratHisService)

	// err := importService.SaveIncomingRatings()

	// if err != nil {
	// 	fmt.Println("Error persisting the incoming data: %w", err)
	// }

	router.Run(":8080")

}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello"})
}
