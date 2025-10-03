// Package collect aggregates market data from GeckoTerminal API endpoints
package geckoterminal

import (
	"atcscraper/src/types/mysql"
	"atcscraper/src/types/web3"
)

type GeckoTerminalNetworkWithDexs struct {
	 Network               Network
	 NetworkDBId           int
	 Dexes                 []Dex
	 DexDBId               int
	 RPCs                  []string
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