package mysql_utils

import (
	logging "atcscraper/src/log"
	"atcscraper/src/util"
	"context"
	"fmt"
	"strings"
)

func ExecuteInsert(InsertQuery string) int64 {

	// Create Connection To DB
	DBConnection := CreateDatabaseConnection()

	// Execute DB Query
	InsertResult, InsertError := DBConnection.ExecContext(context.Background(), InsertQuery)

	// Catch Insert Error
	if InsertError != nil {

		DeadlockError := strings.Contains(InsertError.Error(), "Deadlock")

		if DeadlockError {

			for DeadlockError {

				// Close Connection
				DBConnectionCloseError := DBConnection.Close()
				if DBConnectionCloseError != nil {
					Error := fmt.Sprintf("Error Closing DB Connecting During Deadlock: %v", DBConnectionCloseError.Error())
					logging.NewError(Error)
				}

				// Generate Random Sleep Time
				util.SleepForRandomRange(1, 3)

				// Create Connection To DB
				DBConnection = CreateDatabaseConnection()

				// Retry Execute DB Query
				InsertResult, InsertError = DBConnection.ExecContext(context.Background(), InsertQuery)

				if InsertError != nil {

					DeadlockError = strings.Contains(InsertError.Error(), "Deadlock")

					if !DeadlockError {
						Error := fmt.Sprintf("Error Inserting DB Record During Deadlock: %v", InsertError)
						logging.NewError(Error)
					}

				}

			}
		} else {
			Error := fmt.Sprintf("Error Inserting DB Record: %v", InsertError)
			logging.NewError(Error)
		}

	}

	// Get Inserted Row ID
	InsertedRowID := int64(0)
	if InsertResult != nil {
		InsertedRowID, _ = InsertResult.LastInsertId()
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		Error := fmt.Sprintf("Error Closing DB Connecting: %v", DBConnectionCloseError.Error())
		logging.NewError(Error)
	}

	return InsertedRowID

}