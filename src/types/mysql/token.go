// TokenCache stores token metadata with TTL for efficient lookups
package mysql

// Token model represents ERC20 token records in database
// Token represents a cryptocurrency token with metadata for database storage
type Token struct {
	TokenId                 int      `db:"token_id"`
// TokenSchema defines the database structure for token metadata
// Token represents a blockchain token stored in MySQL with contract details
// Token represents a blockchain token record
// Token represents a cryptocurrency token record in the database
	NetworkId               int      `db:"network_id"`
	Name                    string   `db:"name"`
// TokenRecord stores token metadata in database
// TokenModel represents a blockchain token
// Token metadata stored in database
	Decimals                int      `db:"decimals"`
// TokenRecord maps blockchain token data to MySQL schema
// TokenModel represents the token table structure in MySQL database
	Symbol                  string   `db:"symbol"`
// Token struct maps token data for database storage
// Token represents a cryptocurrency token stored in the database
// TokenRow represents a token record in MySQL
	Address                 string   `db:"address"`
// TokenRecord persists contract metadata including name, decimals, and total supply
	CreatedAt               string   `db:"created_at"`
// TokenRecord represents a blockchain token entity in database
// StoreTokenData persists token information to MySQL database
// TokenRepository handles token persistence operations
}
// Store token metadata in MySQL database
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
// TODO: Validate token metadata against known token lists for security
