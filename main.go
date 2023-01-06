package main

import (
	"atcscraper/src/api/chainlist/requests"
	geckoterminal_api "atcscraper/src/api/geckoterminal/requests"
	mysql_insert "atcscraper/src/db/mysql/insert"
	mysql_query "atcscraper/src/db/mysql/query"
	"atcscraper/src/io"
	logging "atcscraper/src/log"
	geckoterminal_types "atcscraper/src/types/geckoterminal"
	"atcscraper/src/util"
	"atcscraper/src/web3"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// Settings
	WaitTime := 0 * time.Second
	
	// Init Vars
	var CollectedData []geckoterminal_types.GeckoTerminalNetworkWithDexs

	// Env Vars
	// LazyMode := false
	CacheMode, _ := strconv.ParseBool(os.Getenv("CACHE_MODE"))

	logging.LogSeparator(false)
	log.Printf("Collecting Coingecko CoingeckoBuildID")
	logging.LogSeparator(false)

	CoingeckoBuildID := geckoterminal_api.GetGeckoterminalBuildID()
	ChainlistBuildID := chainlist.GetChainlistBuildID()

	log.Printf("Coingecko Build ID: %v", CoingeckoBuildID)
	log.Printf("ChainList Build ID: %v", ChainlistBuildID)
	logging.LogSeparator(true)

	// Sleep For N Seconds
	time.Sleep(WaitTime)

	if !CacheMode {

		////////////////////////////////////////////////////
		// Call GT To Get Networks
		////////////////////////////////////////////////////

		logging.LogSeparator(false)
		log.Printf("Collecting Coingecko Networks")
		logging.LogSeparator(false)

		// Get All Networks On Geckoterminal
		Networks := geckoterminal_api.GetGeckoterminalNetworks(CoingeckoBuildID)

		// Network Count
		NetworkCount := len(Networks.Networks)

		log.Printf("Collected %d Networks", NetworkCount)
		logging.LogSeparator(true)

		// Sleep For N Seconds
		time.Sleep(WaitTime)

		////////////////////////////////////////////////////
		// Iterate Through Networks And Get All Their Dexs
		////////////////////////////////////////////////////

		logging.LogSeparator(false)
		log.Printf("Collecting Network Dexs")
		logging.LogSeparator(false)

		for Index, Network := range Networks.Networks {

			CountIndex := Index + 1

			// Create Collection Object
			var NetworkWithDexsAndPairs geckoterminal_types.GeckoTerminalNetworkWithDexs

			log.Printf("[%d/%d] [%v] Getting RPCs From Chainlist...", CountIndex, NetworkCount, Network.Attributes.Name)

			// Get Chain RPCs
			ChainInfo := chainlist.GetChainInfo(Network.Attributes.ChainID, ChainlistBuildID)

			// Sleep For N Seconds
			time.Sleep(WaitTime)

			var FilteredRPCs []string
			for _, RPC := range ChainInfo.PageProps.Chain.RPC {
				if !(strings.Contains(RPC.URL, "API")) {
					FilteredRPCs = append(FilteredRPCs, RPC.URL)
				}
			}

			// Check If We Have Any RPCs
			if len(FilteredRPCs) > 0 {

				log.Printf("[%d/%d] [%v] Collected %d RPC(s)", CountIndex, NetworkCount, Network.Attributes.Name, len(FilteredRPCs))

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

				log.Printf("[%d/%d] [%v] Getting Dexs From Coingecko...", CountIndex, NetworkCount, Network.Attributes.Name)

				// Call API To Retrieve Network Dexs
				var DexResponse geckoterminal_types.GeckoTerminalDexs
				DexResponse = geckoterminal_api.GetGeckoterminalDexsForNetwork(CoingeckoBuildID, Network.Attributes.Identifier)

				log.Printf("[%d/%d] [%v] Collected %d Dex(s)", CountIndex, NetworkCount, Network.Attributes.Name, len(DexResponse.PageProps.Dexes))
				logging.LogSeparator(false)

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

					// Check If Network Is Already Stored
					DexQueryResults := mysql_query.GetDexFromDB(NetworkWithDexsAndPairs.NetworkDBId, CollectedDex.Identifier)

					var DexDBId int64
					if len(DexQueryResults) > 0 {
						// Get Network DB ID
						DexDBId = int64(DexQueryResults[0].DexId)
					} else {
						// Add Network To DB
						DexDBId = mysql_insert.AddDexToDB(
							NetworkWithDexsAndPairs.NetworkDBId,
							CollectedDex.Identifier,
						)
					}

					// Set Network DB ID
					NetworkWithDexsAndPairs.DexDBId = int(DexDBId)

				}

				// Add Network Stablecoins
				NetworkWithDexsAndPairs.Stablecoins = mysql_query.GetNetworkStablecoinsFromDB(NetworkDBID)

				// Add Network With Data To Master List
				CollectedData = append(CollectedData, NetworkWithDexsAndPairs)

			} else {

				log.Printf("[%d/%d] [%v] No Available RPCs - Skipping", CountIndex, NetworkCount, Network.Attributes.Name)
				logging.LogSeparator(false)

			}

		}

		logging.LogSeparator(true)

		////////////////////////////////////////////////////
		// Save Data To Local Cache
		////////////////////////////////////////////////////

		// Cache Log
		logging.LogSeparator(false)
		log.Printf("Saving Data To Local Cache...")
		logging.LogSeparator(false)

		// Save Data To Local File
		io.SaveGTDataToCache(CollectedData, "GTCollectedData")

		// Finish Log
		log.Printf("Data Saved")
		logging.LogSeparator(true)

	} else {

		////////////////////////////////////////////////////
		// Load Local Cache File
		////////////////////////////////////////////////////

		// Read Local Cache
		CollectedData = io.ReadGTCacheFile("GTCollectedData")

		logging.LogSeparator(false)
		log.Printf("Read %v Networks(s) From Cache", len(CollectedData))
		logging.LogSeparator(true)

	}

	////////////////////////////////////////////////////
	// Get Dex Pairs
	////////////////////////////////////////////////////

	// Iterate Through Networks Dexs
	for _, Network := range CollectedData {

		// Collect Network Stablecoin Addresses
		var NetworkStablecoins []string
		for _, Stablecoin := range Network.Stablecoins {
			NetworkStablecoins = append(NetworkStablecoins, Stablecoin.Address)
		}

		// Get Pairs For Each Dex
		for _, Dex := range Network.Dexes {

			// Get All Pairs For Current Dex
			DexPairs := geckoterminal_api.GetGeckoterminalDexPairs(Network.Network.Identifier, Dex.Identifier, 1)

			// Get Factory Address
			if Dex.FactoryAddress == "" {
				Dex.FactoryAddress = web3.GetPairFactoryAddress(DexPairs.Data[0].Attributes.Address, Network.RPCs[0])
			}

			// Sleep For N Seconds
			time.Sleep(WaitTime)

			// Iterate Through Pairs We Got
			for _, DexPair := range DexPairs.Data {

				// Create New Pair Object
				var Pair geckoterminal_types.Pair

				// Add The Pair Address
				Pair.Address = DexPair.Attributes.Address

				// Get Current Pair Transactions
				PairTransactions := geckoterminal_api.GetGeckoterminalPairTransactionsAndTokens(Network.Network.Identifier, Pair.Address, 1)

				// Sleep For N Seconds
				time.Sleep(WaitTime)

				// Collect Transactions
				for _, PairTransaction := range PairTransactions.Data {
					Pair.Transactions = append(Pair.Transactions, PairTransaction.Attributes.TxHash)
				}

				if len(Pair.Transactions) > 0 {

					// Add Base Token Details
					var BaseToken geckoterminal_types.Token
					BaseToken.Name = PairTransactions.Included[1].Attributes.Name
					BaseToken.Symbol = PairTransactions.Included[1].Attributes.Symbol
					BaseToken.Address = PairTransactions.Included[1].Attributes.Address
					Pair.BaseToken = BaseToken

					// Add Quote Token Details
					var QuoteToken geckoterminal_types.Token
					QuoteToken.Name = PairTransactions.Included[0].Attributes.Name
					QuoteToken.Symbol = PairTransactions.Included[0].Attributes.Symbol
					QuoteToken.Address = PairTransactions.Included[0].Attributes.Address
					Pair.QuoteToken = QuoteToken

					// Add The Pair Address
					Pair.Name = fmt.Sprintf("%v/%v", BaseToken.Symbol, QuoteToken.Symbol)

					// Check If Pair Contains A Stablecoin
					IsStablecoinPair := (util.CheckIfStringIsInList(NetworkStablecoins, BaseToken.Address)) || (util.CheckIfStringIsInList(NetworkStablecoins, QuoteToken.Address))

					// Append The Pair To Out Collection If Pair Contains A Stablecoin
					if IsStablecoinPair {
						Dex.Pairs = append(Dex.Pairs, Pair)
					}

				}

			}

			log.Printf("[%v - %v] %v Pairs(s)", Network.Network.Name, Dex.Name, len(Dex.Pairs))

		}

		// Get Pair Transactions
		geckoterminal_api.GetGeckoterminalPairTransactionsAndTokens(Network.Network.Identifier, "Yo", 1)
		
	}

}
