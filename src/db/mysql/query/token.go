package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	logging "atcscraper/src/log"
	"atcscraper/src/types/mysql"
	"fmt"
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
		Error := fmt.Sprintf("Error Querying DB For Token: %v", QueryError.Error())
		logging.NewError(Error)
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		Error := fmt.Sprintf("Error Closing DB Connecting: %v", DBConnectionCloseError.Error())
		logging.NewError(Error)
	}

	return Token

}