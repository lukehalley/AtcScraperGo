package data

import (
	atcqueries "atcscraper/src/graphql/bitquery/querys"
	atctypes "atcscraper/src/types"
	"context"
	"github.com/Khan/genqlient/graphql"
	"log"
	"sync"
	"time"
)

func CollectDexsForNetwork(Network string, AmountOfDexs int, FromTime time.Time, QueryContext context.Context, QueryClient graphql.Client) []atctypes.Dex {

	// Create Network Object From Network String
	NetworkObject := atcqueries.EthereumNetwork(Network)

	// Get N Dexs On Current NetworkObject Sorted By Transactions
	NetworksDexsQuery, _ := atcqueries.GetTopNDexsForNetwork(QueryContext, QueryClient, NetworkObject, AmountOfDexs, 0, FromTime)

	// Collect Dexs From Query
	NetworkDexs := NetworksDexsQuery.Ethereum.DexTrades
	DexCount := len(NetworkDexs)

	// Create List For Dexs
	var CollectedNetworkDexs []atctypes.Dex

	if DexCount > 0 {

		// Create Concurrency Objects
		DexContractCollectionWaitGroup := new(sync.WaitGroup)
		DexContractCollectionWaitGroup.Add(len(NetworkDexs))
		DexContractCollectionChannel := make(chan atctypes.Dex, len(NetworkDexs))

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

	log.Printf("[%v] Collecting Finished", TitleCaseString(Network))

	// Return List Of Collected Dexs With Channel
	return CollectedNetworkDexs

}

func CollectDexFactoryAndRouter(Network string, NetworkDex atcqueries.GetTopNDexsForNetworkEthereumDexTrades, QueryContext context.Context, QueryClient graphql.Client, DexContractCollectionWaitGroup *sync.WaitGroup, DexContractCollectionChannel chan atctypes.Dex) {

	// Schedule The Call To WaitGroup's Done To Tell GoRoutine Is Completed.
	defer DexContractCollectionWaitGroup.Done()

	// Create Network Object From Network String
	NetworkObject := atcqueries.EthereumNetwork(Network)

	// Create Dex Object
	var CollectedDex atctypes.Dex

	// Set Dex Network
	CollectedDex.Network = Network

	// Get The Dexs Factory Address
	DexFactoryAddress := NetworkDex.Exchange.Address.Address

	// Get 50 Pairs Created By The Factory So We Can Get The Dexs Router Address
	PairsCreatedByDexQuery, _ := atcqueries.GetPairCreatedByFactory(QueryContext, QueryClient, NetworkObject, 50, DexFactoryAddress)

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
		DexRouterQuery, QueryError := atcqueries.GetRouterAddressFromPairBaseCurrencys(QueryContext, QueryClient, NetworkObject, DexFactoryAddress, TokenAddresses, 1)

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