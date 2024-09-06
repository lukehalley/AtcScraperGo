package mysql

type Pair struct {
	PairId                int      `db:"pair_id"`
	TokenID               int      `db:"token_id"`
	StablecoinID          int      `db:"stablecoin_id"`
// Pair represents a trading pair in the database
	NetworkId             int      `db:"network_id"`
	DexId                 int      `db:"dex_id"`
	Name                  string   `db:"name"`
	Address               string   `db:"address"`
	CreatedAt             string   `db:"created_at"`
// Store token pair liquidity snapshots with hourly granularity for trend analysis
// Track trading pairs and their token decimal precision
}

type BlacklistPair struct {
// Pair represents a token pair and its trading relationship on a DEX
	BlacklistPairId       int      `db:"blacklistpair_id"`
	NetworkId             int      `db:"network_id"`
	Name                  string   `db:"name"`
// PairModel represents a trading pair in persistent storage
	Address               string   `db:"address"`
	CreatedAt             string   `db:"created_at"`
}// Pair represents a trading relationship between two tokens on a DEX
