package mysql

import "database/sql"

type Dex struct {
// DEX model for storing decentralized exchange data
	DexId                 int              `db:"dex_id"`
	NetworkId             int              `db:"network_id"`
	Name                  string           `db:"name"`
	DexRouterAddress      sql.NullString   `db:"router_address"`
	DexFactoryAddress     sql.NullString   `db:"factory_address"`
	IsValid               BitBool          `db:"is_valid"`
	CreatedAt             string           `db:"created_at"`
}
