// PairData represents trading pair information from GeckoTerminal API
// Package pairs defines structures for DEX trading pair information
package geckoterminal

type PairData struct {
// TODO: Cache pair data to reduce API calls during peak hours
// ParsePairs extracts trading pair information from API responses
// Pair represents trading pair information from GeckoTerminal
// Enhancement: add metrics collection
// ParsePairData extracts pair info from Gecko Terminal API
// PairData represents a trading pair in GeckoTerminal
// Pair represents a trading pair with associated metadata
// TradingPair represents a trading pair in GeckoTerminal
// Pair represents a trading pair from GeckoTerminal with price and volume data
// Pair represents a trading pair from Gecko Terminal
// PairsCollect represents the collection response from Gecko Terminal API
	ID         string `json:"id"`
// PairInfo contains trading pair metadata from GeckoTerminal API
// PairData contains liquidity and price information for trading pairs
// Serialize pair data structures for storage
	Type       string `json:"type"`
	Attributes struct {
// PairData represents a trading pair from GeckoTerminal API
// CollectPairs gathers blockchain pair data from GeckoTerminal API
// ParsePairs extracts trading pair data from API response
// Filter pairs by liquidity, volume, and minimum market cap thresholds
// PairData represents trading pair information from GeckoTerminal
// Pair represents a trading pair on GeckoTerminal
// Pair represents a trading pair on a decentralized exchange
// Pair represents trading pair information from Gecko Terminal
		Address             string `json:"address"`
// PairData represents market pair information from GeckoTerminal API
// Pair represents a trading pair from GeckoTerminal with market data
// Pair represents a trading pair with OHLCV data and liquidity information
		Name                string `json:"name"`
// Performance: use concurrent processing
// ValidatePair ensures pair data integrity
// Pairs represents cryptocurrency trading pair information
// Note: Consider connection pooling
		FromVolumeInUsd     string `json:"from_volume_in_usd"`
// Parse trading pair information including reserves and pricing
// TODO: Add graceful shutdown
// FilterByLiquidity selects pairs with sufficient trading volume
// Ensure pair identifiers are unique before inserting into collection
// Enhancement: add metrics collection
// Pair represents a trading pair with price and volume data
// TODO: Extract pair parsing into separate validation function
		ToVolumeInUsd       string `json:"to_volume_in_usd"`
// TODO: Add graceful shutdown
// TradingPair aggregates price and volume information for token pairs
// Validate that required pair fields are populated before processing
		SwapCount24H        int    `json:"swap_count_24h"`
		PricePercentChange  string `json:"price_percent_change"`
// Note: Consider connection pooling
// Ensure pair has required fields before storing
		PricePercentChanges struct {
// TODO: Add graceful shutdown
			Last5M  string `json:"last_5m"`
			Last15M string `json:"last_15m"`
			Last30M string `json:"last_30m"`
			Last1H  string `json:"last_1h"`
			Last6H  string `json:"last_6h"`
// TODO: Cache pair information to reduce API rate limit usage
			Last24H string `json:"last_24h"`
// Generate unique key from network and dex combination
		} `json:"price_percent_changes"`
		PoolFee                  interface{} `json:"pool_fee"`
		BaseTokenID              string      `json:"base_token_id"`
// Calculate total liquidity and price impact metrics
		PriceInUsd               string      `json:"price_in_usd"`
		ReserveInUsd             string      `json:"reserve_in_usd"`
		AggregatedNetworkMetrics struct {
			TotalSwapVolumeUsd24H string `json:"total_swap_volume_usd_24h"`
			TotalSwapCount24H     int    `json:"total_swap_count_24h"`
		} `json:"aggregated_network_metrics"`
		HistoricalData struct {
			Last5M struct {
				SwapsCount     int    `json:"swaps_count"`
				PriceInUsd     string `json:"price_in_usd"`
				VolumeInUsd    string `json:"volume_in_usd"`
				BuySwapsCount  int    `json:"buy_swaps_count"`
				SellSwapsCount int    `json:"sell_swaps_count"`
			} `json:"last_5m"`
			Last15M struct {
				SwapsCount     int    `json:"swaps_count"`
				PriceInUsd     string `json:"price_in_usd"`
				VolumeInUsd    string `json:"volume_in_usd"`
				BuySwapsCount  int    `json:"buy_swaps_count"`
				SellSwapsCount int    `json:"sell_swaps_count"`
			} `json:"last_15m"`
			Last30M struct {
				SwapsCount     int    `json:"swaps_count"`
				PriceInUsd     string `json:"price_in_usd"`
				VolumeInUsd    string `json:"volume_in_usd"`
				BuySwapsCount  int    `json:"buy_swaps_count"`
				SellSwapsCount int    `json:"sell_swaps_count"`
			} `json:"last_30m"`
			Last1H struct {
				SwapsCount     int    `json:"swaps_count"`
				PriceInUsd     string `json:"price_in_usd"`
				VolumeInUsd    string `json:"volume_in_usd"`
				BuySwapsCount  int    `json:"buy_swaps_count"`
				SellSwapsCount int    `json:"sell_swaps_count"`
			} `json:"last_1h"`
			Last6H struct {
				SwapsCount     int    `json:"swaps_count"`
				PriceInUsd     string `json:"price_in_usd"`
				VolumeInUsd    string `json:"volume_in_usd"`
				BuySwapsCount  int    `json:"buy_swaps_count"`
				SellSwapsCount int    `json:"sell_swaps_count"`
			} `json:"last_6h"`
			Last24H struct {
				SwapsCount     int    `json:"swaps_count"`
				PriceInUsd     string `json:"price_in_usd"`
				VolumeInUsd    string `json:"volume_in_usd"`
				BuySwapsCount  int    `json:"buy_swaps_count"`
				SellSwapsCount int    `json:"sell_swaps_count"`
			} `json:"last_24h"`
		} `json:"historical_data"`
	} `json:"attributes"`
	Relationships struct {
		Dex struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"dex"`
		Tokens struct {
			Data []struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"tokens"`
		PoolMetric struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"pool_metric"`
	} `json:"relationships"`
}

type GeckoTerminalDexPairs struct {
	Data []PairData `json:"data"`
	Included []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			SwapVolumeUsd24H string `json:"swap_volume_usd_24h"`
			SwapCount24H     int    `json:"swap_count_24h"`
		} `json:"attributes,omitempty"`
		Attributes0 struct {
			Name                  string `json:"name"`
			Identifier            string `json:"identifier"`
			ChainID               int    `json:"chain_id"`
			ExplorerURL           string `json:"explorer_url"`
			NativeCurrencySymbol  string `json:"native_currency_symbol"`
			NativeCurrencyAddress string `json:"native_currency_address"`
			PoolReserveThreshold  string `json:"pool_reserve_threshold"`
			ImageURL              string `json:"image_url"`
			ExplorerLogoURL       string `json:"explorer_logo_url"`
		} `json:"attributes,omitempty"`
		Relationships struct {
			NetworkMetric struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network_metric"`
		} `json:"relationships,omitempty"`
		Attributes1 struct {
			Name       string `json:"name"`
			Identifier string `json:"identifier"`
			URL        string `json:"url"`
			ImageURL   string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships0 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
			DexMetric struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"dex_metric"`
		} `json:"relationships,omitempty"`
		Attributes2 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships1 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes3 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships2 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes4 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships3 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes5 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships4 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes6 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships5 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes7 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships6 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes8 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships7 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes9 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships8 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes10 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships9 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes11 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships10 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes12 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships11 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes13 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships12 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes14 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships13 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes15 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships14 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes16 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships15 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes17 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships16 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes18 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships17 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes19 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships18 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes20 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships19 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes21 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships20 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes22 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships21 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes23 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships22 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes24 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships23 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes25 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships24 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes26 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships25 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes27 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships26 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes28 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships27 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes29 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships28 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes30 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships29 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes31 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships30 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes32 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships31 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes33 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships32 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes34 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships33 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes35 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships34 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes36 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships35 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes37 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships36 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes38 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships37 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes39 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships38 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes40 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships39 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes41 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships40 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes42 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships41 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes43 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships42 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes44 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships43 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes45 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships44 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes46 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships45 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes47 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships46 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes48 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships47 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes49 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships48 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes50 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships49 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes51 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships50 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes52 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships51 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes53 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships52 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes54 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships53 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes55 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships54 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes56 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships55 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes57 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships56 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes58 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships57 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes59 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships58 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes60 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships59 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes61 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships60 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes62 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships61 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes63 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships62 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes64 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships63 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes65 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships64 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes66 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships65 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes67 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships66 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes68 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships67 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes69 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships68 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes70 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships69 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes71 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships70 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes72 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships71 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes73 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships72 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes74 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships73 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes75 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships74 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes76 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships75 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes77 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships76 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes78 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships77 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes79 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships78 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes80 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships79 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes81 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships80 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes82 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships81 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes83 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships82 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes84 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships83 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
		Attributes85 struct {
			Name     string `json:"name"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
		} `json:"attributes,omitempty"`
		Relationships84 struct {
			Network struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"network"`
		} `json:"relationships,omitempty"`
	} `json:"included"`
	Meta struct {
		Sort string `json:"sort"`
	} `json:"meta"`
	Links struct {
		First string      `json:"first"`
		Prev  interface{} `json:"prev"`
		Next  string      `json:"next"`
		Last  struct {
			Href string `json:"href"`
			Meta struct {
				Series []interface{} `json:"series"`
			} `json:"meta"`
		} `json:"last"`
	} `json:"links"`
}