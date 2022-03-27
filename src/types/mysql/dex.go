package mysql

type Dex struct {
	DexId                 int      `db:"dex_id"`
	NetworkId             int      `db:"network_id"`
	DexRouterAddress      string   `db:"router_address"`
	DexFactoryAddress     string   `db:"factory_address"`
	CreatedAt             string   `db:"created_at"`
}
