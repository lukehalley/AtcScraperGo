package mysql_utils

import (
	"atcscraper/src/env"
	logging "atcscraper/src/log"
	"atcscraper/src/types/aws"
	"atcscraper/src/util"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"strings"
)

func LoadAWSDBSecret(SecretName string) aws.AWSDBSecret {

	SecretJSON := os.Getenv(SecretName)

	if SecretJSON == "" {
		Error := fmt.Sprintf("No DB Secret Found With Name: %v", SecretName)
		logging.NewError(Error)
		logging.NewError("No DB Secret Found!")
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

func CheckConnectionError(DBConnectionError error) bool {
	TooManyConnectionsError := strings.Contains(DBConnectionError.Error(), "Error 1040")
	BadConnectionError := strings.Contains(DBConnectionError.Error(), "bad connection")
	TimeoutError := strings.Contains(DBConnectionError.Error(), "operation timed out")
	SyncError := strings.Contains(DBConnectionError.Error(), "commands out of sync")
	IsRetryError := TooManyConnectionsError || BadConnectionError || TimeoutError || SyncError
	return IsRetryError
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
	DBConnection, DBConnectionError := sqlx.Connect("mysql", DBConnectionString)
	if DBConnectionError != nil {

		Error := fmt.Sprintf("Error Connecting To DB: %v", DBConnectionError.Error())

		IsRetryError := CheckConnectionError(DBConnectionError)

		if IsRetryError {

			for IsRetryError {

				// Generate Random Sleep Time
				util.SleepForRandomRange(1, 3)

				// Retry Execute DB Connect
				DBConnection, DBConnectionError = sqlx.Connect("mysql", DBConnectionString)

				if DBConnectionError != nil {

					IsRetryError = CheckConnectionError(DBConnectionError)

					if !IsRetryError {
						Error = fmt.Sprintf("Error Connecting To DB During Retry: %v", DBConnectionError.Error())
						logging.NewError(Error)
					}

				} else {

					IsRetryError = false

				}

			}

		} else {

			Error = fmt.Sprintf("Error Connecting To DB: %v", DBConnectionError.Error())
			logging.NewError(Error)

		}

	}

	// Return DB Connection
	return DBConnection

}
