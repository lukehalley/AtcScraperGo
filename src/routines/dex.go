package routines

import (
	"atcscraper/src/api/explorer"
	mysql_query "atcscraper/src/db/mysql/query"
	geckoterminal_types "atcscraper/src/types/geckoterminal"
	"atcscraper/src/types/web3"
	"fmt"
	"sync"
)

func CollectNetworkDexsABIs(Network geckoterminal_types.GeckoTerminalNetworkWithDexs, NetworkABIsWaitGroup *sync.WaitGroup, NetworkDexABIsChannel chan geckoterminal_types.GeckoTerminalNetworkWithDexs) {

	// Schedule The Call To WaitGroup's Done To Tell GoRoutine Is Completed.
	defer NetworkABIsWaitGroup.Done()

	// Get Network DB Details
	DBNetwork := mysql_query.GetNetwork(Network.Network.Identifier)

	// Create Concurrency Objects
	DexABIsCollectionWaitGroup := new(sync.WaitGroup)
	DexABIsCollectionWaitGroup.Add(len(Network.Dexes))
	DexABIsCollectionChannel := make(chan geckoterminal_types.Dex, len(Network.Dexes))

	// Iterate Through Networks Dexs And Get All Their Pairs
	for _, Dex := range Network.Dexes {
		if len(DBNetwork) > 0 {

			// Get The Results
			NetworkDetails := DBNetwork[0]

			// Check If API Is An Etherscan or Blockscout Type
			if NetworkDetails.ExplorerType.String == "scan" || NetworkDetails.ExplorerType.String == "blockscout" {

				// Get Network DB Details
				DBDex := mysql_query.GetDexFromDB(NetworkDetails.NetworkId, Dex.Identifier)

				if len(DBDex) > 0 {
					DexDetails := DBDex[0]
					if DexDetails.DexRouterAddress.String != "" && DexDetails.DexFactoryAddress.String != "" {

						FetchedABI := web3.AbiAPI{}
						WasSuccessful := false
						if NetworkDetails.ExplorerApiPrefix.String != "" {
							FetchedABI, WasSuccessful = explorer.GetAbiFromExplorer(NetworkDetails.ExplorerApiPrefix.String, NetworkDetails.ExplorerApiKey.String, DexDetails.DexRouterAddress.String)
						} else {
							FetchedABI, WasSuccessful = explorer.GetAbiFromExplorer(NetworkDetails.ExplorerTxUrl.String, "", DexDetails.DexRouterAddress.String)
						}

						if WasSuccessful {
							fmt.Printf("%v", FetchedABI.Message)
						}

					}

				}

			}

		}
	}

	// Wait For All Networks To Come Back
	DexABIsCollectionWaitGroup.Wait()

	// Close The Group Channel
	close(DexABIsCollectionChannel)

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
