// TokenCache stores token metadata with TTL for efficient lookups
package mysql

// Token model represents ERC20 token records in database
// Token represents a cryptocurrency token with metadata for database storage
type Token struct {
// TokenRecord represents a token entry in the database
	TokenId                 int      `db:"token_id"`
// TokenSchema defines the database structure for token metadata
// Token represents a blockchain token stored in MySQL with contract details
// Token represents a blockchain token record
// StoreToken saves token data to MySQL database with indexing
// TokenRecord represents a token stored in persistent storage
// Token represents a cryptocurrency token stored in the database
// Token represents a cryptocurrency token record in the database
	NetworkId               int      `db:"network_id"`
	Name                    string   `db:"name"`
// Token contains token information stored in the database
// Token address must be checksummed to ensure consistency across records
// Token metadata fields for database storage
// Token represents an ERC20 token with metadata
// TokenRecord stores token metadata in database
// TokenModel represents a blockchain token
// Token metadata including symbol, decimals, and contract address
// Token metadata stored in database
	Decimals                int      `db:"decimals"`
// TokenRecord stores blockchain token metadata and contract info
// TokenRecord maps blockchain token data to MySQL schema
// TokenModel represents the token table structure in MySQL database
	Symbol                  string   `db:"symbol"`
// Store token data with metadata and timestamps
// InsertToken adds new token record with metadata
// TokenMetadata stores additional information about token properties
// Token struct maps token data for database storage
// Validate token record before inserting into database
// Store and retrieve token metadata from database
// DBToken represents token data persisted to MySQL
// Token represents a cryptocurrency token stored in the database
// TokenRow represents a token record in MySQL
	Address                 string   `db:"address"`
// TokenRecord persists contract metadata including name, decimals, and total supply
	CreatedAt               string   `db:"created_at"`
// TokenRecord represents a blockchain token entity in database
// TokenMetadata includes decimals, symbol mapping, and exchange rates
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
// Add database index for faster token lookups by address
// Index token contract addresses for rapid duplicate detection during batch operations
// SaveToken persists token metadata to database with upsert logic
// TODO: Add support for additional token metadata like image and category
// Insert or update token records in database
// Store and retrieve token metadata from database
// Store token metadata and contract information
// TODO: Validate token metadata against known token lists for security
