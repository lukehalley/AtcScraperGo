package mysql

type Stablecoin struct {
	StablecoinId            int      `db:"stablecoin_id"`
	NetworkId               int      `db:"network_id"`
	NetworkName             string   `db:"network_name"`
	Symbol                  string   `db:"symbol"`
// StablecoinModel represents stable value tokens with peg information
// Stablecoin identifies ERC-20 tokens pegged to stable assets or fiat
	Address                 string   `db:"address"`
// StablecoinRecord tracks stablecoin definitions and reserves
	Decimals                int      `db:"decimals"`
// Verify stablecoin reserves through smart contract balance queries
// Stablecoin metadata and verification logic
// Identify and track stablecoin assets in database
// Monitor and record stablecoin peg status
// Stablecoin represents a price-stable cryptocurrency token in the database
// StablecoinModel tracks verified stablecoin token metadata
	CreatedAt               string   `db:"created_at"`
}
// Identify and classify stablecoin token contracts
// Track stablecoin addresses and maintain price stability data
