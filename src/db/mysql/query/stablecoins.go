package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	logging "atcscraper/src/log"
	"atcscraper/src/types/mysql"
	"fmt"
)


func GetNetworkStablecoinsFromDB(NetworkName string) []mysql.Stablecoin {

	// Query
	GetStablecoinsQuery := fmt.Sprintf("SELECT stablecoins.* FROM stablecoins WHERE stablecoins.network_name = '%v'", NetworkName)

	// Create Connection To DB
	DBConnection := mysqlutils.CreateDatabaseConnection()

	// Create List Of Stablecoins
	var Stablecoins []mysql.Stablecoin

	// Execute DB Query
	QueryError := DBConnection.Select(&Stablecoins, GetStablecoinsQuery)

	// Catch Any Errors When Querying
	if QueryError != nil {
		Error := fmt.Sprintf("Error Querying DB For Network Stablecoins: %v", QueryError.Error())
		logging.NewError(Error)
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		Error := fmt.Sprintf("Error Closing DB Connecting: %v", DBConnectionCloseError.Error())
		logging.NewError(Error)
	}

	return Stablecoins

}