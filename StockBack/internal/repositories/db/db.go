package db

import (
	"os"
	"stock-app/internal/core/domain"

	_ "github.com/jackc/pgx/v5/stdlib"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	dsn := os.Getenv("DATABASE_URL")
	var err error

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	dbMigration(DB)

	// createTables()
}

func dbMigration(dataBase *gorm.DB) {
	dataBase.AutoMigrate(&domain.Action{})
	dataBase.AutoMigrate(&domain.Broker{})
	dataBase.AutoMigrate(&domain.Rating{})
	dataBase.AutoMigrate(&domain.Stock{})
	dataBase.AutoMigrate(&domain.StockPrices{})
	dataBase.AutoMigrate(&domain.BrokerStock{})
	dataBase.AutoMigrate(&domain.BrokerRatingsHistoric{})
}

// func createTables() {
// 	createBrokerTable := `
// 	CREATE TABLE IF NOT EXISTs
// 		brokerage(
// 			broker_id uuid NOT NULL DEFAULT gen_random_uuid(),
// 			name VARCHAR(255),
// 			PRIMARY KEY(broker_id)
// 		)`
// 	_, error := DB.Exec(createBrokerTable)
// 	if error != nil {
// 		panic(error)
// 	}
// }
