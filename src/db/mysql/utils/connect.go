package mysql_utils

import (
	"atcscraper/src/env"
	logging "atcscraper/src/log"
	"atcscraper/src/types/aws"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

func LoadAWSDBSecret(SecretName string) aws.AWSDBSecret {

	SecretJSON := os.Getenv(SecretName)

	if SecretJSON == "" {
		Error := fmt.Sprintf("No DB Secret Found With Name: %v", SecretName)
		logging.NewError(Error)
	}

	DBSecret := aws.AWSDBSecret{}
	_ = json.Unmarshal([]byte(SecretJSON), &DBSecret)

	if SecretJSON == "" {
		Error := fmt.Sprintf("Error Loading DB Secret Env Var: %v", SecretName)
		logging.NewError(Error)
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
		Error := fmt.Sprintf("Error Connecting To DB: %v", DBError.Error())
		logging.NewError(Error)
	}

	// Return DB Connection
	return DBConnection

}
