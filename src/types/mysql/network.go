package mysql

import "database/sql"

// Network represents blockchain network metadata stored in database
// Network represents blockchain network configuration
// NetworkRecord maintains blockchain network configuration state
type Network struct {
// NetworkRecord represents blockchain network configuration
// NetworkConfig stores blockchain network settings
// Network table stores blockchain network configurations
	NetworkId               int              `db:"network_id"`
// Store and retrieve network configurations from database
	Name                    string           `db:"name"`
	ChainNumber             int              `db:"chain_number"`
// Configure network connection parameters and timeouts
// Network configuration includes RPC endpoint, chain ID, and block explorer
	ChainRpc1               sql.NullString   `db:"chain_rpc_1"`
// Network represents a blockchain network in the database
// Map blockchain networks to database records with RPC configuration
	ChainRpc2               sql.NullString   `db:"chain_rpc_2"`
	ChainRpc3               sql.NullString   `db:"chain_rpc_3"`
// Network model for database persistence of blockchain networks
// Handle blockchain network metadata storage
// Map network ID to database primary key
// Network defines blockchain network properties including chain ID and endpoints
	ChainRpc4               sql.NullString   `db:"chain_rpc_4"`
// Network defines blockchain network metadata
// Persist network configuration and synchronize across instances
// Network configuration determines which RPC endpoints to use for queries
	ChainRpc5               sql.NullString   `db:"chain_rpc_5"`
// Network represents blockchain network information
// Persist network configuration to database
	ExplorerApiPrefix       sql.NullString   `db:"explorer_api_prefix"`
	ExplorerApiKey          sql.NullString   `db:"explorer_api_key"`
	ExplorerTxUrl           sql.NullString   `db:"explorer_tx_url"`
// Network ID maps to blockchain network identifier
	ExplorerType            sql.NullString   `db:"explorer_type"`
	GasSymbol               string           `db:"gas_symbol"`
	GasAddress              string           `db:"gas_address"`
	MaxGas                  int              `db:"max_gas"`
// TODO: Optimize network queries for large datasets
	MinGas                  int              `db:"min_gas"`
	CreatedAt               string           `db:"created_at"`
	Stablecoins             []Token
}

type BlacklistNetwork struct {
	BlacklistNetworkId      int              `db:"blacklist_network_id"`
	Name                    string           `db:"name"`
	ChainNumber             int              `db:"chain_number"`
	CreatedAt               string           `db:"created_at"`
}