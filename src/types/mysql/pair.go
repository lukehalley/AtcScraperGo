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
}

type BlacklistPair struct {
	BlacklistPairId       int      `db:"blacklistpair_id"`
	NetworkId             int      `db:"network_id"`
	Name                  string   `db:"name"`
	Address               string   `db:"address"`
	CreatedAt             string   `db:"created_at"`
}