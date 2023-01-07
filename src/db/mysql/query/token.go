package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	"atcscraper/src/types/mysql"
	"fmt"
	"log"
)

func GetTokenFromDB(NetworkId int, TokenAddress string) []mysql.Token {

	// Query
	GetTokenQuery := fmt.Sprintf("SELECT tokens.* FROM tokens WHERE tokens.network_id = %d AND tokens.address = '%v'", NetworkId, TokenAddress)

	// Create Connection To DB
	DBConnection := mysqlutils.CreateDatabaseConnection()

	// Create List Of Token
	var Token []mysql.Token

	// Execute DB Query
	QueryError := DBConnection.Select(&Token, GetTokenQuery)

	// Catch Any Errors When Querying
	if QueryError != nil {
		log.Fatal("Error Querying DB: ", QueryError)
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		log.Fatal(DBConnectionCloseError)
	}

	return Token

}