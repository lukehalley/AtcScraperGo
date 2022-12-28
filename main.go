package main

import (
	"atcscraper/src/data"
	atcqueries "atcscraper/src/graphql/bitquery/querys"
	atchelpers "atcscraper/src/graphql/helpers"
	atctypes "atcscraper/src/types"
	"context"
	"github.com/Khan/genqlient/graphql"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {

	// Env Vars
	CacheMode, _ := strconv.ParseBool(os.Getenv("CACHE_MODE"))

	// Collection Settings
	// Networks := []string{"ethereum", "bsc", "matic", "klaytn", "avalanche", "fantom", "moonbeam"}
	Networks := []string{"klaytn"}

	// Init Vars
	var FinalCollectedDexs []atctypes.Dex

	////////////////////////////////////////////////////
	// Setup HTTP Client For API Calls
	////////////////////////////////////////////////////

	// Get The Bitquery API Key
	BitqueryApiKey := "BQYZVuAE5Tl7zOvKBmQajtsyaH5JaqmG"

	// Create Query Client
	HTTPClient := atchelpers.CreateHTTPClientWithAuth(BitqueryApiKey)
	QueryClient := graphql.NewClient("https://graphql.bitquery.io", &HTTPClient)
	QueryContext := context.Background()

	////////////////////////////////////////////////////
	// Get Bitquery Account Details
	////////////////////////////////////////////////////

	atchelpers.LogSeparator(false)

	// Get Out Current Bitquery Account Details
	AccountQuery, QueryError := atcqueries.GetBitQueryAccountDetails(QueryContext, QueryClient, BitqueryApiKey)

	// Catch Query Errors
	if QueryError != nil {
		log.Panicln("Error Querying Bitquery Account:", QueryError)
	} else {
		log.Printf("Bitquery Points Remaining: %v", AccountQuery.Utilities.ActivePeriod.PointsRemaining)
	}

	atchelpers.LogSeparator(true)

	if !CacheMode {

		////////////////////////////////////////////////////
		// Get All NetworkDexs On Each NetworkObject
		////////////////////////////////////////////////////

		// Set Our Beginning Time For Timeframe
		FromTime := atchelpers.GenerateGraphQLTimestamp("2022-01-01T01:00:00Z")

		// Set Our NetworkObject
		NumberOfDexsToCollect := 5

		// Start Log
		atchelpers.LogSeparator(false)
		log.Printf("Collecting Dex For %v Network(s)", len(Networks))
		atchelpers.LogSeparator(false)

		// Kick Off The Function Which Gets The Dexs For Each Network
		var NetworkDexs [][]atctypes.Dex
		for Index, Network := range Networks {
			IsLastElement := (len(Networks) - 1) == Index
			CurrentNetworkDexs := data.CollectDexsForNetwork(Network, NumberOfDexsToCollect, FromTime, QueryContext, QueryClient)
			NetworkDexs = append(NetworkDexs, CurrentNetworkDexs)
			if !IsLastElement {
				time.Sleep(60 * time.Second)
			}
		}

		atchelpers.LogSeparator(false)

		// Get Non-Null Dexs
		for _, NetworkDex := range NetworkDexs {
			CurrentNetwork := data.TitleCaseString(NetworkDex[0].Network)
			NetworkDexCount := 0
			for _, Dex := range NetworkDex {
				if Dex.RouterAddress != "" && Dex.FactoryAddress != "" {
					NetworkDexCount = NetworkDexCount + 1
					FinalCollectedDexs = append(FinalCollectedDexs, Dex)
				}
			}
			log.Printf("[%v] %v Dex(s)", CurrentNetwork, NetworkDexCount)
		}

	} else {

		FinalCollectedDexs = []atctypes.Dex {
			{
				Network : "klaytn",
				FactoryAddress : "0xf51082ba95c18315ca51a1ab8d9db537ff52312e",
				RouterAddress : "0x805c3fa4f0425f7b4ec6848e10de7cc7e0841a07",
			},
			{
				Network: "klaytn",
				FactoryAddress: "0xdee3df2560bceb55d3d7ef12f76dcb01785e6b29",
				RouterAddress: "0xf50782a24afcb26acb85d086cf892bfffb5731b5",
			},
			{
				Network: "klaytn",
				FactoryAddress: "0xa25ba09d8837f6319cd65b2345c0bbea99c39cb1",
				RouterAddress: "0x66ec1b0c3bf4c15a76289ac36098704acd44170f",
			},
			{
				Network: "klaytn",
				FactoryAddress: "0x3679c3766e70133ee4a7eb76031e49d3d1f2b50c",
				RouterAddress: "0xf50782a24afcb26acb85d086cf892bfffb5731b5",
			},
		}

	}

	atchelpers.LogSeparator(false)
	log.Printf("Collected %v Dex(s) Across %v Network(s)", len(FinalCollectedDexs), len(Networks))
	atchelpers.LogSeparator(true)

}
