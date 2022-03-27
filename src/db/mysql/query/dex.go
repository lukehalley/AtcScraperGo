package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	"atcscraper/src/types/mysql"
	"fmt"
	"log"
)

func GetDexFromDB(RouterAddress string, FactoryAddress string, NetworkId int) mysql.Dex {

	// Query
	GetDexQuery := fmt.Sprintf("SELECT dexs.* FROM dexs WHERE dexs.network_id = %d AND dexs.router_address = '%v' AND dexs.factory_address = '%v'", NetworkId, RouterAddress, FactoryAddress)

	// Create Connection To DB
	DBConnection := mysqlutils.CreateDatabaseConnection()

	// Create List Of Dex
	var Dex []mysql.Dex

	// Execute DB Query
	DBConnection.Select(&Dex, GetDexQuery)

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		log.Fatal(DBConnectionCloseError)
	}

	if len(Dex) > 1 {
		log.Fatal("Couldn't Find Dex In DB:", RouterAddress, FactoryAddress, NetworkId)
	}

	return Dex[0]

}