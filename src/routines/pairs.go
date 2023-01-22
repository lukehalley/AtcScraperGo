package routines

import (
	geckoterminal_api "atcscraper/src/api/geckoterminal/requests"
	mysql_insert "atcscraper/src/db/mysql/insert"
	mysql_query "atcscraper/src/db/mysql/query"
	geckoterminal_types "atcscraper/src/types/geckoterminal"
	"sync"
)

func CollectGeckoTerminalDexPairs(Network geckoterminal_types.GeckoTerminalNetworkWithDexs, NetworkStablecoins []string, Dex geckoterminal_types.Dex, InvalidDexs []string, BlacklistPairAddresses []string, PagesToCollect int, TxsToCollect int, DexCollectionWaitGroup *sync.WaitGroup, DexCollectionChannel chan geckoterminal_types.Dex) {

	// Schedule The Call To WaitGroup's Done To Tell GoRoutine Is Completed.
	defer DexCollectionWaitGroup.Done()

	// Dex Is Invalid Dex List
	// DexIsInvalid, _ := util.CheckIfStringIsInList(InvalidDexs, Dex.Identifier, false)
	//if !DexIsInvalid {
	//
	//}

	// Check If Dex Is Already Stored
	DexQueryResults := mysql_query.GetDexFromDB(Network.NetworkDBId, Dex.Identifier)

	// Get Dex DB ID
	DexDBId := 0
	if len(DexQueryResults) > 0 {

		// Set Dex DB ID
		DexDBId = DexQueryResults[0].DexId

	} else {

		// Add Dex To DB
		DexDBId = int(mysql_insert.AddDexToDB(
			Network.NetworkDBId,
			Dex.Identifier,
			Dex.RouterAddress,
			Dex.FactoryAddress,
			1,
		))

	}

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

	// Create Concurrency Objects
	PairCollectionWaitGroup := new(sync.WaitGroup)
	PairCollectionWaitGroup.Add(len(DexPairs.Data))
	PairCollectionChannel := make(chan geckoterminal_types.Pair, len(DexPairs.Data))

	// Iterate Through Networks Dexs And Get All Their Pairs
	for _, DexPair := range DexPairs.Data {
		go ScrapePairInfo(Network, NetworkStablecoins, Dex, DexDBId, DexPair, BlacklistPairAddresses, TxsToCollect, PairCollectionWaitGroup, PairCollectionChannel)
	}

	// Wait For All Networks To Come Back
	PairCollectionWaitGroup.Wait()

	// Close The Group Channel
	close(PairCollectionChannel)

	DexCollectionChannel <- Dex
	
}
