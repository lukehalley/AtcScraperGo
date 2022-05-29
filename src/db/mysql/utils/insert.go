package mysql_utils

import (
	logging "atcscraper/src/log"
	"context"
	"fmt"
)

func ExecuteInsert(InsertQuery string) int64 {

	// Create Connection To DB
	DBConnection := CreateDatabaseConnection()

	// Execute DB Query
	InsertResult, InsertError := DBConnection.ExecContext(context.Background(), InsertQuery)

	// Catch Insert Error
	if InsertError != nil {
		Error := fmt.Sprintf("Error Inserting DB Record: %v", InsertError)
		logging.NewError(Error)
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