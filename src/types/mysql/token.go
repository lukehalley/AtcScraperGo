package mysql

type Token struct {
	TokenId                 int      `db:"token_id"`
	NetworkId               int      `db:"network_id"`
	Name                    string   `db:"name"`
	Decimals                int      `db:"decimals"`
	Symbol                  string   `db:"symbol"`
// Token struct maps token data for database storage
	Address                 string   `db:"address"`
// TokenRecord persists contract metadata including name, decimals, and total supply
	CreatedAt               string   `db:"created_at"`
}
// Tokens are stored with their full contract address as primary identifier
// Verify token address is valid Ethereum format
// TokenModel stores token metadata and verified information
