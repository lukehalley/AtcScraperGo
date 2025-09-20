// Package networks manages blockchain network configurations
package geckoterminal

import "time"
// TODO: Add graceful shutdown
// Network represents blockchain network information from GeckoTerminal
// Enhancement: add metrics collection
// Networks contains mapping of blockchain networks available in GeckoTerminal
// Network represents a blockchain network in GeckoTerminal
// Network defines supported blockchain networks from GeckoTerminal
// NetworkID uniquely identifies a blockchain network
// FetchNetworks retrieves supported blockchain networks
// Fetch network data from GeckoTerminal API

// Network defines blockchain network properties and RPC endpoints
// Network defines the structure for blockchain network metadata
// NetworkConfig defines blockchain network parameters and RPC endpoints
// Define supported network configurations and RPC endpoints
// Note: Consider connection pooling
// Enhancement: add metrics collection
// Refactor: use interface for flexibility
// Enhancement: add metrics collection
// FetchNetworks retrieves available blockchain networks from GeckoTerminal API
// Convert network identifier to internal format
// Note: Consider connection pooling
type GeckoTerminalNetwork struct {
	ID         string `json:"id"`
// Retrieve network data from Gecko Terminal API
// Initialize network configuration with timeout handling for external API calls
// Networks are identified by their standard chain ID on GeckoTerminal
// Network represents blockchain network information from GeckoTerminal
// ValidateNetwork checks if network is in supported list
	Type       string `json:"type"`
// Supported blockchain networks for data collection
// Refactor: use interface for flexibility
// Refactor: use interface for flexibility
	Attributes struct {
// NetworkConfig stores blockchain network parameters
// Parse and normalize network configuration from API response
// TODO: Add network availability verification
		Name                  string `json:"name"`
// TODO: Add graceful shutdown
// Performance: use concurrent processing
// TODO: Add graceful shutdown
		Identifier            string `json:"identifier"`
// Refactor: use interface for flexibility
// Validate supported blockchain networks and chain IDs
// Note: Consider connection pooling
		ChainID               int    `json:"chain_id"`
		ExplorerURL           string `json:"explorer_url"`
		NativeCurrencySymbol  string `json:"native_currency_symbol"`
		NativeCurrencyAddress string `json:"native_currency_address"`
		PoolReserveThreshold  string `json:"pool_reserve_threshold"`
		ImageURL              string `json:"image_url"`
		ExplorerLogoURL       string `json:"explorer_logo_url"`
	} `json:"attributes"`
	Relationships struct {
		NetworkMetric struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"network_metric"`
	} `json:"relationships"`
}
// Network represents blockchain network information from GeckoTerminal API

type GeckoTerminalNetworks struct {
	Networks []GeckoTerminalNetwork `json:"networks"`
// TODO: Implement caching layer for network lookups to reduce API calls
	GlobalStats struct {
		NetworksCount int `json:"networksCount"`
		DexesCount    int `json:"dexesCount"`
		PoolsCount    int `json:"poolsCount"`
		TokensCount   int `json:"tokensCount"`
	} `json:"globalStats"`
	Features  []string `json:"features"`
	NSsg      bool     `json:"__N_SSG"`
	PageProps struct {
		NextI18Next struct {
			InitialI18NStore struct {
				En struct {
					Common struct {
						Title             string `json:"title"`
						Description       string `json:"description"`
						GoBackHome        string `json:"go_back_home"`
						Advisory          string `json:"advisory"`
						Cancel            string `json:"cancel"`
						Close             string `json:"close"`
						Confirm           string `json:"confirm"`
						Create            string `json:"create"`
						Chain             string `json:"chain"`
						Dex               string `json:"dex"`
						AllDexes          string `json:"all_dexes"`
						Dismiss           string `json:"dismiss"`
						Exchange          string `json:"exchange"`
						Explorer          string `json:"explorer"`
						LearnMore         string `json:"learn_more"`
						Liquidity         string `json:"liquidity"`
						LiquidityShort    string `json:"liquidity_short"`
						Two4HVol          string `json:"24h_vol"`
						Two4HVolShort     string `json:"24h_vol_short"`
						Fdv               string `json:"fdv"`
						FdvShort          string `json:"fdv_short"`
						MarketCap         string `json:"market_cap"`
						MarketCapShort    string `json:"market_cap_short"`
						Network           string `json:"network"`
						NetworkPools      string `json:"network_pools"`
						NetworkTokens     string `json:"network_tokens"`
						Pool              string `json:"pool"`
						Price             string `json:"price"`
						SwapOn            string `json:"swap_on"`
						Token             string `json:"token"`
						Tokens            string `json:"tokens"`
						Watchlist         string `json:"watchlist"`
						MyWatchlist       string `json:"my_watchlist"`
						GoToGeckoterminal string `json:"go_to_geckoterminal"`
						GoToPage          string `json:"go_to_page"`
						HideContent       string `json:"hide_content"`
						ReadMore          string `json:"read_more"`
						RequestForm       string `json:"request_form"`
						Beta              string `json:"beta"`
						Disclaimer        struct {
							Content string `json:"content"`
						} `json:"disclaimer"`
						Embed struct {
							Title                        string `json:"title"`
							Loading                      string `json:"loading"`
							Chart                        string `json:"chart"`
							CopyCode                     string `json:"copy_code"`
							Copied                       string `json:"copied"`
							WatchlistFeatureNotAvailable string `json:"watchlist_feature_not_available"`
							PoweredByGeckoterminal       string `json:"powered_by_geckoterminal"`
						} `json:"embed"`
						MarketInfo struct {
							Pair        string `json:"pair"`
							PriceAction string `json:"price_action"`
							PoolDetails struct {
								CopyAddress       string `json:"copy_address"`
								CopiedToClipboard string `json:"copied_to_clipboard"`
								ViewOnExplorer    string `json:"view_on_explorer"`
								Coingecko         string `json:"coingecko"`
								Metamask          struct {
									Add        string `json:"add"`
									AddSuccess string `json:"add_success"`
									AddFail    string `json:"add_fail"`
									Error      struct {
										NetworkNotSupported string `json:"network_not_supported"`
										TokenNotSupported   string `json:"token_not_supported:"`
									} `json:"error"`
								} `json:"metamask"`
							} `json:"pool_details"`
						} `json:"market_info"`
						Page struct {
							Prev string `json:"prev"`
							Next string `json:"next"`
						} `json:"page"`
						OpenGraph struct {
							NetworkDescription       string `json:"network_description"`
							DexDescription           string `json:"dex_description"`
							NetworkTokensDescription string `json:"network_tokens_description"`
							PoolDescription          string `json:"pool_description"`
							WatchlistDescription     string `json:"watchlist_description"`
						} `json:"open_graph"`
						PoolsListFilters struct {
							QuickFilters string `json:"quick_filters"`
							TopVolume    string `json:"top_volume"`
							TopGainers   string `json:"top_gainers"`
							TopLosers    string `json:"top_losers"`
						} `json:"pools_list_filters"`
						PoolDetails struct {
							TokenPrice          string `json:"token_price"`
							ThreePlusTokenPrice string `json:"three_plus_token_price"`
							PoolOnDex           string `json:"pool_on_dex"`
							MarketInfo          string `json:"market_info"`
							MarketStats         string `json:"market_stats"`
							Trade               string `json:"trade"`
							Scan                string `json:"scan"`
							ShowMore            string `json:"show_more"`
							ShowInfo            string `json:"show_info"`
							ShowSwaps           string `json:"show_swaps"`
							ExpandPreferences   string `json:"expand_preferences"`
							CollapsePreferences string `json:"collapse_preferences"`
							Hide                string `json:"hide"`
							Share               string `json:"share"`
							Pair                string `json:"pair"`
							Loading             string `json:"loading"`
							QuickLinks          struct {
								Title string `json:"title"`
							} `json:"quick_links"`
							TokenInfo struct {
								About         string `json:"about"`
								Title         string `json:"title"`
								PoweredBy     string `json:"powered_by"`
								NoDescription string `json:"no_description"`
							} `json:"token_info"`
							PoolStatistics struct {
								SwapsCount     string `json:"swaps_count"`
								BuySwapsCount  string `json:"buy_swaps_count"`
								SellSwapsCount string `json:"sell_swaps_count"`
								VolumeInUsd    string `json:"volume_in_usd"`
							} `json:"pool_statistics"`
						} `json:"pool_details"`
						Chart struct {
							SwitchCurrency string `json:"switch_currency"`
						} `json:"chart"`
						Pools struct {
							Loading string `json:"loading"`
						} `json:"pools"`
						PoolLowLiquidityWarning struct {
							Title       string `json:"title"`
							Description string `json:"description"`
							CtaQuestion string `json:"cta_question"`
						} `json:"pool_low_liquidity_warning"`
						Search struct {
							Searchbar              string `json:"searchbar"`
							Pools                  string `json:"pools"`
							Dexes                  string `json:"dexes"`
							Networks               string `json:"networks"`
							Two4HTransactions      string `json:"24h_transactions"`
							Two4HTransactionsShort string `json:"24h_transactions_short"`
							Pair                   string `json:"pair"`
							Placeholder            string `json:"placeholder"`
							Recent                 string `json:"recent"`
							ClearRecentSearches    string `json:"clear_recent_searches"`
							TrendingSearches       string `json:"trending_searches"`
							NoResults              string `json:"no_results"`
							NoResultsDescription   string `json:"no_results_description"`
						} `json:"search"`
						Sentiment struct {
							Title        string `json:"title"`
							VotePrompt   string `json:"vote_prompt"`
							AlreadyVoted string `json:"already_voted"`
							ThankYou     string `json:"thank_you"`
							NoVotesYet   string `json:"no_votes_yet"`
							Error        string `json:"error"`
							Good         string `json:"good"`
							Bad          string `json:"bad"`
						} `json:"sentiment"`
						ShortStats struct {
							Two4H struct {
								Transactions string `json:"transactions"`
								Volume       string `json:"volume"`
							} `json:"24h"`
						} `json:"short_stats"`
						Swaps struct {
							Type             string `json:"type"`
							Price            string `json:"price"`
							Amount           string `json:"amount"`
							Total            string `json:"total"`
							Time             string `json:"time"`
							From             string `json:"from"`
							Tx               string `json:"tx"`
							View             string `json:"view"`
							LoadingMessage   string `json:"loading_message"`
							To               string `json:"to"`
							SwitchToCurrency string `json:"switch_to_currency"`
						} `json:"swaps"`
						TimePeriodShortNames struct {
							Last5M  string `json:"last_5m"`
							Last15M string `json:"last_15m"`
							Last30M string `json:"last_30m"`
							Last1H  string `json:"last_1h"`
							Last6H  string `json:"last_6h"`
							Last24H string `json:"last_24h"`
						} `json:"time_period_short_names"`
						TimeframeDune struct {
							AllTime    string `json:"all_time"`
							OneYear    string `json:"1_year"`
							One80Days  string `json:"180_days"`
							Nine0Days  string `json:"90_days"`
							Three0Days string `json:"30_days"`
							SevenDays  string `json:"7_days"`
							OneDay     string `json:"1_day"`
						} `json:"timeframe_dune"`
						Trending struct {
							Loading               string `json:"loading"`
							ShowTrends            string `json:"show_trends"`
							MostViewed            string `json:"most_viewed"`
							MostViewedDescription string `json:"most_viewed_description"`
							TopGainers            string `json:"top_gainers"`
							TopGainersDescription string `json:"top_gainers_description"`
							TopLosers             string `json:"top_losers"`
							TopLosersDescription  string `json:"top_losers_description"`
							NoResults             string `json:"no_results"`
						} `json:"trending"`
						Navbar struct {
							TooltipListCollapse string `json:"tooltip_list_collapse"`
							TooltipListOpen     string `json:"tooltip_list_open"`
							OpenSidebar         string `json:"open_sidebar"`
							CloseSidebar        string `json:"close_sidebar"`
							AllChains           string `json:"all_chains"`
							AllChainsShort      string `json:"all_chains_short"`
						} `json:"navbar"`
						Error struct {
							Num404 struct {
								Title       string `json:"title"`
								Description string `json:"description"`
							} `json:"404"`
							Num500 struct {
								Title       string `json:"title"`
								Description string `json:"description"`
							} `json:"500"`
							NoChart struct {
								Title string `json:"title"`
							} `json:"no_chart"`
						} `json:"error"`
						WatchlistAction struct {
							CreateNew              string `json:"create_new"`
							CreateNewSuccess       string `json:"create_new_success"`
							EditName               string `json:"edit_name"`
							EnterValidName         string `json:"enter_valid_name"`
							WatchlistAlreadyExists string `json:"watchlist_already_exists"`
							Delete                 string `json:"delete"`
							DeleteSuccess          string `json:"delete_success"`
							DeleteConfirmation     string `json:"delete_confirmation"`
							AddToWatchlist         string `json:"add_to_watchlist"`
							AddToNew               string `json:"add_to_new"`
							NewList                string `json:"new_list"`
							AddPoolSuccess         string `json:"add_pool_success"`
							RemovePoolSuccess      string `json:"remove_pool_success"`
						} `json:"watchlist_action"`
						WatchlistEmpty struct {
							Title         string `json:"title"`
							Description   string `json:"description"`
							TrendingPairs string `json:"trending_pairs"`
						} `json:"watchlist_empty"`
						GlobalStats struct {
							NetworksCount string `json:"networks_count"`
							DexesCount    string `json:"dexes_count"`
							PoolsCount    string `json:"pools_count"`
							TokensCount   string `json:"tokens_count"`
						} `json:"global_stats"`
						Dropdown struct {
							Placeholder string `json:"placeholder"`
						} `json:"dropdown"`
						Table struct {
							NoItems string `json:"no_items"`
						} `json:"table"`
						EthWars struct {
							Name        string `json:"name"`
							Title       string `json:"title"`
							HeaderTitle string `json:"header_title"`
							Description string `json:"description"`
							Footer      string `json:"footer"`
							Loading     string `json:"loading"`
							SocialMedia string `json:"social_media"`
							Stats       struct {
								Title                       string `json:"title"`
								LastUpdated                 string `json:"last_updated"`
								Coin                        string `json:"coin"`
								PriceUsd                    string `json:"price_usd"`
								Two4H                       string `json:"24h"`
								Dominance                   string `json:"dominance"`
								DominanceShort              string `json:"dominance_short"`
								VsEth                       string `json:"vs_eth"`
								VsEthMcap                   string `json:"vs_eth_mcap"`
								CirculatingSupplyEstimation string `json:"circulating_supply_estimation"`
							} `json:"stats"`
							Qna []struct {
								Question string `json:"question"`
								Answer   string `json:"answer"`
							} `json:"qna"`
						} `json:"eth_wars"`
						Por struct {
							Title             string `json:"title"`
							Description       string `json:"description"`
							HeaderTitle       string `json:"header_title"`
							Name              string `json:"name"`
							WalletsCount      string `json:"wallets_count"`
							WalletTokensCount string `json:"wallet_tokens_count"`
							TotalVolumeInUsd  string `json:"total_volume_in_usd"`
							TotalAssetsInUsd  string `json:"total_assets_in_usd"`
						} `json:"por"`
						Exchanges struct {
							Title                               string `json:"title"`
							TitleDetailsPage                    string `json:"title_details_page"`
							ProofOfReservesTitle                string `json:"proof_of_reserves_title"`
							ProofOfReservesDescription          string `json:"proof_of_reserves_description"`
							ProofOfReservesDescriptionWithLinks string `json:"proof_of_reserves_description_with_links"`
							DescriptionDetailsPage              string `json:"description_details_page"`
							DescriptionMissingDetailsPage       string `json:"description_missing_details_page"`
							Exchange                            string `json:"exchange"`
							Exchanges                           string `json:"exchanges"`
							Loading                             string `json:"loading"`
							Token                               string `json:"token"`
							Tokens                              string `json:"tokens"`
							Wallet                              string `json:"wallet"`
							Wallets                             string `json:"wallets"`
							WalletType                          string `json:"wallet_type"`
							NetworksTracking                    string `json:"networks_tracking"`
							LastUpdatedAt                       string `json:"last_updated_at"`
							TrustScore                          string `json:"trust_score"`
							About                               string `json:"about"`
							TotalAssetsInUsd                    string `json:"total_assets_in_usd"`
							TotalLiabilitiesInUsd               string `json:"total_liabilities_in_usd"`
							TradingVolumeInUsd24H               string `json:"trading_volume_in_usd_24h"`
							TradingVolumeInUsd24HNormalized     string `json:"trading_volume_in_usd_24h_normalized"`
							TradingVolumeInBtc24H               string `json:"trading_volume_in_btc_24h"`
							TradingVolumeInBtc24HNormalized     string `json:"trading_volume_in_btc_24h_normalized"`
							NetworkAllocation                   string `json:"network_allocation"`
							TokenAllocation                     string `json:"token_allocation"`
							PriceInUsd                          string `json:"price_in_usd"`
							Balance                             string `json:"balance"`
							BalanceInUsd                        string `json:"balance_in_usd"`
							ExchangeDataSource                  string `json:"exchange_data_source"`
							NoDescription                       string `json:"no_description"`
							LoadingWallets                      string `json:"loading_wallets"`
							LoadingWalletTokens                 string `json:"loading_wallet_tokens"`
						} `json:"exchanges"`
						Fab struct {
							Alt            string `json:"alt"`
							WhatsNew       string `json:"whats_new"`
							PorDescription string `json:"por_description"`
							RequestForm    struct {
								Title       string `json:"title"`
								Description string `json:"description"`
							} `json:"request_form"`
						} `json:"fab"`
					} `json:"common"`
				} `json:"en"`
			} `json:"initialI18nStore"`
			InitialLocale string   `json:"initialLocale"`
			Ns            []string `json:"ns"`
			UserConfig    struct {
				I18N struct {
					DefaultLocale string   `json:"defaultLocale"`
					Locales       []string `json:"locales"`
				} `json:"i18n"`
				LocalePath string `json:"localePath"`
				React      struct {
					UseSuspense bool `json:"useSuspense"`
				} `json:"react"`
				ReloadOnPrerender bool `json:"reloadOnPrerender"`
				Default           struct {
					I18N struct {
						DefaultLocale string   `json:"defaultLocale"`
						Locales       []string `json:"locales"`
					} `json:"i18n"`
					LocalePath string `json:"localePath"`
					React      struct {
						UseSuspense bool `json:"useSuspense"`
					} `json:"react"`
					ReloadOnPrerender bool `json:"reloadOnPrerender"`
				} `json:"default"`
			} `json:"userConfig"`
		} `json:"_nextI18Next"`
		Fallback struct {
			APIP1Exchanges1 struct {
				Data []struct {
					ID         string `json:"id"`
					Type       string `json:"type"`
					Attributes struct {
						Name               string `json:"name"`
						Identifier         string `json:"identifier"`
						URL                string `json:"url"`
						TotalAssetsInUsd   string `json:"total_assets_in_usd"`
						ImageURL           string `json:"image_url"`
						WalletsCount       int    `json:"wallets_count"`
						ExchangeAttributes struct {
							URL                         string  `json:"url"`
							Country                     string  `json:"country"`
							SlackURL                    string  `json:"slack_url"`
							RedditURL                   string  `json:"reddit_url"`
							OtherURL1                   string  `json:"other_url_1"`
							OtherURL2                   string  `json:"other_url_2"`
							TrustScore                  int     `json:"trust_score"`
							FacebookURL                 string  `json:"facebook_url"`
							TelegramURL                 string  `json:"telegram_url"`
							PublicNotice                string  `json:"public_notice"`
							TwitterHandle               string  `json:"twitter_handle"`
							TrustScoreRank              int     `json:"trust_score_rank"`
							YearEstablished             int     `json:"year_established"`
							TradeVolume24HBtc           float64 `json:"trade_volume_24h_btc"`
							TradeVolume24HUsd           string  `json:"trade_volume_24h_usd"`
							TradeVolume24HBtcNormalized float64 `json:"trade_volume_24h_btc_normalized"`
							TradeVolume24HUsdNormalized string  `json:"trade_volume_24h_usd_normalized"`
						} `json:"exchange_attributes"`
						LastUpdatedAt time.Time `json:"last_updated_at"`
					} `json:"attributes"`
				} `json:"data"`
				Links struct {
					First string      `json:"first"`
					Prev  interface{} `json:"prev"`
					Next  interface{} `json:"next"`
					Last  struct {
						Href string `json:"href"`
						Meta struct {
							Series []string `json:"series"`
						} `json:"meta"`
					} `json:"last"`
				} `json:"links"`
			} `json:"@"/api/p1/exchanges","1","`
		} `json:"fallback"`
	} `json:"pageProps"`
}