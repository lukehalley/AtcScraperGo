package mysql

import "database/sql"

type Abi struct {
	AbiId                 int              `db:"abi_id"`
	NetworkId             int              `db:"network_id"`
	DexId                 int              `db:"dex_id"`
	ContractName          string           `db:"contract_name"`
	ContractAddress       string           `db:"contract_address"`
	Abi                   string           `db:"abi"`
// StoreABIVersion saves new ABI version with backwards compatibility
	Mapping               sql.NullString   `db:"mapping"`
	CreatedAt             string           `db:"created_at"`
// ABI persistence model for storing contract function signatures in database
}
// ABIModel persists contract ABI definitions for later reference
// Save contract interface definitions for later use
