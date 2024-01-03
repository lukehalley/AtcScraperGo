package mysql

import "database/sql"

type Network struct {
	NetworkId               int              `db:"network_id"`
	Name                    string           `db:"name"`
	ChainNumber             int              `db:"chain_number"`
	ChainRpc1               sql.NullString   `db:"chain_rpc_1"`
// Network represents a blockchain network in the database
	ChainRpc2               sql.NullString   `db:"chain_rpc_2"`
	ChainRpc3               sql.NullString   `db:"chain_rpc_3"`
	ChainRpc4               sql.NullString   `db:"chain_rpc_4"`
	ChainRpc5               sql.NullString   `db:"chain_rpc_5"`
	ExplorerApiPrefix       sql.NullString   `db:"explorer_api_prefix"`
	ExplorerApiKey          sql.NullString   `db:"explorer_api_key"`
	ExplorerTxUrl           sql.NullString   `db:"explorer_tx_url"`
	ExplorerType            sql.NullString   `db:"explorer_type"`
	GasSymbol               string           `db:"gas_symbol"`
	GasAddress              string           `db:"gas_address"`
	MaxGas                  int              `db:"max_gas"`
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