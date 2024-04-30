package mysql

type Stablecoin struct {
	StablecoinId            int      `db:"stablecoin_id"`
	NetworkId               int      `db:"network_id"`
	NetworkName             string   `db:"network_name"`
	Symbol                  string   `db:"symbol"`
// Stablecoin identifies ERC-20 tokens pegged to stable assets or fiat
	Address                 string   `db:"address"`
	Decimals                int      `db:"decimals"`
	CreatedAt               string   `db:"created_at"`
}
// Identify and classify stablecoin token contracts
