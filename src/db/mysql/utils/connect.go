package mysql_utils

import (
	"atcscraper/src/env"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

func CreateDatabaseConnection() *sqlx.DB {

	// Get DB Credentials
	DBName := env.LoadEnv("DB_NAME")
	DBEndpoint := env.LoadEnv("DB_ENDPOINT")
	DBUsername := env.LoadEnv("DB_USERNAME")
	DBPassword := env.LoadEnv("DB_PASSWORD")

	// Build Connection String
	DBConnectionString := DBUsername + ":" + DBPassword + "@tcp(" + DBEndpoint + ")/" + DBName

	// Connect To DB + Catch Any Errors
	DBConnection, DBError := sqlx.Connect("mysql", DBConnectionString)
	if DBError != nil {
		log.Fatal(DBError)
	}

	// Return DB Connection
	return DBConnection

}
