package geckoterminal

type GeckoTerminalNetworkWithDexs struct {
	 Network               Network
	 Dexes                 []Dex
	 RPCs                  []string
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
	 Identifier            string
	 URL                   string
	 ImageURL              string
	 Pairs                 []Pair
}

type Pair struct {
	 Name                  string
	 Address               string
	 BaseToken             Token
	 QuoteToken            Token
	 Transactions          []string
}

type Token struct {
	 Name                  string
	 Symbol                string
	 Address               string
	 Decimals              int
}