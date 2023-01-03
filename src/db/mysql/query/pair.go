package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	"atcscraper/src/types/mysql"
	"fmt"
	"log"
)

func GetPairFromDB(PairAddress string, NetworkId int) []mysql.Pair {

	// Query
	GetPairQuery := fmt.Sprintf("SELECT pairs.* FROM pairs WHERE pairs.network_id = %d AND pairs.address = '%v'", NetworkId, PairAddress)

	// Create Connection To DB
	DBConnection := mysqlutils.CreateDatabaseConnection()

	// Create List Of Pair
	var Pair []mysql.Pair

	// Execute DB Query
	QueryError := DBConnection.Select(&Pair, GetPairQuery)

	// Catch Errors
	if QueryError != nil {
		log.Println("Couldn't Find Pair In DB With Address: ", PairAddress, " & NetworkId: ", NetworkId)
		log.Panicln("Error: ", QueryError.Error())
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		log.Fatal(DBConnectionCloseError)
	}

	return Pair

}