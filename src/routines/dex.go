package routines

import (
	"atcscraper/src/api/explorer"
	mysql_query "atcscraper/src/db/mysql/query"
	geckoterminal_types "atcscraper/src/types/geckoterminal"
	"atcscraper/src/types/web3"
	"sync"
)

func CollectDexsABI(Network geckoterminal_types.GeckoTerminalNetworkWithDexs, ContractAddress string) string {

	// Get Network DB Details
	DBNetworkQuery := mysql_query.GetNetwork(Network.Network.Identifier)
	NetworkDetails := DBNetworkQuery[0]

	// Blank Abi Strings
	CollectedAbi := ""

	// Check If API Is An Etherscan or Blockscout Type
	if NetworkDetails.ExplorerType.String == "scan" || NetworkDetails.ExplorerType.String == "blockscout" {

		// Get The Router ABI If We Have The Address
		if ContractAddress != "" {

			// Fetch Router ABI
			FetchedABI := web3.AbiAPI{}
			WasSuccessful := false
			if NetworkDetails.ExplorerApiPrefix.String != "" {
				FetchedABI, WasSuccessful = explorer.GetAbiFromExplorer(NetworkDetails.ExplorerApiPrefix.String, NetworkDetails.ExplorerApiKey.String, ContractAddress)
			} else {
				FetchedABI, WasSuccessful = explorer.GetAbiFromExplorer(NetworkDetails.ExplorerTxUrl.String, "", ContractAddress)
			}

			NotBlankABI := FetchedABI.Result != ""

			// Add To Router ABI To DB If Successful
			if WasSuccessful && NotBlankABI {

				// Set Our Dex Objects ABI
				CollectedAbi = FetchedABI.Result

			}

		}

	}

	return CollectedAbi

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
