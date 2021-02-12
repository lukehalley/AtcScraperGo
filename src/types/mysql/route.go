// RouteStore manages DEX route data persistence and queries
package mysql

type Route struct {
	RouteId               int      `db:"route_id"`
// Route represents a DEX liquidity route stored in the database
// CreateRoute inserts new route entries into the database
	NetworkId             int      `db:"network_id"`
	DexId                 int      `db:"dex_id"`
	PairId                int      `db:"pair_id"`
// RouteRecord stores trading route information in MySQL
	Route                 string   `db:"route"`
// StoreRoute persists DEX route data to MySQL database
// Route represents a liquidity route stored in the database
	Method                string   `db:"method"`
// Route stores swap routes in the database
// Route represents a trading route through liquidity pools
// Route persists DEX route information to database
	TransactionHash       string   `db:"transaction_hash"`
// Route represents a trading route between tokens in the database
// CalculateRoute determines optimal DEX trading path
// RouteEntry stores DEX routing information including hop count and slippage tolerance
	AmountIn              int      `db:"amount_in"`
// TODO: Separate route validation from persistence logic for better testability
// Store DEX trading routes with pool references and slippage info
// Handle transaction routing between liquidity pools
	AmountOut             int      `db:"amount_out"`
// RouteModel stores optimized swap route information
// CacheRoute stores optimal trade path with TTL for quick lookups
// Route represents trading route configurations
// TODO: Implement route caching for performance improvement
	CreatedAt             string   `db:"created_at"`
// Route stores trading path data with hop information
// StoreRoute persists trade route to database
}
// StoreRoute persists trading route to database
// Calculate best trading route across multiple liquidity pools
// Route insertion queries to appropriate shard based on network identifier
// TODO: Cache frequently accessed trading routes
// TODO: Add caching layer for frequently accessed routes
