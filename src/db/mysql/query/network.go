package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	logging "atcscraper/src/log"
	"atcscraper/src/types/mysql"
	"fmt"
)

func GetNetwork(NetworkName string) []mysql.Network {

	// Query
	GetNetworkQuery := fmt.Sprintf("SELECT networks.* FROM networks WHERE networks.name = '%v'", NetworkName)

	// Create Connection To DB
	DBConnection := mysqlutils.CreateDatabaseConnection()

	// Create List Of Networks
	var Networks []mysql.Network

	// Execute DB Query
	QueryError := DBConnection.Select(&Networks, GetNetworkQuery)

	// Catch Any Errors When Querying
	if QueryError != nil {
		Error := fmt.Sprintf("Error Querying DB For Network: %v", QueryError.Error())
		logging.NewError(Error)
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		Error := fmt.Sprintf("Error Closing DB Connecting: %v", DBConnectionCloseError.Error())
		logging.NewError(Error)
	}

	return Networks

}

func GetBlacklistNetwork(NetworkName string) []mysql.BlacklistNetwork {

	// Query
	GetBlacklistNetworksQuery := fmt.Sprintf("SELECT blacklist_networks.* FROM blacklist_networks WHERE blacklist_networks.name = '%v'", NetworkName)

	// Create Connection To DB
	DBConnection := mysqlutils.CreateDatabaseConnection()

	// Create List Of BlacklistNetworks
	var BlacklistNetworks []mysql.BlacklistNetwork

	// Execute DB Query
	QueryError := DBConnection.Select(&BlacklistNetworks, GetBlacklistNetworksQuery)

	// Catch Any Errors When Querying
	if QueryError != nil {
		Error := fmt.Sprintf("Error Querying DB For Blacklisted Network: %v", QueryError.Error())
		logging.NewError(Error)
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		Error := fmt.Sprintf("Error Closing DB Connecting: %v", DBConnectionCloseError.Error())
		logging.NewError(Error)
	}

	return BlacklistNetworks

}

func GetBlacklistedNetworks() []int {

	// Query
	GetBlacklistNetworks := fmt.Sprintf("SELECT blacklist_networks.* FROM blacklist_networks")

	// Create Connection To DB
	DBConnection := mysqlutils.CreateDatabaseConnection()

	// Create List Of BlacklistNetworks
	var BlacklistNetworks []mysql.BlacklistNetwork

	// Execute DB Query
	QueryError := DBConnection.Select(&BlacklistNetworks, GetBlacklistNetworks)

	// Catch Any Errors When Querying
	if QueryError != nil {
		Error := fmt.Sprintf("Error Querying DB For Blacklisted Networks: %v", QueryError.Error())
		logging.NewError(Error)
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		Error := fmt.Sprintf("Error Closing DB Connecting: %v", DBConnectionCloseError.Error())
		logging.NewError(Error)
	}

	// Get Blacklist Network Chain IDs
	var BlacklistNetworkChainIds []int
	for _, BlacklistNetwork := range BlacklistNetworks {
		BlacklistNetworkChainIds = append(BlacklistNetworkChainIds, BlacklistNetwork.ChainNumber)
	}

	return BlacklistNetworkChainIds

}

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
		Error := fmt.Sprintf("Error Querying DB For Bitquery Compatible Networks: %v", QueryError.Error())
		logging.NewError(Error)
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		Error := fmt.Sprintf("Error Closing DB Connecting: %v", DBConnectionCloseError.Error())
		logging.NewError(Error)
	}

	return Networks

}