package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	logging "atcscraper/src/log"
	"atcscraper/src/types/mysql"
	"fmt"
)

func GetAbiFromDBByAddress(NetworkId int, ContractAddress string) []mysql.Abi {

	// Query
	GetAbiQuery := fmt.Sprintf("SELECT abis.* FROM abis WHERE abis.network_id = %d AND abis.contract_address = '%v'", NetworkId, ContractAddress)

	// Create Connection To DB
	DBConnection := mysqlutils.CreateDatabaseConnection()

	// Create List Of Abi
	var Abi []mysql.Abi

	// Execute DB Query
	QueryError := DBConnection.Select(&Abi, GetAbiQuery)

	// Catch Any Errors When Querying
	if QueryError != nil {
		Error := fmt.Sprintf("Error Querying DB For Abi: %v", QueryError.Error())
		logging.NewError(Error)
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		Error := fmt.Sprintf("Error Closing DB Connecting: %v", DBConnectionCloseError.Error())
		logging.NewError(Error)
	}

	return Abi

}

func GetAbiFromDBById(AbiId int) []mysql.Abi {

	// Query
	GetAbiQuery := fmt.Sprintf("SELECT abis.* FROM abis WHERE abis.abi_id = %d", AbiId)

	// Create Connection To DB
	DBConnection := mysqlutils.CreateDatabaseConnection()

	// Create List Of Abi
	var Abi []mysql.Abi

	// Execute DB Query
	QueryError := DBConnection.Select(&Abi, GetAbiQuery)

	// Catch Any Errors When Querying
	if QueryError != nil {
		Error := fmt.Sprintf("Error Querying DB For Abi: %v", QueryError.Error())
		logging.NewError(Error)
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		Error := fmt.Sprintf("Error Closing DB Connecting: %v", DBConnectionCloseError.Error())
		logging.NewError(Error)
	}

	return Abi

}