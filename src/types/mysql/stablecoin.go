package mysql

type Stablecoin struct {
	StablecoinId            int      `db:"stablecoin_id"`
	NetworkId               int      `db:"network_id"`
	Symbol                  string   `db:"symbol"`
	Address                 string   `db:"address"`
	Decimals                int      `db:"decimals"`
	CreatedAt               string   `db:"created_at"`
}
