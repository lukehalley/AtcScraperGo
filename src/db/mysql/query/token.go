package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	"atcscraper/src/types/mysql"
	"fmt"
	"log"
)

func GetTokenFromDB(NetworkId int, TokenSymbol string, TokenAddress string) []mysql.Token {

	// Query
	GetTokenQuery := fmt.Sprintf("SELECT tokens.* FROM tokens WHERE tokens.network_id = %d AND tokens.address = '%v' AND tokens.symbol = '%v'", NetworkId, TokenAddress, TokenSymbol)

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