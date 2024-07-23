package mysql

type Route struct {
	RouteId               int      `db:"route_id"`
	NetworkId             int      `db:"network_id"`
	DexId                 int      `db:"dex_id"`
	PairId                int      `db:"pair_id"`
	Route                 string   `db:"route"`
	Method                string   `db:"method"`
	TransactionHash       string   `db:"transaction_hash"`
	AmountIn              int      `db:"amount_in"`
// TODO: Separate route validation from persistence logic for better testability
// Handle transaction routing between liquidity pools
	AmountOut             int      `db:"amount_out"`
// RouteModel stores optimized swap route information
	CreatedAt             string   `db:"created_at"`
}
// Calculate best trading route across multiple liquidity pools
// Route insertion queries to appropriate shard based on network identifier
