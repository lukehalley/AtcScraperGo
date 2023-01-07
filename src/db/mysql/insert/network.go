package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	"fmt"
)

func AddNetworkToDB(NetworkName string, ChainNumber int, ChainRPCs []string, ExplorerAPIPrefix string, ExplorerAPIKey string, ExplorerTxURL string, ExplorerType string, GasSymbol string, GasAddress string, MinGas int, MaxGas int) int64 {

	// Make Sure We Only Put Max 5 RPCs In
	if len(ChainRPCs) > 5 {
		ChainRPCs = ChainRPCs[0:4]
	}

	// Build Our RPC Insert String
	var RPCSQLString string
	var RPCKeys string
	for Index, ChainRPC := range ChainRPCs {
		IndexPlus := Index + 1

		IsLastRPC := len(ChainRPCs) == IndexPlus

		RPCSQLString = RPCSQLString + fmt.Sprintf("'%v' AS chain_rpc_%d", ChainRPC, IndexPlus)
		RPCKeys = RPCKeys + fmt.Sprintf("chain_rpc_%d", IndexPlus)

		if !IsLastRPC {
			RPCSQLString = RPCSQLString + ", "
			RPCKeys = RPCKeys + ", "
		}

	}

	DBKeys := fmt.Sprintf("name, chain_number, %v, explorer_api_prefix, explorer_api_key, explorer_tx_url, explorer_type, gas_symbol, gas_address, max_gas, min_gas", RPCKeys)
	SelectStatement := fmt.Sprintf("(SELECT '%v' AS name, %d AS chain_number, %v, '%v' AS explorer_api_prefix, '%v' AS explorer_api_key, '%v' AS explorer_tx_url, '%v' AS explorer_type, '%v' AS gas_symbol, '%v' AS gas_address, %d AS max_gas, %d AS min_gas)", NetworkName, ChainNumber, RPCSQLString, ExplorerAPIPrefix, ExplorerAPIKey, ExplorerTxURL, ExplorerType, GasSymbol, GasAddress, MinGas, MaxGas)
	CompareStatement := fmt.Sprintf("networks.name = '%v' AND networks.chain_number = %d", NetworkName, ChainNumber)

	InsertQuery := "INSERT INTO networks(" + DBKeys + ") SELECT * FROM " + SelectStatement + " AS tmp WHERE NOT EXISTS (SELECT * FROM networks WHERE " + CompareStatement + ") LIMIT 1"

	InsertedNetworkID := mysqlutils.ExecuteInsert(InsertQuery)

	return InsertedNetworkID

}

func AddBlacklistNetworkToDB(NetworkName string, ChainNumber int) int64 {

	DBKeys := fmt.Sprintf("name, chain_number")
	SelectStatement := fmt.Sprintf("(SELECT '%v' AS name, %d AS chain_number)", NetworkName, ChainNumber)
	CompareStatement := fmt.Sprintf("blacklist_networks.name = '%v' AND blacklist_networks.chain_number = %d", NetworkName, ChainNumber)

	InsertQuery := "INSERT INTO blacklist_networks(" + DBKeys + ") SELECT * FROM " + SelectStatement + " AS tmp WHERE NOT EXISTS (SELECT * FROM blacklist_networks WHERE " + CompareStatement + ") LIMIT 1"

	InsertedNetworkID := mysqlutils.ExecuteInsert(InsertQuery)

	return InsertedNetworkID

}
