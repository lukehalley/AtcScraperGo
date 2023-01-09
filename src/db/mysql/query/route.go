package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	logging "atcscraper/src/log"
	"atcscraper/src/types/mysql"
	"fmt"
	"log"
)

func GetRouteFromDB(NetworkId int, DexId int, PairId int, TxHash string) []mysql.Route {

	// Query
	GetRouteQuery := fmt.Sprintf("SELECT routes.* FROM routes WHERE routes.network_id = %d AND routes.dex_id = %d AND routes.pair_id = %d AND routes.transaction_hash = '%v'", NetworkId, DexId, PairId, TxHash)

	// Create Connection To DB
	DBConnection := mysqlutils.CreateDatabaseConnection()

	// Create List Of Route
	var Route []mysql.Route

	// Execute DB Query
	QueryError := DBConnection.Select(&Route, GetRouteQuery)

	// Catch Errors
	if QueryError != nil {
		log.Println("Couldn't Find Route In DB!")
		log.Panicln("Error: ", QueryError.Error())
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		Error := fmt.Sprintf("Error Closing DB Connecting: %v", DBConnectionCloseError.Error())
		logging.NewError(Error)
	}

	return Route

}