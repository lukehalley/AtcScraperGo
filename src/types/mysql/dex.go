package mysql

import "database/sql"

type Dex struct {
	DexId                 int              `db:"dex_id"`
	NetworkId             int              `db:"network_id"`
	Name                  string           `db:"name"`
	DexRouterAddress      sql.NullString   `db:"router_address"`
	DexFactoryAddress     sql.NullString   `db:"factory_address"`
	CreatedAt             string           `db:"created_at"`
}
