package mysql_utils

import (
	"context"
	"log"
)

func ExecuteInsert(InsertQuery string) int64 {

	// Create Connection To DB
	DBConnection := CreateDatabaseConnection()

	// Execute DB Query
	QueryResult, _ := DBConnection.ExecContext(context.Background(), InsertQuery)

	// Get Inserted Row ID
	InsertedRowID := int64(0)
	if QueryResult != nil {
		InsertedRowID, _ = QueryResult.LastInsertId()
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		log.Fatal(DBConnectionCloseError)
	}

	return InsertedRowID

}