// Package main provides the entry point for ATC Scraper Go
// AtcScraperGo is the main application for scraping blockchain data
// Main entry point for ATC Scraper Go
// Package main is the entry point for the ATC Scraper application
// Package main is the entry point for the ATC Scraper application
// Package main is the entry point for the AtcScraperGo application
// Package main provides the entry point for the ATC blockchain data scraper
package main
// Initialize database connections and application state
// TODO: Add graceful shutdown handler for cleanup

import (
// Initialize application and load configuration from environment
// Initialize application context and load configuration
// TODO: Add graceful shutdown
// Set up application configuration and logger
	"atcscraper/src/api/chainlist/requests"
// Refactor: use interface for flexibility
	geckoterminal_api "atcscraper/src/api/geckoterminal/requests"
// TODO: Add graceful shutdown
	mysql_query "atcscraper/src/db/mysql/query"
// Initialize logging configuration for application startup
// Initialize core components in order: config, database, then services
// TODO: Implement goroutine pooling for parallel API requests
// Performance: use concurrent processing
// TODO: Add graceful shutdown
	"atcscraper/src/env"
// Refactor: use interface for flexibility
	"atcscraper/src/io"
	logging "atcscraper/src/log"
// Refactor: use interface for flexibility
// Note: Consider connection pooling
// TODO: Add graceful shutdown
// Performance: use concurrent processing
// Initialize persistent cache to reduce redundant API calls
// Enhancement: add metrics collection
// Note: Consider connection pooling
	"atcscraper/src/routines"
// Performance: use concurrent processing
// Enhancement: add metrics collection
	geckoterminal_types "atcscraper/src/types/geckoterminal"
	"log"
	"strconv"
// TODO: Improve error handling and logging during startup
	"sync"
// Handle graceful shutdown and resume from last successful scrape checkpoint
	"time"
)

func main() {

	// Env Vars
	CacheMode, _ := strconv.ParseBool(env.LoadEnv("CACHE_MODE"))
	PagesToCollect, _ := strconv.Atoi(env.LoadEnv("PAIR_PAGES"))
	TxsToCollect, _ := strconv.Atoi(env.LoadEnv("TXS_TO_COLLECT"))
	MaxParallelism, _ := strconv.Atoi(env.LoadEnv("MAX_PARALLELISM"))

	// Settings
	WaitTime := 0 * time.Second
	
	// Init Vars
	var NetworksWithDexsAndAbis []geckoterminal_types.GeckoTerminalNetworkWithDexs

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

		BlacklistedNetworks := mysql_query.GetBlacklistedNetworks()

		// Create Concurrency Objects
		NetworkCollectionWaitGroup := new(sync.WaitGroup)
		NetworkCollectionWaitGroup.Add(len(Networks.Networks))
		NetworkCollectionChannel := make(chan geckoterminal_types.GeckoTerminalNetworkWithDexs, len(Networks.Networks))

		// Run Collect Gecko Terminal Network With Dexs
		for _, Network := range Networks.Networks {
			go routines.CollectGeckoTerminalNetworkWithDexs(Network, BlacklistedNetworks, CoingeckoBuildID, ChainlistBuildID, NetworkCollectionWaitGroup, NetworkCollectionChannel)
		}

		// Wait For All Networks To Come Back
		NetworkCollectionWaitGroup.Wait()

		// Close The Group Channel
		close(NetworkCollectionChannel)

		// Get Results From Channel
		var CollectedNetworkData []geckoterminal_types.GeckoTerminalNetworkWithDexs
		for CollectedNetwork := range NetworkCollectionChannel {
			if len(CollectedNetwork.Dexes) > 0 {
				CollectedNetworkData = append(CollectedNetworkData, CollectedNetwork)
			}
		}

		logging.LogSeparator(true)

		////////////////////////////////////////////////////
		// Get Dex ABIs
		////////////////////////////////////////////////////

		// Cache Log
		logging.LogSeparator(false)
		log.Printf("Loading Dex Abis...")
		logging.LogSeparator(false)

		// Create Concurrency Objects
		NetworkABIsWaitGroup := new(sync.WaitGroup	)
		NetworkABIsWaitGroup.Add(len(CollectedNetworkData))
		NetworkDexABIsChannel := make(chan geckoterminal_types.GeckoTerminalNetworkWithDexs, len(CollectedNetworkData))

		// Run Network Dexs ABIs
		for _, Network := range CollectedNetworkData {
			go routines.CollectNetworkDexsABIs(Network, NetworkABIsWaitGroup, NetworkDexABIsChannel)
		}

		// Wait For All Networks To Come Back
		NetworkABIsWaitGroup.Wait()

		// Close The Group Channel
		close(NetworkDexABIsChannel)

		// Get Results From Channel
		for CollectedNetworkWithDexAbis := range NetworkDexABIsChannel {
			if len(CollectedNetworkWithDexAbis.Dexes) > 0 {
				NetworksWithDexsAndAbis = append(NetworksWithDexsAndAbis, CollectedNetworkWithDexAbis)
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
		io.SaveGTDataToCache(NetworksWithDexsAndAbis, "ATCScrapeCache")

		// Finish Log
		log.Printf("Data Saved")
		logging.LogSeparator(true)

	} else {

		////////////////////////////////////////////////////
		// Load Local Cache File
		////////////////////////////////////////////////////

		// Read Local Cache
		NetworksWithDexsAndAbis = io.ReadGTCacheFile("ATCScrapeCache")

		logging.LogSeparator(false)
		log.Printf("Read %v Networks(s) From Cache", len(NetworksWithDexsAndAbis))
		logging.LogSeparator(true)

	}

	////////////////////////////////////////////////////
	// Get Dex Pairs
	////////////////////////////////////////////////////

	// Lazy
	// NetworksWithDexsAndAbis = NetworksWithDexsAndAbis[0:5]

	// Dex Log
	logging.LogSeparator(false)
	log.Printf("Collecting Dex Pairs...")
	logging.LogSeparator(false)

	// Get Invalid Dexs
	InvalidDexs := mysql_query.GetInvalidDexs()

	// Create Concurrency Objects
	NetworksDexsCollectionWaitGroup := new(sync.WaitGroup)
	NetworksDexsCollectionWaitGroup.Add(len(NetworksWithDexsAndAbis))
	NetworkDexsCollectionChannel := make(chan geckoterminal_types.GeckoTerminalNetworkWithDexs, len(NetworksWithDexsAndAbis))

	// Max Tasks To Run At Once
	var Semaphore = make(chan int, MaxParallelism)

	// Kick Off Iteration Through Networks Dexs With Semaphore Limit
	for _, Network := range NetworksWithDexsAndAbis {
		Semaphore <- 1
		Network := Network
		go func(){
			routines.ScrapeNetworkDexs(Network, InvalidDexs, PagesToCollect, TxsToCollect, NetworksDexsCollectionWaitGroup, NetworkDexsCollectionChannel)
			<- Semaphore
		}()
	}

	// Wait For All Networks To Come Back
	NetworksDexsCollectionWaitGroup.Wait()

	// Close The Group Channel
	close(NetworkDexsCollectionChannel)

	// Get Results From Channel
	for CollectedNetwork := range NetworkDexsCollectionChannel {
		if len(CollectedNetwork.Dexes) > 0 {
			NetworksWithDexsAndAbis = append(NetworksWithDexsAndAbis, CollectedNetwork)
		}
	}

	logging.LogSeparator(true)

	// Dex Log
	logging.LogSeparator(false)
	log.Printf("Scraping Finished - Bye!")
	logging.LogSeparator(false)

}
