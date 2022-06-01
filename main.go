package main

import (
	"atcscraper/src/api/chainlist/requests"
	geckoterminal_api "atcscraper/src/api/geckoterminal/requests"
	mysql_insert "atcscraper/src/db/mysql/insert"
	mysql_query "atcscraper/src/db/mysql/query"
	mysql_utils "atcscraper/src/db/mysql/utils"
	"atcscraper/src/env"
	"atcscraper/src/io"
	logging "atcscraper/src/log"
	"atcscraper/src/requests"
	geckoterminal_types "atcscraper/src/types/geckoterminal"
	"atcscraper/src/util"
	"atcscraper/src/web3"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {

	// Create Connection To DB
	mysql_utils.CreateDatabaseConnection()

	// Env Vars
	CacheMode, _ := strconv.ParseBool(env.LoadEnv("CACHE_MODE"))
	CollectionFailLimit, _ := strconv.Atoi(env.LoadEnv("COLLECTION_FAIL_LIMIT"))
	PairPages, _ := strconv.Atoi(env.LoadEnv("PAIR_PAGES"))

	// Settings
	WaitTime := 0 * time.Second
	
	// Init Vars
	var CollectedNetworkData []geckoterminal_types.GeckoTerminalNetworkWithDexs

	// ABIs
	RouterAbi := io.LoadAbiAsString("IUniswapV2Router02.json")

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

		for Index, Network := range Networks.Networks {

			// Counter
			CountIndex := Index + 1

			// Create Collection Object
			var NetworkWithDexsAndPairs geckoterminal_types.GeckoTerminalNetworkWithDexs

			// Check If Network Blacklisted
			NetworkIsBlacklisted, _ := util.CheckIfIntIsInList(BlacklistedNetworks, Network.Attributes.ChainID)

			if !NetworkIsBlacklisted {

				log.Printf("[%d/%d] [%v] Getting RPCs From Chainlist...", CountIndex, NetworkCount, Network.Attributes.Name)

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

					}

					// Add Network Stablecoins
					NetworkWithDexsAndPairs.Stablecoins = mysql_query.GetNetworkStablecoinsFromDB(Network.Attributes.Identifier)

					// Add Network With Data To Master List
					CollectedNetworkData = append(CollectedNetworkData, NetworkWithDexsAndPairs)

				} else {

					// Check If Network Is Already Blacklisted
					BlacklistNetworkQueryResults := mysql_query.GetBlacklistNetwork(Network.Attributes.Name)

					if len(BlacklistNetworkQueryResults) <= 0 {

						// Add Network To Blacklist Table
						mysql_insert.AddBlacklistNetworkToDB(Network.Attributes.Name, Network.Attributes.ChainID)

					}

					log.Printf("[%d/%d] [%v] No Available RPCs - Skipping", CountIndex, NetworkCount, Network.Attributes.Name)
					logging.LogSeparator(false)

				}

			} else {

				log.Printf("[%d/%d] [%v] Network Blacklisted - Skipping", CountIndex, NetworkCount, Network.Attributes.Name)
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
		io.SaveGTDataToCache(CollectedNetworkData, "GTCollectedData")

		// Finish Log
		log.Printf("Data Saved")
		logging.LogSeparator(true)

	} else {

		////////////////////////////////////////////////////
		// Load Local Cache File
		////////////////////////////////////////////////////

		// Read Local Cache
		CollectedNetworkData = io.ReadGTCacheFile("GTCollectedData")

		logging.LogSeparator(false)
		log.Printf("Read %v Networks(s) From Cache", len(CollectedNetworkData))
		logging.LogSeparator(true)

	}

	////////////////////////////////////////////////////
	// Get Dex Pairs
	////////////////////////////////////////////////////

	// Dex Log
	logging.LogSeparator(false)
	log.Printf("Collecting Dex Pairs...")
	logging.LogSeparator(false)

	// Get Invalid Dexs
	InvalidDexs := mysql_query.GetInvalidDexs()
	
	// Network Count
	NetworkCount := len(CollectedNetworkData)

	// Iterate Through Networks Dexs
	for NetworkIndex, Network := range CollectedNetworkData {

		// Get All Network Blacklist Pairs So We Can Skip Them
		BlacklistPairAddresses := mysql_query.GetBlacklistedPairsForNetwork(Network.NetworkDBId)

		// Count For Network Iteration
		NetworkCountIndex := NetworkIndex + 1

		// Collect Network Stablecoin Addresses
		var NetworkStablecoins []string

		for _, Stablecoin := range Network.Stablecoins {
			NetworkStablecoins = append(NetworkStablecoins, Stablecoin.Address)
		}

		// Dex Count
		DexCount := len(Network.Dexes)

		log.Printf("[%d/%d] [%v] Collecting Pairs Across %d Dex(s)", NetworkCountIndex, NetworkCount, Network.Network.Name, DexCount)

		// Get Pairs For Each Dex
		for DexIndex, Dex := range Network.Dexes {

			// Count For Dex Iteration
			DexCountIndex := DexIndex + 1

			log.Printf("  [%d/%d] Dex: %v", DexCountIndex, DexCount, Dex.Name)

			// Dex Is Invalid Dex List
			DexIsInvalid, _ := util.CheckIfStringIsInList(InvalidDexs, Dex.Identifier, false)

			if !DexIsInvalid {

				// Dex DB Vars
				DexDBId := -1

				// State Vars
				DexFailCount := 0
				DexIsValid := true

				// Get First Page Dexs
				DexPairs := geckoterminal_api.GetGeckoterminalDexPairs(Network.Network.Identifier, Dex.Identifier, 1)

				// If We Want To Collect More Than One Page
				if PairPages > 1 {

					for i := 2; i <= PairPages; i++ {

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

				// Sleep For N Seconds
				time.Sleep(WaitTime)

				// Pair Count
				PairCount := len(DexPairs.Data)

				// Iterate Through Pairs We Got
				for PairIndex, DexPair := range DexPairs.Data {

					// Count For Pair Iteration
					PairCountIndex := PairIndex + 1

					// Create New Pair Object
					var Pair geckoterminal_types.Pair

					// Add The Pair Address
					Pair.Address = DexPair.Attributes.Address

					// Dex Is Invalid Dex List
					PairIsBlacklisted, _ := util.CheckIfStringIsInList(BlacklistPairAddresses, Pair.Address, false)

					// If The Pair Is Not Blacklisted Get Info On It
					if !PairIsBlacklisted {

						// Get Factory Address
						if Dex.FactoryAddress == "" {
							Dex.FactoryAddress = web3.GetPairFactoryAddress(DexPair.Attributes.Address, Network.RPCs[0])
						}

						// Get Current Pair Transactions
						PairTransactions := geckoterminal_api.GetGeckoterminalPairTransactionsAndTokens(Network.Network.Identifier, Pair.Address, 1)

						// Sleep For N Seconds
						time.Sleep(WaitTime)

						// Collect Transactions
						for _, Transaction := range PairTransactions.Data {
							var PairTransaction geckoterminal_types.Transaction
							PairTransaction.Hash = Transaction.Attributes.TxHash
							Pair.Transactions = append(Pair.Transactions, PairTransaction)
						}

						if len(Pair.Transactions) > 0 {


							// Get First Transaction
							TransactionReceived, LocalTransactionReceipt := web3.GetTransactionReceipt(Network.RPCs[0], Pair.Transactions[0].Hash)

							if TransactionReceived {

								// Get Router Address
								Dex.RouterAddress = LocalTransactionReceipt.To().Hex()

								// Decode First Transaction To See If We Can - If So We Will Continue
								DexIsValid, _, _ = web3.DecodeTransactionInputData(RouterAbi, LocalTransactionReceipt.Data())

								if DexIsValid {

									// Check If We Got Transactions
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

											// Collect Transactions
											for PairTransactionIndex, PairTransaction := range Pair.Transactions {

												// Decode Pair Transactions
												DecodeSuccessful, Method, DecodedInputData := web3.DecodeTransactionInputData(RouterAbi, LocalTransactionReceipt.Data())

												// Add Data If Decode Successful
												if DecodeSuccessful {
													PairTransaction.InputData = DecodedInputData
													PairTransaction.MethodName = Method
													Pair.Transactions[PairTransactionIndex] = PairTransaction
												} else {
													// Remove It If We Could Decode
													Pair.Transactions = append(Pair.Transactions[:PairTransactionIndex], Pair.Transactions[PairTransactionIndex+1:]...)
												}

											}

											// Check If We Have Any Decoded Transactions
											if len(Pair.Transactions) > 0 {

												if DexDBId < 0 {

													// Check If Dex Is Already Stored
													DexQueryResults := mysql_query.GetDexFromDB(Network.NetworkDBId, Dex.Identifier)

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
												}

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
															log.Panicln("Invalid Token DB Id: ", TokenDBId)
														}

														// Add Pair To DB
														PairDBId = mysql_insert.AddPairToDB(TokenDBId, StablecoinDetails.StablecoinId, Network.NetworkDBId, DexDBId, Pair.Name, Pair.Address)

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

														// Set Pair DB ID
														PairDBId = int64(PairQueryResults[0].PairId)

													} else {

														if TokenDBId < 1 {
															log.Panicln("Invalid Token DB Id: ", TokenDBId)
														}

														// Add Pair To DB
														PairDBId = mysql_insert.AddPairToDB(TokenDBId, StablecoinDetails.StablecoinId, Network.NetworkDBId, DexDBId, Pair.Name, Pair.Address)

													}

												}

												// Add Pair Routes To DB
												for _, PairTransaction := range Pair.Transactions {

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
													RouteQueryResults := mysql_query.GetRouteFromDB(Network.NetworkDBId, DexDBId, int(PairDBId), PairTransaction.Hash)
													if len(RouteQueryResults) <= 0 {

														// Add Route To DB
														mysql_insert.AddRouteToDB(Network.NetworkDBId, DexDBId, PairDBId, RouteString, PairTransaction.MethodName, PairTransaction.Hash, PairTransaction.InputData.AmountIn, PairTransaction.InputData.AmountOutMin)

													}

												}

												log.Printf("    [%d/%d] Added: %v", PairCountIndex, PairCount, Pair.Name)

											}

										} else {

											// Check If Pair Is Blacklisted
											BlacklistPairQueryResults := mysql_query.GetBlacklistPairFromDB(Pair.Address, Network.NetworkDBId)
											if len(BlacklistPairQueryResults) <= 0 {

												// Add Pair To Pair Blacklist DB
												mysql_insert.AddBlacklistPairToDB(Network.NetworkDBId, Pair.Name, Pair.Address)

											}


										}

									}

									DexFailCount = 0

								} else {

									DexFailCount = DexFailCount + 1

									if DexFailCount >= CollectionFailLimit {

										// Check If Dex Is Already Stored
										DexQueryResults := mysql_query.GetDexFromDB(Network.NetworkDBId, Dex.Identifier)

										if len(DexQueryResults) <= 0 {

											// Set Dex DB ID
											DexDBId = int(mysql_insert.AddDexToDB(
												Network.NetworkDBId,
												Dex.Identifier,
												Dex.RouterAddress,
												Dex.FactoryAddress,
												0,
											))

										}

										break

									}

								}

							}

						}

					}

				}

				if DexIsValid {
					log.Printf("    > [%d/%d] [%v] Collected %d Pairs", DexCountIndex, DexCount, Dex.Name, len(Dex.Pairs))
				} else {
					log.Printf("    > [%d/%d] [%v] Skipping", DexCountIndex, DexCount, Dex.Name)
				}

			} else {

				log.Printf("    > [%d/%d] [%v] Skipping", DexCountIndex, DexCount, Dex.Name)

			}

		}

		logging.LogSeparator(false)
		
	}

}
