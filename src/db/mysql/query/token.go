package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	"atcscraper/src/types/mysql"
	"fmt"
	"log"
)


func GetStablecoinsFromDB() []mysql.Token {

	// Query
	GetStablecoinsQuery := "SELECT tokens.* FROM tokens WHERE tokens.stablecoin"

	// Create Connection To DB
	DBConnection := mysqlutils.CreateDatabaseConnection()

	// Create List Of Stablecoins
	var Stablecoins []mysql.Token

	// Execute DB Query
	QueryError := DBConnection.Select(&Stablecoins, GetStablecoinsQuery)

	// Catch Any Errors When Querying
	if QueryError != nil {
		log.Fatal(QueryError)
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		log.Fatal(DBConnectionCloseError)
	}

	return Stablecoins

}

func GetTokenFromDB(NetworkId int, TokenAddress string) mysql.Token {

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
		log.Fatal(QueryError)
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		log.Fatal(DBConnectionCloseError)
	}

	if len(Token) > 1 {
		return Token[0]
	} else {
		var NullToken mysql.Token
		NullToken.TokenId = 0
		return NullToken
	}

}