package db

import (
	"os"
	"stock-app/internal/core/domain"

	_ "github.com/jackc/pgx/v5/stdlib"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {

	dsn := os.Getenv("DATABASE_URL")
	var err error

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// dbMigration(DB)
	return DB
	// createTables()
}

func dbMigration(dataBase *gorm.DB) {
	dataBase.AutoMigrate(&domain.Action{})
	dataBase.AutoMigrate(&domain.Broker{})
	dataBase.AutoMigrate(&domain.Rating{})
	dataBase.AutoMigrate(&domain.Stock{})
	dataBase.AutoMigrate(&domain.StockPrice{})
	dataBase.AutoMigrate(&domain.BrokerStock{})
	dataBase.AutoMigrate(&domain.RatingHistoric{})
}
