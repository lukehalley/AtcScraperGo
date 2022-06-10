package routines

import (
	chainlist "atcscraper/src/api/chainlist/requests"
	geckoterminal_api "atcscraper/src/api/geckoterminal/requests"
	mysql_insert "atcscraper/src/db/mysql/insert"
	mysql_query "atcscraper/src/db/mysql/query"
	"atcscraper/src/requests"
	geckoterminal_types "atcscraper/src/types/geckoterminal"
	"atcscraper/src/util"
	"log"
	"strings"
	"sync"
	"time"
)

func CollectGeckoTerminalNetworkWithDexs(Network geckoterminal_types.GeckoTerminalNetwork, BlacklistedNetworks []int, CoingeckoBuildID string, ChainlistBuildID string, WaitTime time.Duration, NetworkCollectionWaitGroup *sync.WaitGroup, NetworkCollectionChannel chan geckoterminal_types.GeckoTerminalNetworkWithDexs) {

	// Schedule The Call To WaitGroup's Done To Tell GoRoutine Is Completed.
	defer NetworkCollectionWaitGroup.Done()

	// Create Collection Object
	var NetworkWithDexsAndPairs geckoterminal_types.GeckoTerminalNetworkWithDexs

	// Check If Network Blacklisted
	NetworkIsBlacklisted, _ := util.CheckIfIntIsInList(BlacklistedNetworks, Network.Attributes.ChainID)

	if !NetworkIsBlacklisted {

		// log.Printf("[%d/%d] [%v] Getting RPCs From Chainlist...", CountIndex, NetworkCount, Network.Attributes.Name)

		// Get Chain RPCs
		ChainInfo := chainlist.GetChainInfo(Network.Attributes.ChainID, ChainlistBuildID)

		// Sleep For N Seconds
		time.Sleep(WaitTime)

		var FilteredRPCs []string
		for _, RPC := range ChainInfo.PageProps.Chain.RPC {
			if !(strings.Contains(RPC.URL, "API")) {
				HealthyRPC := requests.CheckURLIsReachable(RPC.URL)
				if HealthyRPC {
					FilteredRPCs = append(FilteredRPCs, RPC.URL)
				}

			}
		}

		// Check If We Have Any RPCs
		if len(FilteredRPCs) > 0 {

			// log.Printf("[%d/%d] [%v] Collected %d RPC(s)", CountIndex, NetworkCount, Network.Attributes.Name, len(FilteredRPCs))

			// Add RPCs
			for _, RPCUrl := range FilteredRPCs {
				NetworkWithDexsAndPairs.RPCs = append(NetworkWithDexsAndPairs.RPCs, RPCUrl)
			}

			// Add Network Details
			NetworkWithDexsAndPairs.Network = struct {
				Name                  string
				Identifier            string
				ChainID               int
				ExplorerURL           string
				NativeCurrencySymbol  string
				NativeCurrencyAddress string
				PoolReserveThreshold  string
				ImageURL              string
				ExplorerLogoURL       string
			} (Network.Attributes)

			// Check If Network Is Already Stored
			NetworkQueryResults := mysql_query.GetNetwork(NetworkWithDexsAndPairs.Network.Identifier)

			// Chain Explorer
			var ChainExplorer string
			if len(ChainInfo.PageProps.Chain.Explorers) > 0 {
				ChainExplorer = ChainInfo.PageProps.Chain.Explorers[0].Standard
			} else {
				ChainExplorer = "none"
			}

			var NetworkDBID int64
			if len(NetworkQueryResults) > 0 {
				// Get Network DB ID
				NetworkDBID = int64(NetworkQueryResults[0].NetworkId)
			} else {
				// Add Network To DB
				NetworkDBID = mysql_insert.AddNetworkToDB(
					NetworkWithDexsAndPairs.Network.Identifier,
					NetworkWithDexsAndPairs.Network.ChainID,
					FilteredRPCs,
					"",
					"",
					ChainExplorer,
					NetworkWithDexsAndPairs.Network.ExplorerURL,
					NetworkWithDexsAndPairs.Network.NativeCurrencySymbol,
					NetworkWithDexsAndPairs.Network.NativeCurrencyAddress,
					1,
					5,
				)
			}

			// Set Network DB ID
			NetworkWithDexsAndPairs.NetworkDBId = int(NetworkDBID)

			// log.Printf("[%d/%d] [%v] Getting Dexs From Coingecko...", CountIndex, NetworkCount, Network.Attributes.Name)

			// Call API To Retrieve Network Dexs
			var DexResponse geckoterminal_types.GeckoTerminalDexs
			DexResponse = geckoterminal_api.GetGeckoterminalDexsForNetwork(CoingeckoBuildID, Network.Attributes.Identifier)

			log.Printf("[%v] Collected %d Dex(s)", Network.Attributes.Name, len(DexResponse.PageProps.Dexes))
			// logging.LogSeparator(false)

			// Sleep For N Seconds
			time.Sleep(WaitTime)

			// Iterate Through Networks Dexs
			for _, Dex := range DexResponse.PageProps.Dexes {

				// Create New Dex Object
				var CollectedDex geckoterminal_types.Dex

				// Get Values
				CollectedDex.Name = Dex.Attributes.Name
				CollectedDex.Identifier = Dex.Attributes.Identifier
				CollectedDex.ImageURL = Dex.Attributes.ImageURL
				CollectedDex.URL = Dex.Attributes.URL

				// Add It To Networks Dex List
				NetworkWithDexsAndPairs.Dexes = append(NetworkWithDexsAndPairs.Dexes, CollectedDex)

			}

			// Add Network Stablecoins
			NetworkWithDexsAndPairs.Stablecoins = mysql_query.GetNetworkStablecoinsFromDB(Network.Attributes.Identifier)

		} else {

			// Check If Network Is Already Blacklisted
			BlacklistNetworkQueryResults := mysql_query.GetBlacklistNetwork(Network.Attributes.Name)

			if len(BlacklistNetworkQueryResults) <= 0 {

				// Add Network To Blacklist Table
				mysql_insert.AddBlacklistNetworkToDB(Network.Attributes.Name, Network.Attributes.ChainID)

			}

			// log.Printf("[%d/%d] [%v] No Available RPCs - Skipping", CountIndex, NetworkCount, Network.Attributes.Name)
			// logging.LogSeparator(false)

		}

	} else {

		log.Printf("[%v] Network Blacklisted - Skipping", Network.Attributes.Name)
		// logging.LogSeparator(false)

	}

	// Send Network To Channel
	NetworkCollectionChannel <- NetworkWithDexsAndPairs
	
}
