package mysql_utils

import (
	"atcscraper/src/env"
	"atcscraper/src/types/aws"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

func LoadAWSDBSecret(SecretName string) aws.AWSDBSecret {

	SecretJSON := os.Getenv(SecretName)

	if SecretJSON == "" {
		log.Fatal("No DB Secret Found!")
	}

	DBSecret := aws.AWSDBSecret{}
	_ = json.Unmarshal([]byte(SecretJSON), &DBSecret)

	if SecretJSON == "" {
		log.Fatalf("Error Loading DB Secret Env Var: '%s'", SecretName)
	}

	// Return DBSecret
	return DBSecret
}

func CreateDatabaseConnection() *sqlx.DB {

	// Parse DB Secret
	DBSecret := LoadAWSDBSecret("ATC_DB_Credentials")

	// Get DB Credentials
	DBName := env.LoadEnv("DB_NAME")
	DBEndpoint := env.LoadEnv("DB_ENDPOINT")
	DBUsername := DBSecret.Username
	DBPassword := DBSecret.Password

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
