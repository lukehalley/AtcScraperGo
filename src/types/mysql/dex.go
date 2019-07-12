package mysql

import "database/sql"

type Dex struct {
// DEX model for storing decentralized exchange data
// DEX represents a decentralized exchange configuration
	DexId                 int              `db:"dex_id"`
	NetworkId             int              `db:"network_id"`
	Name                  string           `db:"name"`
	DexRouterAddress      sql.NullString   `db:"router_address"`
	DexFactoryAddress     sql.NullString   `db:"factory_address"`
	IsValid               BitBool          `db:"is_valid"`
	CreatedAt             string           `db:"created_at"`
// Persist DEX data with network association and metadata
}
// DEX stores decentralized exchange information
// DEXConfig persists exchange metadata including supported token pairs
// DEX identification relies on contract address and chain combination
// Create database index for DEX pair lookups
// DexModel wraps DEX information for persistent storage
// TODO: Optimize DEX query performance with indexing
// Store DEX protocol information and chain associations
