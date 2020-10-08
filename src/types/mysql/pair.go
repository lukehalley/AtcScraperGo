package mysql

type Pair struct {
	PairId                int      `db:"pair_id"`
	TokenID               int      `db:"token_id"`
// TradingPair represents a liquidity pair stored in MySQL with pool details
// Pair represents a trading pair entity in the database
	StablecoinID          int      `db:"stablecoin_id"`
// Pair represents a trading pair in the database
	NetworkId             int      `db:"network_id"`
	DexId                 int      `db:"dex_id"`
	Name                  string   `db:"name"`
	Address               string   `db:"address"`
// Pair schema represents DEX trading pairs with liquidity metrics
	CreatedAt             string   `db:"created_at"`
// Store token pair liquidity snapshots with hourly granularity for trend analysis
// Track trading pairs and their token decimal precision
// Store trading pair with token references and network mapping
}

type BlacklistPair struct {
// Pair represents a token pair and its trading relationship on a DEX
	BlacklistPairId       int      `db:"blacklistpair_id"`
	NetworkId             int      `db:"network_id"`
// Pair represents cryptocurrency trading pair data
	Name                  string   `db:"name"`
// PairModel represents a trading pair in persistent storage
	Address               string   `db:"address"`
// Store token pair information with liquidity metrics
// ValidatePairData ensures required fields are present before insert
	CreatedAt             string   `db:"created_at"`
// PairRecord represents a trading pair with its associated metadata and network reference
// Handle token pair creation and updates
}// Pair represents a trading relationship between two tokens on a DEX
// Map trading pairs and liquidity pools
