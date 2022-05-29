package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	logging "atcscraper/src/log"
	"atcscraper/src/types/mysql"
	"fmt"
)

func GetInvalidDexs() []string {

	// Query
	GetInvalidDexQuery := fmt.Sprintf("SELECT dexs.* FROM dexs WHERE dexs.is_valid = 0")

	// Create Connection To DB
	DBConnection := mysqlutils.CreateDatabaseConnection()

	// Create List Of Dexs
	var Dexs []mysql.Dex

	// Execute DB Query
	QueryError := DBConnection.Select(&Dexs, GetInvalidDexQuery)

	// Catch Any Errors When Querying
	if QueryError != nil {
		Error := fmt.Sprintf("Error Querying DB For Invalid Dexs: %v", QueryError.Error())
		logging.NewError(Error)
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		Error := fmt.Sprintf("Error Closing DB Connecting: %v", DBConnectionCloseError.Error())
		logging.NewError(Error)
	}

	var InvalidDexs []string
	for _, Dex := range Dexs {
		InvalidDexs = append(InvalidDexs, Dex.Name)
	}

	return InvalidDexs

}

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
		Error := fmt.Sprintf("Error Querying DB For Dex: %v", QueryError.Error())
		logging.NewError(Error)
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		Error := fmt.Sprintf("Error Closing DB Connecting: %v", DBConnectionCloseError.Error())
		logging.NewError(Error)
	}

	return Dex

}