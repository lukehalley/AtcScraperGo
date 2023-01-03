package mysql_utils

import (
	"context"
	"log"
)

func ExecuteInsert(InsertQuery string) int64 {

	// Create Connection To DB
	DBConnection := CreateDatabaseConnection()

	// Execute DB Query
	QueryResult, QueryError := DBConnection.ExecContext(context.Background(), InsertQuery)

	// Catch Insert Error
	if QueryError != nil {
		log.Fatal("Error Inserting DB Record: ", QueryError)
	}

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