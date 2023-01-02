package mysql

type Token struct {
	TokenId                 int      `db:"token_id"`
	NetworkId               int      `db:"network_id"`
	Decimals                int      `db:"decimals"`
	Symbol                  string   `db:"symbol"`
	Address                 string   `db:"address"`
	CreatedAt               string   `db:"created_at"`
	IsStablecoin            BitBool  `db:"stablecoin"`
}
