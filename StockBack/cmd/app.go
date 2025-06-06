package main

import (
	"fmt"
	"net/http"
	web "stock-app/internal/handlers/http"
	"stock-app/internal/repositories/db"

	"github.com/gin-gonic/gin"
)

func main() {
	schema, err := db.NewSchema(false, false)
	if err != nil {
		fmt.Println(fmt.Errorf("error with database connection: %w", err))
	}
	fmt.Println("Succesful DB connection")
	router := web.NewRouter(&schema)
	router.Run(":8080")

}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello"})
}
