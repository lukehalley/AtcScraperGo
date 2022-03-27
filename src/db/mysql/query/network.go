package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	"atcscraper/src/types/mysql"
	"log"
)


func GetBitqueryCompatibleNetworks() []mysql.Network {

	// Query
	GetBitqueryCompatibleNetworksQuery := "SELECT networks.* FROM networks WHERE networks.bitquery_compatible"

	// Create Connection To DB
	DBConnection := mysqlutils.CreateDatabaseConnection()

	// Create List Of Networks
	var Networks []mysql.Network

	// Execute DB Query
	QueryError := DBConnection.Select(&Networks, GetBitqueryCompatibleNetworksQuery)

	// Catch Any Errors When Querying
	if QueryError != nil {
		log.Fatal(QueryError)
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		log.Fatal(DBConnectionCloseError)
	}

	return Networks

}