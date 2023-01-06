package data

import (
	"atcscraper/src/db/graphql/bitquery/querys"
	"atcscraper/src/types/bitquery"
	"atcscraper/src/types/mysql"
	"atcscraper/src/types/web3"
	"atcscraper/src/util"
	"context"
	"github.com/Khan/genqlient/graphql"
	"log"
	"sort"
	"sync"
	"time"
)

func CollectDexsForNetwork(Network mysql.Network, AmountOfDexs int, FromTime time.Time, QueryContext context.Context, QueryClient graphql.Client) []bitquery.Dex {

	// Get N Dexs On Current NetworkObject Sorted By Transactions
	NetworksDexsQuery, _ := atcqueries.GetTopNDexsForNetwork(QueryContext, QueryClient, atcqueries.EthereumNetwork(Network.Name), AmountOfDexs, 0, FromTime)

	// Collect Dexs From Query
	NetworkDexs := NetworksDexsQuery.Ethereum.DexTrades
	DexCount := len(NetworkDexs)

	// Create List For Dexs
	var CollectedNetworkDexs []bitquery.Dex

	if DexCount > 0 {

		// Create Concurrency Objects
		DexContractCollectionWaitGroup := new(sync.WaitGroup)
		DexContractCollectionWaitGroup.Add(len(NetworkDexs))
		DexContractCollectionChannel := make(chan bitquery.Dex, len(NetworkDexs))

		// Kick Off The Function Which Gets The Router + Factory Contract For Each Dex
		for _, NetworkDex := range NetworkDexs {
			go CollectDexFactoryAndRouter(Network, NetworkDex, QueryContext, QueryClient, DexContractCollectionWaitGroup, DexContractCollectionChannel)
		}

		// Wait For All Groups To Come Back
		DexContractCollectionWaitGroup.Wait()

		// Close The Group Channel
		close(DexContractCollectionChannel)

		// Collect Dexs From Channel
		for DexContract := range DexContractCollectionChannel {
			CollectedNetworkDexs = append(CollectedNetworkDexs, DexContract)
		}

	}

	log.Printf("[%v] Collecting Finished", util.TitleCaseString(string(Network.Name)))

	// Return List Of Collected Dexs With Channel
	return CollectedNetworkDexs

}

func CollectDexFactoryAndRouter(Network mysql.Network, NetworkDex atcqueries.GetTopNDexsForNetworkEthereumDexTrades, QueryContext context.Context, QueryClient graphql.Client, DexContractCollectionWaitGroup *sync.WaitGroup, DexContractCollectionChannel chan bitquery.Dex) {

	// Schedule The Call To WaitGroup's Done To Tell GoRoutine Is Completed.
	defer DexContractCollectionWaitGroup.Done()

	// Create Dex Object
	var CollectedDex bitquery.Dex

	// Set Dex Network
	CollectedDex.Network = Network

	// Get The Dexs Factory Address
	DexFactoryAddress := NetworkDex.Exchange.Address.Address

	// Get 50 Pairs Created By The Factory So We Can Get The Dexs Router Address
	PairsCreatedByDexQuery, _ := atcqueries.GetPairCreatedByFactory(QueryContext, QueryClient, atcqueries.EthereumNetwork(Network.Name), 50, DexFactoryAddress)

	var TokenAddresses []string

	// Check If We Got Any Pairs Back
	if len(PairsCreatedByDexQuery.Ethereum.Arguments) > 0 {

		// Get The Primary Address Of The Pair
		PairsCreatedByDex := PairsCreatedByDexQuery.Ethereum.Arguments

		// Loop Through The Pairs Until We Get A Router Address
		for _, PairCreatedByDex := range PairsCreatedByDex {
			TokenAddresses = append(TokenAddresses, PairCreatedByDex.Reference.Address)
		}

		// Get The Router Address From One Of The Primary Tokens Dex Trades
		DexRouterQuery, QueryError := atcqueries.GetRouterAddressFromPairBaseCurrencys(QueryContext, QueryClient, atcqueries.EthereumNetwork(Network.Name), DexFactoryAddress, TokenAddresses, 1)

		// Catch Query Errors
		if QueryError == nil {

			// Check If We Got Trades Back
			if len(DexRouterQuery.Ethereum.DexTrades) > 0 {

				// Collect The Router Address
				DexRouterAddress := DexRouterQuery.Ethereum.DexTrades[0].Transaction.To.Address

				// Add Contract Addresses To Dex Object
				CollectedDex.FactoryAddress = DexFactoryAddress
				CollectedDex.RouterAddress = DexRouterAddress

			}
		}

	}

	DexContractCollectionChannel <- CollectedDex

}

func SearchPairTransactions(PairTransactions []web3.PairTransactions, BaseCurrency string, QuoteCurrency string) (bool, int) {

	Index := sort.Search(len(PairTransactions), func(i int) bool {
		return PairTransactions[i].BaseCurrency >= BaseCurrency && PairTransactions[i].QuoteCurrency >= QuoteCurrency
	})

	if Index < len(PairTransactions) && PairTransactions[Index].BaseCurrency == BaseCurrency && PairTransactions[Index].QuoteCurrency == QuoteCurrency {
		return true, Index
	} else {
		return false, -1
	}
}

func CollectStablecoinPairsForDex(Dex bitquery.Dex, Limit int, FromTime time.Time, QueryContext context.Context, QueryClient graphql.Client) ([]atcqueries.GetAllStablecoinPairsCreatedForDexEthereumDexTrades, []web3.PairTransactions) {

	// Collect Stablecoin Addresses
	var Stablecoins []string
	for _, DBStablecoin := range Dex.Network.Stablecoins {
		Stablecoins = append(Stablecoins, DBStablecoin.Address)
	}

	// List For Unique Pairs
	var UniquePairs []atcqueries.GetAllStablecoinPairsCreatedForDexEthereumDexTrades

	// Query For Dex Stablecoins
	DexStablecoinPairsQuery, _ := atcqueries.GetAllStablecoinPairsCreatedForDex(QueryContext, QueryClient, atcqueries.EthereumNetwork(Dex.Network.Name), FromTime, Limit, Dex.FactoryAddress, Stablecoins, 1)

	// Create A List Of Pair Transactions
	var PairTransactions []web3.PairTransactions

	// Collect All Unique Pairs
	CollectedPairs := DexStablecoinPairsQuery.Ethereum.DexTrades
	if len(CollectedPairs) > 0 {
		var CollectedBaseCurrencyAddresses []string
		for _, CollectedPair := range CollectedPairs {

			PairIndexCreated, ResultIndex := SearchPairTransactions(PairTransactions, CollectedPair.BaseCurrency.Address, CollectedPair.QuoteCurrency.Address)

			if !PairIndexCreated {
				var NewPair web3.PairTransactions
				NewPair.BaseCurrency = CollectedPair.BaseCurrency.Address
				NewPair.QuoteCurrency = CollectedPair.QuoteCurrency.Address
				NewPair.Transactions = append(NewPair.Transactions, CollectedPair.Transaction.Hash)
				PairTransactions = append(PairTransactions, NewPair)
			} else {
				PairTransactions[ResultIndex].Transactions = append(PairTransactions[ResultIndex].Transactions, CollectedPair.Transaction.Hash)
			}

			target := CollectedPair.BaseCurrency.Address
			sort.Strings(CollectedBaseCurrencyAddresses)
			i := sort.SearchStrings(CollectedBaseCurrencyAddresses, target)
			TokenAlreadyCollected := i < len(CollectedBaseCurrencyAddresses) && CollectedBaseCurrencyAddresses[i] == target
			if !TokenAlreadyCollected {
				CollectedBaseCurrencyAddresses = append(CollectedBaseCurrencyAddresses, CollectedPair.BaseCurrency.Address)
				UniquePairs = append(UniquePairs, CollectedPair)
			}
		}
	}

	return UniquePairs, PairTransactions
}