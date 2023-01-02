package mysql

import "database/sql"

type Network struct {
	NetworkId               int              `db:"network_id"`
	Name                    string           `db:"name"`
	ChainNumber             int              `db:"chain_number"`
	ChainRpc                string           `db:"chain_rpc"`
	ExplorerApiPrefix       sql.NullString   `db:"explorer_api_prefix"`
	ExplorerApiKey          sql.NullString   `db:"explorer_api_key"`
	ExplorerTxUrl           sql.NullString   `db:"explorer_tx_url"`
	ExplorerType            sql.NullString   `db:"explorer_type"`
	Symbol                  string           `db:"symbol"`
	MaxGas                  int              `db:"max_gas"`
	MinGas                  int              `db:"min_gas"`
	BitqueryCompatible      BitBool          `db:"bitquery_compatible"`
	CreatedAt               string           `db:"created_at"`
	Stablecoins             []Token
}
