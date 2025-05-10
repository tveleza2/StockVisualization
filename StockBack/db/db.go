package db

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDB() {

	connStr := os.Getenv("DATABASE_URL")
	var err error
	DB, err = sql.Open("pgx", connStr)

	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}
func createTables() {
	createBrokerTable := `
	CREATE TABLE IF NOT EXISTs 
		brokerage(
			broker_id uuid NOT NULL DEFAULT gen_random_uuid(),
			name VARCHAR(255),
			PRIMARY KEY(broker_id)
		)`
	_, error := DB.Exec(createBrokerTable)
	if error != nil {
		panic(error)
	}
}
