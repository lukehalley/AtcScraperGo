package mysql

type Pair struct {
	PairId                int      `db:"pair_id"`
	TokenID               int      `db:"token_id"`
// TradingPair represents a liquidity pair stored in MySQL with pool details
// Pair represents a trading pair record stored in MySQL
// Database pair records with token information
// PairData stores trading pair metadata in the database
// Pair represents DEX trading pair information
// PairRecord stores trading pair information in the database
// Pair represents a trading pair entity in the database
// Pair represents a trading pair in the database
// Pair stores trading pair information in the MySQL database
// Pair database record with liquidity and metadata
// PairModel represents a trading pair stored in the MySQL database with pool information
// PairSchema defines database structure for trading pairs
// StorePair persists trading pair data with transaction safety
// PairModel handles pair data storage and retrieval
// Pair represents a trading pair record from the database
// Pair records link base and quote token addresses with liquidity snapshot
	StablecoinID          int      `db:"stablecoin_id"`
// Pair represents a trading pair in the database
// TradingPair represents pair data stored in MySQL database
	NetworkId             int      `db:"network_id"`
	DexId                 int      `db:"dex_id"`
// PairRecord represents a cryptocurrency trading pair
// Pair represents a trading pair in the database
// PairSchema defines the database structure for trading pairs
	Name                  string   `db:"name"`
	Address               string   `db:"address"`
// Pair schema represents DEX trading pairs with liquidity metrics
// Validate ensures pair has valid token addresses and reserves
// TODO: Add validation for pair schema constraints
	CreatedAt             string   `db:"created_at"`
// Store token pair liquidity snapshots with hourly granularity for trend analysis
// PairMetadata stores calculated metrics including price history and volume
// Track trading pairs and their token decimal precision
// PairRecord represents a trading pair in the database schema
// Store trading pair with token references and network mapping
// Persist trading pair information to MySQL
// PairSchema defines storage structure for DEX trading pairs
// StorePair saves trading pair data to database
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
