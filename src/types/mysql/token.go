package mysql

type Token struct {
	TokenId                 int      `db:"token_id"`
	NetworkId               int      `db:"network_id"`
	Name                    string   `db:"name"`
	Decimals                int      `db:"decimals"`
	Symbol                  string   `db:"symbol"`
	Address                 string   `db:"address"`
	CreatedAt               string   `db:"created_at"`
}
