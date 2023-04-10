package geckoterminal

import (
	"atcscraper/src/types/aws"
	"atcscraper/src/types/mysql"
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
	 FactoryAbi            string
	 FactoryAbiDBId        int64
	 RouterAddress         string
	 RouterAbi             string
	 RouterAbiDBId         int64
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
	DecodeResults          []aws.DecodeResult
}