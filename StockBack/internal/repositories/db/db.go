package db

import (
	"os"
	"stock-app/internal/core/domain"
	"stock-app/internal/core/services"
	"stock-app/internal/infrastructure"

	_ "github.com/jackc/pgx/v5/stdlib"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Schema struct {
	DB            *gorm.DB
	ActionService services.ActionService
	BrokerService services.BrokerService
	BSService     services.BrokerStockService
	RHService     services.RatingHistoricService
	RatingService services.RatingService
	StockService  services.StockService
}

func NewSchema(migrateSchema bool, importData bool) (Schema, error) {
	var err error
	DB := InitDB(migrateSchema)
	actionRepo := NewActionRepository(DB)
	actionService := services.NewActionService(actionRepo)

	brokerRepo := NewBrokerRepository(DB)
	brokerService := services.NewBrokerService(brokerRepo)

	stockRepo := NewStockRepository(DB)
	stockService := services.NewStockService(stockRepo)

	ratingRepo := NewRatingRepository(DB)
	ratingService := services.NewRatingService(ratingRepo)

	brokerStockRepo := NewBrokerStockRepository(DB)
	brokerStockService := services.NewBrokerStockService(brokerStockRepo, *brokerService, *stockService)

	ratHisRepo := NewRatingHistoricRepository(DB)
	ratHisService := services.NewRatingHistoricService(ratHisRepo, *brokerStockService, *actionService, *ratingService)

	if importData {
		importRepository := infrastructure.NewImportRepository()
		importService := services.NewExternalResourcesService(importRepository, *ratHisService)
		err = importService.SaveIncomingRatings()
		ratHisService.UpdateStockScores()
	} else {
		err = nil
	}
	return Schema{
		DB:            DB,
		ActionService: *actionService,
		BrokerService: *brokerService,
		BSService:     *brokerStockService,
		RHService:     *ratHisService,
		RatingService: *ratingService,
		StockService:  *stockService,
	}, err
}

var DB *gorm.DB

func InitDB(migrate bool) *gorm.DB {

	dsn := os.Getenv("DATABASE_URL")
	var err error

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	if migrate {
		dbMigration(DB)
	}
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
