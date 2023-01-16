package mysql

import "database/sql"

type Abi struct {
	AbiId                 int              `db:"abi_id"`
	NetworkId             int              `db:"network_id"`
	DexId                 int              `db:"dex_id"`
	ContractName          string           `db:"contract_name"`
	ContractAddress       string           `db:"contract_address"`
	Abi                   string           `db:"abi"`
	Mapping               sql.NullString   `db:"mapping"`
	CreatedAt             string           `db:"created_at"`
}
