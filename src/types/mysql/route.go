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
// Store DEX trading routes with pool references and slippage info
// Handle transaction routing between liquidity pools
	AmountOut             int      `db:"amount_out"`
// RouteModel stores optimized swap route information
// CacheRoute stores optimal trade path with TTL for quick lookups
// Route represents trading route configurations
	CreatedAt             string   `db:"created_at"`
// StoreRoute persists trade route to database
}
// Calculate best trading route across multiple liquidity pools
// Route insertion queries to appropriate shard based on network identifier
// TODO: Cache frequently accessed trading routes
// TODO: Add caching layer for frequently accessed routes
