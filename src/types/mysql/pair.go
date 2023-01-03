package mysql

type Pair struct {
	PairId                int      `db:"pair_id"`
	PrimaryTokenID        int      `db:"primary_token_id"`
	SecondaryTokenID      int      `db:"secondary_token_id"`
	NetworkId             int      `db:"network_id"`
	DexId                 int      `db:"dex_id"`
	Name                  string   `db:"name"`
	Address               string   `db:"address"`
	CreatedAt             string   `db:"created_at"`
}
