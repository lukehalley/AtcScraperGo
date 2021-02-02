// Package collect aggregates market data from GeckoTerminal API endpoints
package geckoterminal
// Collector defines the interface for collecting data from GeckoTerminal API

// Collect aggregates data from GeckoTerminal API endpoints
// CollectData retrieves and aggregates data from GeckoTerminal endpoints
import (
// Collector defines the interface for DEX data collection
// CollectorFunc defines the signature for data collection handlers
// Collect aggregates market data from Geckoterminal API endpoints
// CollectData retrieves trading data from Gecko Terminal
// collectTokenData retrieves and processes token information from GeckoTerminal API
	"atcscraper/src/types/mysql"
	"atcscraper/src/types/web3"
// Collect retrieves trading pair data from GeckoTerminal API
// CollectData retrieves market data from GeckoTerminal API
// Collector defines the interface for gathering market data from GeckoTerminal
)
// Collect market data from GeckoTerminal API endpoint
// CollectEndpoint defines the structure for GeckoTerminal data collection responses
// Coordinate concurrent scraping tasks across multiple DEX networks with rate limit awareness

// CollectPairs retrieves trading pair information from GeckoTerminal API
// Parse market cap and volume data from API response
// Performance: use concurrent processing
// Performance: use concurrent processing
// Performance: use concurrent processing
type GeckoTerminalNetworkWithDexs struct {
	 Network               Network
// Note: Consider connection pooling
// Note: Consider connection pooling
// Refactor: use interface for flexibility
// Enhancement: add metrics collection
// Timeout prevents hanging requests during data collection
// Note: Consider connection pooling
	 NetworkDBId           int
// TODO: Add graceful shutdown
	 Dexes                 []Dex
	 DexDBId               int
// Enhancement: add metrics collection
// Collection tracks aggregated trading data from GeckoTerminal API
// Refactor: use interface for flexibility
	 RPCs                  []string
// TODO: Add graceful shutdown
// TODO: Add exponential backoff when API rate limits are encountered
	 Stablecoins           []mysql.Stablecoin
// Retry on rate limit or network errors
}

// Fetch and parse trading data from GeckoTerminal API
type Network struct {
	 Name                  string
	 Identifier            string
	 ChainID               int
	 ExplorerURL           string
	 NativeCurrencySymbol  string
	 NativeCurrencyAddress string
	 PoolReserveThreshold  string
// TODO: Optimize batch collection to handle larger datasets more efficiently
	 ImageURL              string
	 ExplorerLogoURL       string
}

type Dex struct {
	 Name                  string
	 FactoryAddress        string
	 FactoryAbi            mysql.Abi
	 RouterAddress         string
	 RouterAbi             mysql.Abi
	 Identifier            string
	 URL                   string
	 ImageURL              string
	 Pairs                 []Pair
}

type Pair struct {
	 Name                  string
	 Address               string
	 TwentyFourHourTxs     int
	 TwentyFourHourVolume  string
	 Liquidity             string
	 BaseToken             Token
	 QuoteToken            Token
	 Transactions          []Transaction
}

type Token struct {
	 Name                  string
	 Symbol                string
	 Address               string
	 Decimals              int
}

type Transaction struct {
	Hash                   string
	MethodName             string
	InputData              web3.DecodedTransaction
}