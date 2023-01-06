package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	"atcscraper/src/types/mysql"
	"fmt"
	"log"
)

func GetDexFromDB(NetworkId int, DexName string) []mysql.Dex {

	// Query
	GetDexQuery := fmt.Sprintf("SELECT dexs.* FROM dexs WHERE dexs.network_id = %d AND dexs.name = '%v'", NetworkId, DexName)

	// Create Connection To DB
	DBConnection := mysqlutils.CreateDatabaseConnection()

	// Create List Of Dex
	var Dex []mysql.Dex

	// Execute DB Query
	QueryError := DBConnection.Select(&Dex, GetDexQuery)

	// Catch Any Errors When Querying
	if QueryError != nil {
		log.Fatal(QueryError)
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		log.Fatal(DBConnectionCloseError)
	}

	return Dex

}