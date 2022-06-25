package routines

import (
	"atcscraper/src/api/explorer"
	query "atcscraper/src/db/mysql/insert"
	mysql_query "atcscraper/src/db/mysql/query"
	geckoterminal_types "atcscraper/src/types/geckoterminal"
	"atcscraper/src/types/web3"
	"log"
	"sync"
	"time"
)

func CollectNetworkDexsABIs(Network geckoterminal_types.GeckoTerminalNetworkWithDexs, NetworkABIsWaitGroup *sync.WaitGroup, NetworkDexABIsChannel chan geckoterminal_types.GeckoTerminalNetworkWithDexs) {

	// Schedule The Call To WaitGroup's Done To Tell GoRoutine Is Completed.
	defer NetworkABIsWaitGroup.Done()

	// Get Network DB Details
	DBNetworkQuery := mysql_query.GetNetwork(Network.Network.Identifier)

	// To Check If We Loaded Any ABIs
	LoadedAbis := false

	// Rate Limit
	IterationCount := 0

	// Iterate Through Networks Dexs And Get All Their Pairs
	for DexIndex, Dex := range Network.Dexes {

		if len(DBNetworkQuery) > 0 {

			// Get The Results
			NetworkDetails := DBNetworkQuery[0]

			// Check If API Is An Etherscan or Blockscout Type
			if NetworkDetails.ExplorerType.String == "scan" || NetworkDetails.ExplorerType.String == "blockscout" {

				// Get Network DB Details
				DBDexQuery := mysql_query.GetDexFromDB(NetworkDetails.NetworkId, Dex.Identifier)

				if len(DBDexQuery) > 0 {

					// Get Dex Details From Query
					DexDetails := DBDexQuery[0]

					// Check If Router Has Been Stored
					DBDexRouterABIQuery := mysql_query.GetAbiFromDBByAddress(NetworkDetails.NetworkId, DexDetails.DexRouterAddress.String)
					RouterABIStored := len(DBDexRouterABIQuery) > 0

					// Add The Router ABI To The DB If It Isn't Already In DB
					if !RouterABIStored {

						if DexDetails.DexRouterAddress.String != "" {

							// Fetch Router ABI
							FetchedRouterABI := web3.AbiAPI{}
							WasSuccessful := false
							if NetworkDetails.ExplorerApiPrefix.String != "" {
								FetchedRouterABI, WasSuccessful = explorer.GetAbiFromExplorer(NetworkDetails.ExplorerApiPrefix.String, NetworkDetails.ExplorerApiKey.String, DexDetails.DexRouterAddress.String)
							} else {
								FetchedRouterABI, WasSuccessful = explorer.GetAbiFromExplorer(NetworkDetails.ExplorerTxUrl.String, "", DexDetails.DexRouterAddress.String)
							}

							IterationCount = IterationCount + 1

							// Add To Router ABI To DB If Successful
							if WasSuccessful {

								// Add Router ABI To DB
								AddedAbiDBId := query.AddABIToDB(Network.NetworkDBId, DexDetails.DexId, "router", DexDetails.DexRouterAddress.String, FetchedRouterABI.Result, "")

								// Get The Added ABI
								DBDexRouterABIQuery = mysql_query.GetAbiFromDBById(int(AddedAbiDBId))

								// Add The Added Router ABI To The Dex Struct
								if len(DBDexRouterABIQuery) > 0 {

									// Set Our Dex Objects ABI
									Network.Dexes[DexIndex].RouterAbi = DBDexRouterABIQuery[0]

									// Set That Abi Was Loaded
									LoadedAbis = true

								}

							}

						}

					} else {

						// Add The Existing Router ABI To The Dex Struct
						Network.Dexes[DexIndex].RouterAbi = DBDexRouterABIQuery[0]

						// Set That Abi Was Loaded
						LoadedAbis = true

					}

					// Check If Factory Has Been Stored
					DBDexFactoryABIQuery := mysql_query.GetAbiFromDBByAddress(NetworkDetails.NetworkId, DexDetails.DexFactoryAddress.String)
					FactoryABIStored := len(DBDexFactoryABIQuery) > 0

					// Add The Factory ABI To The DB If It Isn't Already In DB
					if !FactoryABIStored {

						if DexDetails.DexFactoryAddress.String != "" {

							// Fetch Factory ABI
							FetchedFactoryABI := web3.AbiAPI{}
							WasSuccessful := false
							if NetworkDetails.ExplorerApiPrefix.String != "" {
								FetchedFactoryABI, WasSuccessful = explorer.GetAbiFromExplorer(NetworkDetails.ExplorerApiPrefix.String, NetworkDetails.ExplorerApiKey.String, DexDetails.DexFactoryAddress.String)
							} else {
								FetchedFactoryABI, WasSuccessful = explorer.GetAbiFromExplorer(NetworkDetails.ExplorerTxUrl.String, "", DexDetails.DexFactoryAddress.String)
							}

							IterationCount = IterationCount + 1

							// Add Factory ABI To DB If Successful
							if WasSuccessful {

								// Add Factory ABI To DB
								AddedAbiDBId := query.AddABIToDB(Network.NetworkDBId, DexDetails.DexId, "factory", DexDetails.DexFactoryAddress.String, FetchedFactoryABI.Result, "")

								// Get The Added ABI
								DBDexFactoryABIQuery = mysql_query.GetAbiFromDBById(int(AddedAbiDBId))

								// Add The Added Factory ABI To The Dex Struct
								if len(DBDexFactoryABIQuery) > 0 {

									// Add The Existing Router ABI To The Dex Struct
									Network.Dexes[DexIndex].FactoryAbi = DBDexFactoryABIQuery[0]

									// Set That Abi Was Loaded
									LoadedAbis = true

								}
							}

						}

					} else {

						// Add The Existing Factory ABI To The Dex Struct
						Network.Dexes[DexIndex].FactoryAbi = DBDexFactoryABIQuery[0]

						// Set That Abi Was Loaded
						LoadedAbis = true

					}

				}

			}

		}

		// Rate Limit
		if IterationCount >= 4 {
			time.Sleep(time.Duration(1) * time.Second)
			IterationCount = 0
		}

	}

	if LoadedAbis {
		// Log Network Dex Load
		log.Printf("[%v] Loaded ABIs", Network.Network.Name)
	}

	NetworkDexABIsChannel <- Network

}

func ScrapeNetworkDexs(Network geckoterminal_types.GeckoTerminalNetworkWithDexs, InvalidDexs []string, PagesToCollect int, TxsToCollect int, DexCollectionWaitGroup *sync.WaitGroup, DexCollectionChannel chan geckoterminal_types.GeckoTerminalNetworkWithDexs) {

	// Schedule The Call To WaitGroup's Done To Tell GoRoutine Is Completed.
	defer DexCollectionWaitGroup.Done()

	// Get All Network Blacklist Pairs So We Can Skip Them
	BlacklistPairAddresses := mysql_query.GetBlacklistedPairsForNetwork(Network.NetworkDBId)

	// Collect Network Stablecoin Addresses
	var NetworkStablecoins []string
	for _, Stablecoin := range Network.Stablecoins {
		NetworkStablecoins = append(NetworkStablecoins, Stablecoin.Address)
	}

	// Create Concurrency Objects
	DexPairsCollectionWaitGroup := new(sync.WaitGroup)
	DexPairsCollectionWaitGroup.Add(len(Network.Dexes))
	DexPairsCollectionChannel := make(chan geckoterminal_types.Dex, len(Network.Dexes))
	
	// Iterate Through Networks Dexs And Get All Their Pairs
	for _, Dex := range Network.Dexes {
		go CollectGeckoTerminalDexPairs(Network, NetworkStablecoins, Dex, InvalidDexs, BlacklistPairAddresses, PagesToCollect, TxsToCollect, DexPairsCollectionWaitGroup, DexPairsCollectionChannel)
	}

	// Wait For All Networks To Come Back
	DexPairsCollectionWaitGroup.Wait()

	// Close The Group Channel
	close(DexPairsCollectionChannel)

	DexCollectionChannel <- Network

}
