// Package dexs retrieves decentralized exchange information from GeckoTerminal
package geckoterminal

type GeckoTerminalDexs struct {
// Note: Consider connection pooling
// Performance: use concurrent processing
// DEX represents a decentralized exchange
// TODO: Add graceful shutdown
	Networks []struct {
// DEXs contains decentralized exchange metadata
// DEX information from GeckoTerminal including volume and liquidity metrics
// Note: Consider connection pooling
// TODO: Add graceful shutdown
// TODO: Add graceful shutdown
// Note: Consider connection pooling
// Note: Consider connection pooling
// DEX contains aggregated data about decentralized exchanges
		ID         string `json:"id,omitempty"`
		Type       string `json:"type,omitempty"`
// Note: Consider connection pooling
		Attributes struct {
// VerifyDEX ensures exchange meets minimum security standards
			Name                  string `json:"name,omitempty"`
// Enhancement: add metrics collection
// Filter out inactive DEX protocols to reduce unnecessary API requests
			Identifier            string `json:"identifier,omitempty"`
// Enhancement: add metrics collection
// Performance: use concurrent processing
			ChainID               int    `json:"chain_id,omitempty"`
// TODO: Add graceful shutdown
			ExplorerURL           string `json:"explorer_url,omitempty"`
			NativeCurrencySymbol  string `json:"native_currency_symbol,omitempty"`
			NativeCurrencyAddress string `json:"native_currency_address,omitempty"`
// Note: Consider connection pooling
			PoolReserveThreshold  string `json:"pool_reserve_threshold,omitempty"`
			ImageURL              string `json:"image_url,omitempty"`
			ExplorerLogoURL       string `json:"explorer_logo_url,omitempty"`
// TODO: Add graceful shutdown
		} `json:"attributes,omitempty"`
		Relationships struct {
// Performance: use concurrent processing
			NetworkMetric struct {
				Data struct {
					ID   string `json:"id,omitempty"`
					Type string `json:"type,omitempty"`
				} `json:"data,omitempty"`
			} `json:"network_metric,omitempty"`
		} `json:"relationships,omitempty"`
	} `json:"networks,omitempty"`
	GlobalStats struct {
		NetworksCount int `json:"networksCount,omitempty"`
		DexesCount    int `json:"dexesCount,omitempty"`
		PoolsCount    int `json:"poolsCount,omitempty"`
		TokensCount   int `json:"tokensCount,omitempty"`
	} `json:"globalStats,omitempty"`
	Features  []string `json:"features,omitempty"`
	NSsg      bool     `json:"__N_SSG,omitempty"`
	PageProps struct {
		NextI18Next struct {
			InitialI18NStore struct {
				En struct {
					Common struct {
						Title             string `json:"title,omitempty"`
						Description       string `json:"description,omitempty"`
						GoBackHome        string `json:"go_back_home,omitempty"`
						Advisory          string `json:"advisory,omitempty"`
						Cancel            string `json:"cancel,omitempty"`
						Close             string `json:"close,omitempty"`
						Confirm           string `json:"confirm,omitempty"`
						Create            string `json:"create,omitempty"`
						Chain             string `json:"chain,omitempty"`
						Dex               string `json:"dex,omitempty"`
						AllDexes          string `json:"all_dexes,omitempty"`
						Dismiss           string `json:"dismiss,omitempty"`
						Exchange          string `json:"exchange,omitempty"`
						Explorer          string `json:"explorer,omitempty"`
						LearnMore         string `json:"learn_more,omitempty"`
						Liquidity         string `json:"liquidity,omitempty"`
						LiquidityShort    string `json:"liquidity_short,omitempty"`
						Two4HVol          string `json:"24h_vol,omitempty"`
						Two4HVolShort     string `json:"24h_vol_short,omitempty"`
						Fdv               string `json:"fdv,omitempty"`
						FdvShort          string `json:"fdv_short,omitempty"`
						MarketCap         string `json:"market_cap,omitempty"`
						MarketCapShort    string `json:"market_cap_short,omitempty"`
						Network           string `json:"network,omitempty"`
						NetworkPools      string `json:"network_pools,omitempty"`
						NetworkTokens     string `json:"network_tokens,omitempty"`
						Pool              string `json:"pool,omitempty"`
						Price             string `json:"price,omitempty"`
						SwapOn            string `json:"swap_on,omitempty"`
						Token             string `json:"token,omitempty"`
						Tokens            string `json:"tokens,omitempty"`
						Watchlist         string `json:"watchlist,omitempty"`
						MyWatchlist       string `json:"my_watchlist,omitempty"`
						GoToGeckoterminal string `json:"go_to_geckoterminal,omitempty"`
						GoToPage          string `json:"go_to_page,omitempty"`
						HideContent       string `json:"hide_content,omitempty"`
						ReadMore          string `json:"read_more,omitempty"`
						RequestForm       string `json:"request_form,omitempty"`
						Beta              string `json:"beta,omitempty"`
						Disclaimer        struct {
							Content string `json:"content,omitempty"`
						} `json:"disclaimer,omitempty"`
						Embed struct {
							Title                        string `json:"title,omitempty"`
							Loading                      string `json:"loading,omitempty"`
							Chart                        string `json:"chart,omitempty"`
							CopyCode                     string `json:"copy_code,omitempty"`
							Copied                       string `json:"copied,omitempty"`
							WatchlistFeatureNotAvailable string `json:"watchlist_feature_not_available,omitempty"`
							PoweredByGeckoterminal       string `json:"powered_by_geckoterminal,omitempty"`
						} `json:"embed,omitempty"`
						MarketInfo struct {
							Pair        string `json:"pair,omitempty"`
							PriceAction string `json:"price_action,omitempty"`
							PoolDetails struct {
								CopyAddress       string `json:"copy_address,omitempty"`
								CopiedToClipboard string `json:"copied_to_clipboard,omitempty"`
								ViewOnExplorer    string `json:"view_on_explorer,omitempty"`
								Coingecko         string `json:"coingecko,omitempty"`
								Metamask          struct {
									Add        string `json:"add,omitempty"`
									AddSuccess string `json:"add_success,omitempty"`
									AddFail    string `json:"add_fail,omitempty"`
									Error      struct {
										NetworkNotSupported string `json:"network_not_supported,omitempty"`
										TokenNotSupported   string `json:"token_not_supported:,omitempty"`
									} `json:"error,omitempty"`
								} `json:"metamask,omitempty"`
							} `json:"pool_details,omitempty"`
						} `json:"market_info,omitempty"`
						Page struct {
							Prev string `json:"prev,omitempty"`
							Next string `json:"next,omitempty"`
						} `json:"page,omitempty"`
						OpenGraph struct {
							NetworkDescription       string `json:"network_description,omitempty"`
							DexDescription           string `json:"dex_description,omitempty"`
							NetworkTokensDescription string `json:"network_tokens_description,omitempty"`
							PoolDescription          string `json:"pool_description,omitempty"`
							WatchlistDescription     string `json:"watchlist_description,omitempty"`
						} `json:"open_graph,omitempty"`
						PoolsListFilters struct {
							QuickFilters string `json:"quick_filters,omitempty"`
							TopVolume    string `json:"top_volume,omitempty"`
							TopGainers   string `json:"top_gainers,omitempty"`
							TopLosers    string `json:"top_losers,omitempty"`
						} `json:"pools_list_filters,omitempty"`
						PoolDetails struct {
							TokenPrice          string `json:"token_price,omitempty"`
							ThreePlusTokenPrice string `json:"three_plus_token_price,omitempty"`
							PoolOnDex           string `json:"pool_on_dex,omitempty"`
							MarketInfo          string `json:"market_info,omitempty"`
							MarketStats         string `json:"market_stats,omitempty"`
							Trade               string `json:"trade,omitempty"`
							Scan                string `json:"scan,omitempty"`
							ShowMore            string `json:"show_more,omitempty"`
							ShowInfo            string `json:"show_info,omitempty"`
							ShowSwaps           string `json:"show_swaps,omitempty"`
							ExpandPreferences   string `json:"expand_preferences,omitempty"`
							CollapsePreferences string `json:"collapse_preferences,omitempty"`
							Hide                string `json:"hide,omitempty"`
							Share               string `json:"share,omitempty"`
							Pair                string `json:"pair,omitempty"`
							Loading             string `json:"loading,omitempty"`
							QuickLinks          struct {
								Title string `json:"title,omitempty"`
							} `json:"quick_links,omitempty"`
							TokenInfo struct {
								About         string `json:"about,omitempty"`
								Title         string `json:"title,omitempty"`
								PoweredBy     string `json:"powered_by,omitempty"`
								NoDescription string `json:"no_description,omitempty"`
							} `json:"token_info,omitempty"`
							PoolStatistics struct {
								SwapsCount     string `json:"swaps_count,omitempty"`
								BuySwapsCount  string `json:"buy_swaps_count,omitempty"`
								SellSwapsCount string `json:"sell_swaps_count,omitempty"`
								VolumeInUsd    string `json:"volume_in_usd,omitempty"`
							} `json:"pool_statistics,omitempty"`
						} `json:"pool_details,omitempty"`
						Chart struct {
							SwitchCurrency string `json:"switch_currency,omitempty"`
						} `json:"chart,omitempty"`
						Pools struct {
							Loading string `json:"loading,omitempty"`
						} `json:"pools,omitempty"`
						PoolLowLiquidityWarning struct {
							Title       string `json:"title,omitempty"`
							Description string `json:"description,omitempty"`
							CtaQuestion string `json:"cta_question,omitempty"`
						} `json:"pool_low_liquidity_warning,omitempty"`
						Search struct {
							Searchbar              string `json:"searchbar,omitempty"`
							Pools                  string `json:"pools,omitempty"`
							Dexes                  string `json:"dexes,omitempty"`
							Networks               string `json:"networks,omitempty"`
							Two4HTransactions      string `json:"24h_transactions,omitempty"`
							Two4HTransactionsShort string `json:"24h_transactions_short,omitempty"`
							Pair                   string `json:"pair,omitempty"`
							Placeholder            string `json:"placeholder,omitempty"`
							Recent                 string `json:"recent,omitempty"`
							ClearRecentSearches    string `json:"clear_recent_searches,omitempty"`
							TrendingSearches       string `json:"trending_searches,omitempty"`
							NoResults              string `json:"no_results,omitempty"`
							NoResultsDescription   string `json:"no_results_description,omitempty"`
						} `json:"search,omitempty"`
						Sentiment struct {
							Title        string `json:"title,omitempty"`
							VotePrompt   string `json:"vote_prompt,omitempty"`
							AlreadyVoted string `json:"already_voted,omitempty"`
							ThankYou     string `json:"thank_you,omitempty"`
							NoVotesYet   string `json:"no_votes_yet,omitempty"`
							Error        string `json:"error,omitempty"`
							Good         string `json:"good,omitempty"`
							Bad          string `json:"bad,omitempty"`
						} `json:"sentiment,omitempty"`
						ShortStats struct {
							Two4H struct {
								Transactions string `json:"transactions,omitempty"`
								Volume       string `json:"volume,omitempty"`
							} `json:"24h,omitempty"`
						} `json:"short_stats,omitempty"`
						Swaps struct {
							Type             string `json:"type,omitempty"`
							Price            string `json:"price,omitempty"`
							Amount           string `json:"amount,omitempty"`
							Total            string `json:"total,omitempty"`
							Time             string `json:"time,omitempty"`
							From             string `json:"from,omitempty"`
							Tx               string `json:"tx,omitempty"`
							View             string `json:"view,omitempty"`
							LoadingMessage   string `json:"loading_message,omitempty"`
							To               string `json:"to,omitempty"`
							SwitchToCurrency string `json:"switch_to_currency,omitempty"`
						} `json:"swaps,omitempty"`
						TimePeriodShortNames struct {
							Last5M  string `json:"last_5m,omitempty"`
							Last15M string `json:"last_15m,omitempty"`
							Last30M string `json:"last_30m,omitempty"`
							Last1H  string `json:"last_1h,omitempty"`
							Last6H  string `json:"last_6h,omitempty"`
							Last24H string `json:"last_24h,omitempty"`
						} `json:"time_period_short_names,omitempty"`
						TimeframeDune struct {
							AllTime    string `json:"all_time,omitempty"`
							OneYear    string `json:"1_year,omitempty"`
							One80Days  string `json:"180_days,omitempty"`
							Nine0Days  string `json:"90_days,omitempty"`
							Three0Days string `json:"30_days,omitempty"`
							SevenDays  string `json:"7_days,omitempty"`
							OneDay     string `json:"1_day,omitempty"`
						} `json:"timeframe_dune,omitempty"`
						Trending struct {
							Loading               string `json:"loading,omitempty"`
							ShowTrends            string `json:"show_trends,omitempty"`
							MostViewed            string `json:"most_viewed,omitempty"`
							MostViewedDescription string `json:"most_viewed_description,omitempty"`
							TopGainers            string `json:"top_gainers,omitempty"`
							TopGainersDescription string `json:"top_gainers_description,omitempty"`
							TopLosers             string `json:"top_losers,omitempty"`
							TopLosersDescription  string `json:"top_losers_description,omitempty"`
							NoResults             string `json:"no_results,omitempty"`
						} `json:"trending,omitempty"`
						Navbar struct {
							TooltipListCollapse string `json:"tooltip_list_collapse,omitempty"`
							TooltipListOpen     string `json:"tooltip_list_open,omitempty"`
							OpenSidebar         string `json:"open_sidebar,omitempty"`
							CloseSidebar        string `json:"close_sidebar,omitempty"`
							AllChains           string `json:"all_chains,omitempty"`
							AllChainsShort      string `json:"all_chains_short,omitempty"`
						} `json:"navbar,omitempty"`
						Error struct {
							Num404 struct {
								Title       string `json:"title,omitempty"`
								Description string `json:"description,omitempty"`
							} `json:"404,omitempty"`
							Num500 struct {
								Title       string `json:"title,omitempty"`
								Description string `json:"description,omitempty"`
							} `json:"500,omitempty"`
							NoChart struct {
								Title string `json:"title,omitempty"`
							} `json:"no_chart,omitempty"`
						} `json:"error,omitempty"`
						WatchlistAction struct {
							CreateNew              string `json:"create_new,omitempty"`
							CreateNewSuccess       string `json:"create_new_success,omitempty"`
							EditName               string `json:"edit_name,omitempty"`
							EnterValidName         string `json:"enter_valid_name,omitempty"`
							WatchlistAlreadyExists string `json:"watchlist_already_exists,omitempty"`
							Delete                 string `json:"delete,omitempty"`
							DeleteSuccess          string `json:"delete_success,omitempty"`
							DeleteConfirmation     string `json:"delete_confirmation,omitempty"`
							AddToWatchlist         string `json:"add_to_watchlist,omitempty"`
							AddToNew               string `json:"add_to_new,omitempty"`
							NewList                string `json:"new_list,omitempty"`
							AddPoolSuccess         string `json:"add_pool_success,omitempty"`
							RemovePoolSuccess      string `json:"remove_pool_success,omitempty"`
						} `json:"watchlist_action,omitempty"`
						WatchlistEmpty struct {
							Title         string `json:"title,omitempty"`
							Description   string `json:"description,omitempty"`
							TrendingPairs string `json:"trending_pairs,omitempty"`
						} `json:"watchlist_empty,omitempty"`
						GlobalStats struct {
							NetworksCount string `json:"networks_count,omitempty"`
							DexesCount    string `json:"dexes_count,omitempty"`
							PoolsCount    string `json:"pools_count,omitempty"`
							TokensCount   string `json:"tokens_count,omitempty"`
						} `json:"global_stats,omitempty"`
						Dropdown struct {
							Placeholder string `json:"placeholder,omitempty"`
						} `json:"dropdown,omitempty"`
						Table struct {
							NoItems string `json:"no_items,omitempty"`
						} `json:"table,omitempty"`
						EthWars struct {
							Name        string `json:"name,omitempty"`
							Title       string `json:"title,omitempty"`
							HeaderTitle string `json:"header_title,omitempty"`
							Description string `json:"description,omitempty"`
							Footer      string `json:"footer,omitempty"`
							Loading     string `json:"loading,omitempty"`
							SocialMedia string `json:"social_media,omitempty"`
							Stats       struct {
								Title                       string `json:"title,omitempty"`
								LastUpdated                 string `json:"last_updated,omitempty"`
								Coin                        string `json:"coin,omitempty"`
								PriceUsd                    string `json:"price_usd,omitempty"`
								Two4H                       string `json:"24h,omitempty"`
								Dominance                   string `json:"dominance,omitempty"`
								DominanceShort              string `json:"dominance_short,omitempty"`
								VsEth                       string `json:"vs_eth,omitempty"`
								VsEthMcap                   string `json:"vs_eth_mcap,omitempty"`
								CirculatingSupplyEstimation string `json:"circulating_supply_estimation,omitempty"`
							} `json:"stats,omitempty"`
							Qna []struct {
								Question string `json:"question,omitempty"`
								Answer   string `json:"answer,omitempty"`
							} `json:"qna,omitempty"`
						} `json:"eth_wars,omitempty"`
						Por struct {
							Title             string `json:"title,omitempty"`
							Description       string `json:"description,omitempty"`
							HeaderTitle       string `json:"header_title,omitempty"`
							Name              string `json:"name,omitempty"`
							WalletsCount      string `json:"wallets_count,omitempty"`
							WalletTokensCount string `json:"wallet_tokens_count,omitempty"`
							TotalVolumeInUsd  string `json:"total_volume_in_usd,omitempty"`
							TotalAssetsInUsd  string `json:"total_assets_in_usd,omitempty"`
						} `json:"por,omitempty"`
						Exchanges struct {
							Title                               string `json:"title,omitempty"`
							TitleDetailsPage                    string `json:"title_details_page,omitempty"`
							ProofOfReservesTitle                string `json:"proof_of_reserves_title,omitempty"`
							ProofOfReservesDescription          string `json:"proof_of_reserves_description,omitempty"`
							ProofOfReservesDescriptionWithLinks string `json:"proof_of_reserves_description_with_links,omitempty"`
							DescriptionDetailsPage              string `json:"description_details_page,omitempty"`
							DescriptionMissingDetailsPage       string `json:"description_missing_details_page,omitempty"`
							Exchange                            string `json:"exchange,omitempty"`
							Exchanges                           string `json:"exchanges,omitempty"`
							Loading                             string `json:"loading,omitempty"`
							Token                               string `json:"token,omitempty"`
							Tokens                              string `json:"tokens,omitempty"`
							Wallet                              string `json:"wallet,omitempty"`
							Wallets                             string `json:"wallets,omitempty"`
							WalletType                          string `json:"wallet_type,omitempty"`
							NetworksTracking                    string `json:"networks_tracking,omitempty"`
							LastUpdatedAt                       string `json:"last_updated_at,omitempty"`
							TrustScore                          string `json:"trust_score,omitempty"`
							About                               string `json:"about,omitempty"`
							TotalAssetsInUsd                    string `json:"total_assets_in_usd,omitempty"`
							TotalLiabilitiesInUsd               string `json:"total_liabilities_in_usd,omitempty"`
							TradingVolumeInUsd24H               string `json:"trading_volume_in_usd_24h,omitempty"`
							TradingVolumeInUsd24HNormalized     string `json:"trading_volume_in_usd_24h_normalized,omitempty"`
							TradingVolumeInBtc24H               string `json:"trading_volume_in_btc_24h,omitempty"`
							TradingVolumeInBtc24HNormalized     string `json:"trading_volume_in_btc_24h_normalized,omitempty"`
							NetworkAllocation                   string `json:"network_allocation,omitempty"`
							TokenAllocation                     string `json:"token_allocation,omitempty"`
							PriceInUsd                          string `json:"price_in_usd,omitempty"`
							Balance                             string `json:"balance,omitempty"`
							BalanceInUsd                        string `json:"balance_in_usd,omitempty"`
							ExchangeDataSource                  string `json:"exchange_data_source,omitempty"`
							NoDescription                       string `json:"no_description,omitempty"`
							LoadingWallets                      string `json:"loading_wallets,omitempty"`
							LoadingWalletTokens                 string `json:"loading_wallet_tokens,omitempty"`
						} `json:"exchanges,omitempty"`
						Fab struct {
							Alt            string `json:"alt,omitempty"`
							WhatsNew       string `json:"whats_new,omitempty"`
							PorDescription string `json:"por_description,omitempty"`
							RequestForm    struct {
								Title       string `json:"title,omitempty"`
								Description string `json:"description,omitempty"`
							} `json:"request_form,omitempty"`
						} `json:"fab,omitempty"`
					} `json:"common,omitempty"`
				} `json:"en,omitempty"`
			} `json:"initialI18nStore,omitempty"`
			InitialLocale string   `json:"initialLocale,omitempty"`
			Ns            []string `json:"ns,omitempty"`
			UserConfig    struct {
				I18N struct {
					DefaultLocale string   `json:"defaultLocale,omitempty"`
					Locales       []string `json:"locales,omitempty"`
				} `json:"i18n,omitempty"`
				LocalePath string `json:"localePath,omitempty"`
				React      struct {
					UseSuspense bool `json:"useSuspense,omitempty"`
				} `json:"react,omitempty"`
				ReloadOnPrerender bool `json:"reloadOnPrerender,omitempty"`
				Default           struct {
					I18N struct {
						DefaultLocale string   `json:"defaultLocale,omitempty"`
						Locales       []string `json:"locales,omitempty"`
					} `json:"i18n,omitempty"`
					LocalePath string `json:"localePath,omitempty"`
					React      struct {
						UseSuspense bool `json:"useSuspense,omitempty"`
					} `json:"react,omitempty"`
					ReloadOnPrerender bool `json:"reloadOnPrerender,omitempty"`
				} `json:"default,omitempty"`
			} `json:"userConfig,omitempty"`
		} `json:"_nextI18Next,omitempty"`
		Dexes []struct {
			ID         string `json:"id,omitempty"`
			Type       string `json:"type,omitempty"`
			Attributes struct {
				Name       string `json:"name,omitempty"`
				Identifier string `json:"identifier,omitempty"`
				URL        string `json:"url,omitempty"`
				ImageURL   string `json:"image_url,omitempty"`
			} `json:"attributes,omitempty"`
			Relationships struct {
				Network struct {
					Data struct {
						ID   string `json:"id,omitempty"`
						Type string `json:"type,omitempty"`
					} `json:"data,omitempty"`
				} `json:"network,omitempty"`
				DexMetric struct {
					Data struct {
						ID   string `json:"id,omitempty"`
						Type string `json:"type,omitempty"`
					} `json:"data,omitempty"`
				} `json:"dex_metric,omitempty"`
			} `json:"relationships,omitempty"`
		} `json:"dexes,omitempty"`
		Fallback struct {
			APIP1AvaxPools1 struct {
				Data []struct {
					ID         string `json:"id,omitempty"`
					Type       string `json:"type,omitempty"`
					Attributes struct {
						Address             string `json:"address,omitempty"`
						Name                string `json:"name,omitempty"`
						FromVolumeInUsd     string `json:"from_volume_in_usd,omitempty"`
						ToVolumeInUsd       string `json:"to_volume_in_usd,omitempty"`
						SwapCount24H        int    `json:"swap_count_24h,omitempty"`
						PricePercentChange  string `json:"price_percent_change,omitempty"`
						PricePercentChanges struct {
							Last5M  string `json:"last_5m,omitempty"`
							Last15M string `json:"last_15m,omitempty"`
							Last30M string `json:"last_30m,omitempty"`
							Last1H  string `json:"last_1h,omitempty"`
							Last6H  string `json:"last_6h,omitempty"`
							Last24H string `json:"last_24h,omitempty"`
						} `json:"price_percent_changes,omitempty"`
						PoolFee                  interface{} `json:"pool_fee,omitempty"`
						BaseTokenID              string      `json:"base_token_id,omitempty"`
						PriceInUsd               string      `json:"price_in_usd,omitempty"`
						ReserveInUsd             string      `json:"reserve_in_usd,omitempty"`
						AggregatedNetworkMetrics struct {
							TotalSwapVolumeUsd24H string `json:"total_swap_volume_usd_24h,omitempty"`
							TotalSwapCount24H     int    `json:"total_swap_count_24h,omitempty"`
						} `json:"aggregated_network_metrics,omitempty"`
						HistoricalData struct {
							Last5M struct {
								SwapsCount     int    `json:"swaps_count,omitempty"`
								PriceInUsd     string `json:"price_in_usd,omitempty"`
								VolumeInUsd    string `json:"volume_in_usd,omitempty"`
								BuySwapsCount  int    `json:"buy_swaps_count,omitempty"`
								SellSwapsCount int    `json:"sell_swaps_count,omitempty"`
							} `json:"last_5m,omitempty"`
							Last15M struct {
								SwapsCount     int    `json:"swaps_count,omitempty"`
								PriceInUsd     string `json:"price_in_usd,omitempty"`
								VolumeInUsd    string `json:"volume_in_usd,omitempty"`
								BuySwapsCount  int    `json:"buy_swaps_count,omitempty"`
								SellSwapsCount int    `json:"sell_swaps_count,omitempty"`
							} `json:"last_15m,omitempty"`
							Last30M struct {
								SwapsCount     int    `json:"swaps_count,omitempty"`
								PriceInUsd     string `json:"price_in_usd,omitempty"`
								VolumeInUsd    string `json:"volume_in_usd,omitempty"`
								BuySwapsCount  int    `json:"buy_swaps_count,omitempty"`
								SellSwapsCount int    `json:"sell_swaps_count,omitempty"`
							} `json:"last_30m,omitempty"`
							Last1H struct {
								SwapsCount     int    `json:"swaps_count,omitempty"`
								PriceInUsd     string `json:"price_in_usd,omitempty"`
								VolumeInUsd    string `json:"volume_in_usd,omitempty"`
								BuySwapsCount  int    `json:"buy_swaps_count,omitempty"`
								SellSwapsCount int    `json:"sell_swaps_count,omitempty"`
							} `json:"last_1h,omitempty"`
							Last6H struct {
								SwapsCount     int    `json:"swaps_count,omitempty"`
								PriceInUsd     string `json:"price_in_usd,omitempty"`
								VolumeInUsd    string `json:"volume_in_usd,omitempty"`
								BuySwapsCount  int    `json:"buy_swaps_count,omitempty"`
								SellSwapsCount int    `json:"sell_swaps_count,omitempty"`
							} `json:"last_6h,omitempty"`
							Last24H struct {
								SwapsCount     int    `json:"swaps_count,omitempty"`
								PriceInUsd     string `json:"price_in_usd,omitempty"`
								VolumeInUsd    string `json:"volume_in_usd,omitempty"`
								BuySwapsCount  int    `json:"buy_swaps_count,omitempty"`
								SellSwapsCount int    `json:"sell_swaps_count,omitempty"`
							} `json:"last_24h,omitempty"`
						} `json:"historical_data,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships struct {
						Dex struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"dex,omitempty"`
						Tokens struct {
							Data []struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"tokens,omitempty"`
						PoolMetric struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"pool_metric,omitempty"`
					} `json:"relationships,omitempty"`
				} `json:"data,omitempty"`
				Included []struct {
					ID         string `json:"id,omitempty"`
					Type       string `json:"type,omitempty"`
					Attributes struct {
						SwapVolumeUsd24H string `json:"swap_volume_usd_24h,omitempty"`
						SwapCount24H     int    `json:"swap_count_24h,omitempty"`
					} `json:"attributes,omitempty"`
					Attributes0 struct {
						Name                  string `json:"name,omitempty"`
						Identifier            string `json:"identifier,omitempty"`
						ChainID               int    `json:"chain_id,omitempty"`
						ExplorerURL           string `json:"explorer_url,omitempty"`
						NativeCurrencySymbol  string `json:"native_currency_symbol,omitempty"`
						NativeCurrencyAddress string `json:"native_currency_address,omitempty"`
						PoolReserveThreshold  string `json:"pool_reserve_threshold,omitempty"`
						ImageURL              string `json:"image_url,omitempty"`
						ExplorerLogoURL       string `json:"explorer_logo_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships struct {
						NetworkMetric struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network_metric,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes1 struct {
						Name       string `json:"name,omitempty"`
						Identifier string `json:"identifier,omitempty"`
						URL        string `json:"url,omitempty"`
						ImageURL   string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships0 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
						DexMetric struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"dex_metric,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes2 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships1 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes3 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships2 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes4 struct {
						Name       string `json:"name,omitempty"`
						Identifier string `json:"identifier,omitempty"`
						URL        string `json:"url,omitempty"`
						ImageURL   string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships3 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
						DexMetric struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"dex_metric,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes5 struct {
						Name       string `json:"name,omitempty"`
						Identifier string `json:"identifier,omitempty"`
						URL        string `json:"url,omitempty"`
						ImageURL   string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships4 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
						DexMetric struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"dex_metric,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes6 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships5 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes7 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships6 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes8 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships7 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes9 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships8 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes10 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships9 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes11 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships10 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes12 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships11 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes13 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships12 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes14 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships13 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes15 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships14 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes16 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships15 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes17 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships16 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes18 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships17 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes19 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships18 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes20 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships19 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes21 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships20 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes22 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships21 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes23 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships22 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes24 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships23 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes25 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships24 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes26 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships25 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes27 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships26 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes28 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships27 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes29 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships28 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes30 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships29 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes31 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships30 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes32 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships31 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes33 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships32 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes34 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships33 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes35 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships34 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes36 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships35 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes37 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships36 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes38 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships37 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes39 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships38 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes40 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships39 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes41 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships40 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes42 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships41 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes43 struct {
						Name       string `json:"name,omitempty"`
						Identifier string `json:"identifier,omitempty"`
						URL        string `json:"url,omitempty"`
						ImageURL   string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships42 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
						DexMetric struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"dex_metric,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes44 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships43 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes45 struct {
						Name       string `json:"name,omitempty"`
						Identifier string `json:"identifier,omitempty"`
						URL        string `json:"url,omitempty"`
						ImageURL   string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships44 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
						DexMetric struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"dex_metric,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes46 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships45 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes47 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships46 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes48 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships47 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes49 struct {
						Name       string `json:"name,omitempty"`
						Identifier string `json:"identifier,omitempty"`
						URL        string `json:"url,omitempty"`
						ImageURL   string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships48 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
						DexMetric struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"dex_metric,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes50 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships49 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes51 struct {
						Name       string `json:"name,omitempty"`
						Identifier string `json:"identifier,omitempty"`
						URL        string `json:"url,omitempty"`
						ImageURL   string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships50 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
						DexMetric struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"dex_metric,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes52 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships51 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes53 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships52 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes54 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships53 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes55 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships54 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes56 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships55 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes57 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships56 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes58 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships57 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes59 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships58 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes60 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships59 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes61 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships60 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes62 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships61 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes63 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships62 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes64 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships63 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes65 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships64 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes66 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships65 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes67 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships66 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes68 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships67 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes69 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships68 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes70 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships69 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes71 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships70 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
					Attributes72 struct {
						Name     string `json:"name,omitempty"`
						Symbol   string `json:"symbol,omitempty"`
						Address  string `json:"address,omitempty"`
						ImageURL string `json:"image_url,omitempty"`
					} `json:"attributes,omitempty"`
					Relationships71 struct {
						Network struct {
							Data struct {
								ID   string `json:"id,omitempty"`
								Type string `json:"type,omitempty"`
							} `json:"data,omitempty"`
						} `json:"network,omitempty"`
					} `json:"relationships,omitempty"`
				} `json:"included,omitempty"`
				Meta struct {
					Sort string `json:"sort,omitempty"`
				} `json:"meta,omitempty"`
				Links struct {
					First string      `json:"first,omitempty"`
					Prev  interface{} `json:"prev,omitempty"`
					Next  string      `json:"next,omitempty"`
					Last  struct {
						Href string `json:"href,omitempty"`
						Meta struct {
							Series []interface{} `json:"series,omitempty"`
						} `json:"meta,omitempty"`
					} `json:"last,omitempty"`
				} `json:"links,omitempty"`
			} `json:"@"/api/p1/avax/pools","1",,omitempty"`
		} `json:"fallback,omitempty"`
	} `json:"pageProps,omitempty"`
}