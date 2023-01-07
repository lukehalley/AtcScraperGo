package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	"atcscraper/src/types/mysql"
	"fmt"
	"log"
)


func GetNetworkStablecoinsFromDB(NetworkDBId int64) []mysql.Stablecoin {

	// Query
	GetStablecoinsQuery := fmt.Sprintf("SELECT stablecoins.* FROM stablecoins WHERE stablecoins.network_id = %d", NetworkDBId)

	// Create Connection To DB
	DBConnection := mysqlutils.CreateDatabaseConnection()

	// Create List Of Stablecoins
	var Stablecoins []mysql.Stablecoin

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