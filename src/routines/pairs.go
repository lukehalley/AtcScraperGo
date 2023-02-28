package routines

import (
	geckoterminal_api "atcscraper/src/api/geckoterminal/requests"
	mysql_insert "atcscraper/src/db/mysql/insert"
	mysql_query "atcscraper/src/db/mysql/query"
	"atcscraper/src/io"
	geckoterminal_types "atcscraper/src/types/geckoterminal"
	"atcscraper/src/web3"
	"github.com/ethereum/go-ethereum/core/types"
	"sync"
)

func CollectGeckoTerminalDexPairs(Network geckoterminal_types.GeckoTerminalNetworkWithDexs, NetworkStablecoins []string, Dex geckoterminal_types.Dex, InvalidDexs []string, BlacklistPairAddresses []string, PagesToCollect int, TxsToCollect int, DexCollectionWaitGroup *sync.WaitGroup, DexCollectionChannel chan geckoterminal_types.Dex) {

	// Schedule The Call To WaitGroup's Done To Tell GoRoutine Is Completed.
	defer DexCollectionWaitGroup.Done()

	// Get First Page Dexs
	DexPairs := geckoterminal_api.GetGeckoterminalDexPairs(Network.Network.Identifier, Dex.Identifier, 1)

	// If We Want To Collect More Than One Page
	if PagesToCollect > 1 {

		for i := 2; i <= PagesToCollect; i++ {

			// Get All Pairs For Current Dex
			CurrentPageDexPairs := geckoterminal_api.GetGeckoterminalDexPairs(Network.Network.Identifier, Dex.Identifier, i)

			// Add The Rest Of The Pairs To The List
			if len(CurrentPageDexPairs.Data) > 0 {
				for _, CurrentPageDexPair := range CurrentPageDexPairs.Data {
					DexPairs.Data = append(DexPairs.Data, CurrentPageDexPair)
				}
			} else {
				break
			}

		}

	}
	
	// Get Dex Factory & Router Address
	DexContractsReceived := false
	RouterAbiObtained := false
	FactoryAbiObtained := false
	for _, DexPair := range DexPairs.Data {

		// Create New Pair Object
		var Pair geckoterminal_types.Pair

		// Add The Pair Address
		Pair.Address = DexPair.Attributes.Address

		// Get Current Pair Transactions
		PairTransactions := geckoterminal_api.GetGeckoterminalPairTransactionsAndTokens(Network.Network.Identifier, Pair.Address, 1)

		// Collect Transactions
		for _, Transaction := range PairTransactions.Data {
			var PairTransaction geckoterminal_types.Transaction
			PairTransaction.Hash = Transaction.Attributes.TxHash
			Pair.Transactions = append(Pair.Transactions, PairTransaction)
		}

		// Get Factory Address
		if Dex.FactoryAddress == "" {
			Dex.FactoryAddress = web3.GetPairFactoryAddress(DexPair.Attributes.Address, Network.RPCs[0])
		}

		// Get Router Address
		var TransactionReceipt *types.Transaction
		var TransactionReceived bool
		for _, PairTransaction := range Pair.Transactions {
			for _, RPCUrl := range Network.RPCs {
				TransactionReceived, TransactionReceipt = web3.GetTransactionReceipt(RPCUrl, PairTransaction.Hash)
				if TransactionReceived {
					Dex.RouterAddress = TransactionReceipt.To().Hex()
				}
				if Dex.RouterAddress != "" {
					break
				}
			}
			if Dex.RouterAddress != "" {
				break
			}
		}

		if Dex.RouterAddress != "" && Dex.FactoryAddress != "" {

			// Check If Dex Is Already Stored
			DexQueryResults := mysql_query.GetDexFromDB(Network.NetworkDBId, Dex.Identifier)

			// Get Dex DB ID
			if len(DexQueryResults) > 0 {

				// Set Dex DB ID
				Network.DexDBId = DexQueryResults[0].DexId

			} else {

				// Add Dex To DB
				Network.DexDBId = int(mysql_insert.AddDexToDB(
					Network.NetworkDBId,
					Dex.Identifier,
					Dex.RouterAddress,
					Dex.FactoryAddress,
					1,
				))

			}

			if Network.DexDBId > 0 {

				// Check If Router Has Been Stored
				DBDexRouterABIQuery := mysql_query.GetAbiFromDBByAddress(Network.NetworkDBId, Dex.RouterAddress)
				RouterABIStored := len(DBDexRouterABIQuery) > 0

				// Get Router Dex ABI
				if !RouterABIStored {
					Dex.RouterAbi, RouterAbiObtained = CollectDexsABI(Network, Dex.RouterAddress)
					if !RouterAbiObtained {
						Dex.RouterAbi = io.LoadAbiFromFile("IUniswapV2Router02.json")
					}
					Dex.RouterAbiDBId = mysql_insert.AddABIToDB(Network.NetworkDBId, Network.DexDBId, "router", Dex.RouterAddress, Dex.RouterAbi, "")
				} else {
					Dex.RouterAbi = DBDexRouterABIQuery[0].Abi
					Dex.RouterAbiDBId = int64(DBDexRouterABIQuery[0].AbiId)
				}

				// Check If Factory Has Been Stored
				DBDexFactoryABIQuery := mysql_query.GetAbiFromDBByAddress(Network.NetworkDBId, Dex.FactoryAddress)
				FactoryABIStored := len(DBDexFactoryABIQuery) > 0

				// Get Factory Dex ABI
				if !FactoryABIStored {
					Dex.FactoryAbi, FactoryAbiObtained = CollectDexsABI(Network, Dex.FactoryAddress)
					if !FactoryAbiObtained {
						Dex.FactoryAbi = io.LoadAbiFromFile("IUniswapV2Factory.json")
					}
					Dex.FactoryAbiDBId = mysql_insert.AddABIToDB(Network.NetworkDBId, Network.DexDBId, "factory", Dex.FactoryAddress, Dex.FactoryAbi, "")
				} else {
					Dex.FactoryAbi = DBDexFactoryABIQuery[0].Abi
					Dex.FactoryAbiDBId = int64(DBDexFactoryABIQuery[0].AbiId)
				}

				DexContractsReceived = Network.DexDBId > 0 && Dex.RouterAddress != "" && Dex.RouterAbi != "" && Dex.FactoryAddress != "" && Dex.FactoryAbi != ""

				if DexContractsReceived {
					break
				}

			}

		}
		
	}

	// Iterate Through Networks Dexs And Get All Their Pairs
	if DexContractsReceived {
		for _, DexPair := range DexPairs.Data {
			ScrapePairInfo(Network, NetworkStablecoins, Dex, DexPair, BlacklistPairAddresses, TxsToCollect)
		}
	}

	DexCollectionChannel <- Dex
	
}
