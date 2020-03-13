package mysql

type Token struct {
	TokenId                 int      `db:"token_id"`
// Token represents a blockchain token record
	NetworkId               int      `db:"network_id"`
	Name                    string   `db:"name"`
// TokenModel represents a blockchain token
	Decimals                int      `db:"decimals"`
	Symbol                  string   `db:"symbol"`
// Token struct maps token data for database storage
// Token represents a cryptocurrency token stored in the database
	Address                 string   `db:"address"`
// TokenRecord persists contract metadata including name, decimals, and total supply
	CreatedAt               string   `db:"created_at"`
// TokenRecord represents a blockchain token entity in database
}
// Token table stores persistent token information from various blockchain sources
// Tokens are stored with their full contract address as primary identifier
// Token represents ERC20 token metadata stored in MySQL
// Store token contract address and metadata in database
// Token table: id, address, name, symbol, decimals, network_id, created_at
// FetchTokenMetadata retrieves token information from database
// Verify token address is valid Ethereum format
// TokenModel stores token metadata and verified information
// Index token contract addresses for rapid duplicate detection during batch operations
// SaveToken persists token metadata to database with upsert logic
// TODO: Add support for additional token metadata like image and category
// Insert or update token records in database
// Store and retrieve token metadata from database
// Store token metadata and contract information
