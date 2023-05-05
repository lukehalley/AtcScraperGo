package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	logging "atcscraper/src/log"
	"atcscraper/src/types/mysql"
	"fmt"
	"log"
)

func GetDecodedTransactionFromDB(NetworkId int, DexId int, PairId int, TxHash string) []mysql.DecodedTransaction {

	// Query
	GetDecodedTransactionQuery := fmt.Sprintf("SELECT decoded_transactions.* FROM decoded_transactions WHERE decoded_transactions.network_id = %d AND decoded_transactions.dex_id = %d AND decoded_transactions.pair_id = %d AND decoded_transactions.transaction_hash = '%v'", NetworkId, DexId, PairId, TxHash)

	// Create Connection To DB
	DBConnection := mysqlutils.CreateDatabaseConnection()

	// Create List Of DecodedTransaction
	var DecodedTransaction []mysql.DecodedTransaction

	// Execute DB Query
	QueryError := DBConnection.Select(&DecodedTransaction, GetDecodedTransactionQuery)

	// Catch Errors
	if QueryError != nil {
		log.Println("Couldn't Find DecodedTransaction In DB!")
		log.Panicln("Error: ", QueryError.Error())
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		Error := fmt.Sprintf("Error Closing DB Connecting: %v", DBConnectionCloseError.Error())
		logging.NewError(Error)
	}

	return DecodedTransaction

}

func CheckIfPairHasDecodedTransactionInDB(NetworkId int, DexId int, PairId int) bool {

	// Query
	GetDecodedTransactionQuery := fmt.Sprintf("SELECT decoded_transactions.* FROM decoded_transactions WHERE decoded_transactions.network_id = %d AND decoded_transactions.dex_id = %d AND decoded_transactions.pair_id = %d", NetworkId, DexId, PairId)

	// Create Connection To DB
	DBConnection := mysqlutils.CreateDatabaseConnection()

	// Create List Of DecodedTransaction
	var DecodedTransaction []mysql.DecodedTransaction

	// Execute DB Query
	QueryError := DBConnection.Select(&DecodedTransaction, GetDecodedTransactionQuery)

	// Catch Errors
	if QueryError != nil {
		log.Println("Couldn't Find DecodedTransaction In DB!")
		log.Panicln("Error: ", QueryError.Error())
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		Error := fmt.Sprintf("Error Closing DB Connecting: %v", DBConnectionCloseError.Error())
		logging.NewError(Error)
	}

	return len(DecodedTransaction) > 0

}