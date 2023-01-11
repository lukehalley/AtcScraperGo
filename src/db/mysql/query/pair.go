package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	logging "atcscraper/src/log"
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
		Error := fmt.Sprintf("Error Querying Pairs Table: %v", QueryError.Error())
		logging.NewError(Error)
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		Error := fmt.Sprintf("Error Closing DB Connecting: %v", DBConnectionCloseError.Error())
		logging.NewError(Error)
	}

	return Pair

}

func GetBlacklistPairFromDB(PairAddress string, NetworkId int) []mysql.BlacklistPair {

	// Query
	GetBlacklistPairQuery := fmt.Sprintf("SELECT blacklist_pairs.* FROM blacklist_pairs WHERE blacklist_pairs.network_id = %d AND blacklist_pairs.address = '%v'", NetworkId, PairAddress)

	// Create Connection To DB
	DBConnection := mysqlutils.CreateDatabaseConnection()

	// Create List Of BlacklistPair
	var BlacklistPair []mysql.BlacklistPair

	// Execute DB Query
	QueryError := DBConnection.Select(&BlacklistPair, GetBlacklistPairQuery)

	// Catch Errors
	if QueryError != nil {
		log.Println("Error Querying Blacklist Pairs Table!")
		log.Panicln("Error: ", QueryError.Error())
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		Error := fmt.Sprintf("Error Closing DB Connecting: %v", DBConnectionCloseError.Error())
		logging.NewError(Error)
	}

	return BlacklistPair

}

func GetBlacklistedPairsForNetwork(NetworkId int) []string {

	// Query
	GetBlacklistedPairsQuery := fmt.Sprintf("SELECT blacklist_pairs.* FROM blacklist_pairs WHERE blacklist_pairs.network_id = %d", NetworkId)

	// Create Connection To DB
	DBConnection := mysqlutils.CreateDatabaseConnection()

	// Create List Of BlacklistPairs
	var BlacklistPairs []mysql.BlacklistPair

	// Execute DB Query
	QueryError := DBConnection.Select(&BlacklistPairs, GetBlacklistedPairsQuery)

	// Catch Any Errors When Querying
	if QueryError != nil {
		Error := fmt.Sprintf("Error Querying DB For Blacklisted Pairs For Network: %v", QueryError.Error())
		logging.NewError(Error)
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		Error := fmt.Sprintf("Error Closing DB Connecting: %v", DBConnectionCloseError.Error())
		logging.NewError(Error)
	}

	// Get All Blacklisted Pair Addresses
	var BlacklistPairAddresses []string
	for _, BlacklistPair := range BlacklistPairs {
		BlacklistPairAddresses = append(BlacklistPairAddresses, BlacklistPair.Address)
	}

	return BlacklistPairAddresses

}