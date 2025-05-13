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
	router := web.NewRouter(*actionService)

	router.Run(":8080")

}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello"})
}
