package routines

import (
	geckoterminal_api "atcscraper/src/api/geckoterminal/requests"
	mysql_insert "atcscraper/src/db/mysql/insert"
	mysql_query "atcscraper/src/db/mysql/query"
	geckoterminal_types "atcscraper/src/types/geckoterminal"
	"atcscraper/src/util"
	"atcscraper/src/web3"
	"fmt"
	"log"
	"sync"
)

func ScrapePairInfo(Network geckoterminal_types.GeckoTerminalNetworkWithDexs, NetworkStablecoins []string, Dex geckoterminal_types.Dex, DexPair geckoterminal_types.PairData, BlacklistPairAddresses []string, TxsToCollect int) {

	// Create New Pair Object
	var Pair geckoterminal_types.Pair

	// Add The Pair Address
	Pair.Address = DexPair.Attributes.Address

	// Get Current Pair Transactions
	PairTransactions := geckoterminal_api.GetGeckoterminalPairTransactionsAndTokens(Network.Network.Identifier, Pair.Address, 1)

	// Collect Transactions
	for _, Transaction := range PairTransactions.Data {
		var PairTransaction geckoterminal_types.Transaction
		PairTransaction.Hash = Transaction.Attributes.TxHash
		Pair.Transactions = append(Pair.Transactions, PairTransaction)
	}

	// Check If We Got Transactions
	if len(Pair.Transactions) > 0 {

		// Get Only Txs We Want
		if len(Pair.Transactions) > TxsToCollect {
			Pair.Transactions = Pair.Transactions[0:TxsToCollect]
		}

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

		// Add Market Metadata
		Pair.TwentyFourHourVolume = DexPair.Attributes.ToVolumeInUsd
		Pair.TwentyFourHourTxs = DexPair.Attributes.SwapCount24H
		Pair.Liquidity = DexPair.Attributes.ReserveInUsd

		// Add The Pair Address
		Pair.Name = fmt.Sprintf("%v/%v", BaseToken.Symbol, QuoteToken.Symbol)

		// Check If Pair Contains A Stablecoin
		BaseTokenIsStablecoin, BaseTokenMatchIndex := util.CheckIfStringIsInList(NetworkStablecoins, BaseToken.Address, false)
		QuoteTokenIsStablecoin, QuoteTokenMatchIndex := util.CheckIfStringIsInList(NetworkStablecoins, QuoteToken.Address, false)

		// Check If Both Tokens Are Stablecoins
		BothTokensAreStablecoins := BaseTokenIsStablecoin && QuoteTokenIsStablecoin

		// Check If Pair Contains A Stablecoin
		if (BaseTokenIsStablecoin || QuoteTokenIsStablecoin) && !BothTokensAreStablecoins {

			// Create Concurrency Objects
			TxDecodeWaitGroup := new(sync.WaitGroup)
			TxDecodeWaitGroup.Add(len(Pair.Transactions))
			TxDecodeChannel := make(chan geckoterminal_types.Transaction, len(Pair.Transactions))

			// Collect Transactions
			for _, PairTransaction := range Pair.Transactions {
				go DecodeTransaction(Network, PairTransaction, Dex.RouterAddress, Dex.RouterAbiDBId, TxDecodeWaitGroup, TxDecodeChannel)
			}

			// Wait For All Txs To Come Back
			TxDecodeWaitGroup.Wait()

			// Close The Group Channel
			close(TxDecodeChannel)

			// Get Results From Channel
			var CollectedDecodedTxs []geckoterminal_types.Transaction
			for DecodedTx := range TxDecodeChannel {
				if DecodedTx.MethodName != "" && len(DecodedTx.InputData.Path) > 0 {
					CollectedDecodedTxs = append(CollectedDecodedTxs, DecodedTx)
				}
			}

			// Check If We Have Any Decoded Transactions
			if len(CollectedDecodedTxs) > 0 {

				// Append The Pair To Out Collection
				Dex.Pairs = append(Dex.Pairs, Pair)

				// DB Ids
				var TokenDBId int64
				var PairDBId int64

				if BaseTokenIsStablecoin {

					// Check If Our Token Is Already In DB
					TokenQueryResults := mysql_query.GetTokenFromDB(Network.NetworkDBId, QuoteToken.Address)

					if len(TokenQueryResults) > 0 {

						// Set Token DB ID
						TokenDBId = int64(TokenQueryResults[0].TokenId)

					} else {

						// Get Token Decimals
						QuoteToken.Decimals = int(web3.GetTokenDecimals(QuoteToken.Address, Network.RPCs[0]))

						// Add Token To DB
						TokenDBId = mysql_insert.AddTokenToDB(QuoteToken.Symbol, QuoteToken.Address, QuoteToken.Decimals, Network.NetworkDBId)
					}

					// Get Out Stablecoin Details
					StablecoinDetails := Network.Stablecoins[BaseTokenMatchIndex]

					// Check If Our Pair Is Already In DB
					PairQueryResults := mysql_query.GetPairFromDB(Pair.Address, Network.NetworkDBId)

					if len(PairQueryResults) > 0 {

						// Set Pair DB ID
						PairDBId = int64(PairQueryResults[0].PairId)

					} else {

						if TokenDBId < 1 {

							// Check If Our Token Is Already In DB
							TokenQueryResults = mysql_query.GetTokenFromDB(Network.NetworkDBId, QuoteToken.Address)

							if len(TokenQueryResults) > 0 {

								// Set Token DB ID
								TokenDBId = int64(TokenQueryResults[0].TokenId)

							} else {

								log.Panicln("Invalid Token DB Id: ", TokenDBId)

							}

						}

						// Add Pair To DB
						PairDBId = mysql_insert.AddPairToDB(TokenDBId, StablecoinDetails.StablecoinId, Network.NetworkDBId, Network.DexDBId, Pair.Name, Pair.Address)

					}

				} else {

					// Check If Our Token Is Already In DB
					TokenQueryResults := mysql_query.GetTokenFromDB(Network.NetworkDBId, BaseToken.Address)
					if len(TokenQueryResults) > 0 {

						// Set Token DB ID
						TokenDBId = int64(TokenQueryResults[0].TokenId)

					} else {

						// Get Token Decimals
						BaseToken.Decimals = int(web3.GetTokenDecimals(BaseToken.Address, Network.RPCs[0]))

						// Add Token To DB
						TokenDBId = mysql_insert.AddTokenToDB(BaseToken.Symbol, BaseToken.Address, BaseToken.Decimals, Network.NetworkDBId)
					}

					// Get Out Stablecoin Details
					StablecoinDetails := Network.Stablecoins[QuoteTokenMatchIndex]

					// Check If Our Pair Is Already In DB
					PairQueryResults := mysql_query.GetPairFromDB(Pair.Address, Network.NetworkDBId)
					if len(PairQueryResults) > 0 {

						// Log Addition
						log.Printf("[%v] [%v] Present: %v", Network.Network.Name, Dex.Name, Pair.Name)

						// Set Pair DB ID
						PairDBId = int64(PairQueryResults[0].PairId)

					} else {

						if TokenDBId < 1 {
							log.Panicln("Invalid Token DB Id: ", TokenDBId)
						}

						// Add Pair To DB
						PairDBId = mysql_insert.AddPairToDB(TokenDBId, StablecoinDetails.StablecoinId, Network.NetworkDBId, Network.DexDBId, Pair.Name, Pair.Address)

						// Log Addition
						log.Printf("[%v] [%v] Added: %v", Network.Network.Name, Dex.Name, Pair.Name)

					}

				}

				// Add Pair Routes To DB
				for _, PairTransaction := range CollectedDecodedTxs {

					// Create A Comma Seperated String For Route
					var RouteString string
					for RouteAddressIndex, RouteAddress := range PairTransaction.InputData.Path {
						if RouteAddressIndex <= 0 {
							RouteString = fmt.Sprintf("%v", RouteAddress)
						} else {
							RouteString = fmt.Sprintf("%v,%v", RouteString, RouteAddress)
						}

					}

					// Check If Our Route Is Stored
					RouteQueryResults := mysql_query.GetRouteFromDB(Network.NetworkDBId, Network.DexDBId, int(PairDBId), PairTransaction.Hash)
					if len(RouteQueryResults) <= 0 {

						// Add Route To DB
						mysql_insert.AddRouteToDB(Network.NetworkDBId, Network.DexDBId, PairDBId, RouteString, PairTransaction.MethodName, PairTransaction.Hash, PairTransaction.InputData.AmountIn, PairTransaction.InputData.AmountOutMin)

					}

				}

			}

		} else {

			// Check If Pair Is Blacklisted
			BlacklistPairQueryResults := mysql_query.GetBlacklistPairFromDB(Pair.Address, Network.NetworkDBId)
			if len(BlacklistPairQueryResults) <= 0 {

				// Add Pair To Pair Blacklist DB
				// mysql_insert.AddBlacklistPairToDB(Network.NetworkDBId, Pair.Name, Pair.Address)

			}

		}

	}
	
}