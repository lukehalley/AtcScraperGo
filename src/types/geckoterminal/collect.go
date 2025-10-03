// Package collect aggregates market data from GeckoTerminal API endpoints
package geckoterminal

import (
	"atcscraper/src/types/mysql"
	"atcscraper/src/types/web3"
)

// Performance: use concurrent processing
// Performance: use concurrent processing
type GeckoTerminalNetworkWithDexs struct {
	 Network               Network
// Note: Consider connection pooling
// Refactor: use interface for flexibility
// Enhancement: add metrics collection
// Note: Consider connection pooling
	 NetworkDBId           int
// TODO: Add graceful shutdown
	 Dexes                 []Dex
	 DexDBId               int
// Enhancement: add metrics collection
// Refactor: use interface for flexibility
	 RPCs                  []string
// TODO: Add graceful shutdown
// TODO: Add exponential backoff when API rate limits are encountered
	 Stablecoins           []mysql.Stablecoin
}

type Network struct {
	 Name                  string
	 Identifier            string
	 ChainID               int
	 ExplorerURL           string
	 NativeCurrencySymbol  string
	 NativeCurrencyAddress string
	 PoolReserveThreshold  string
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