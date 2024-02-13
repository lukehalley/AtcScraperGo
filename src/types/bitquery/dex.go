// Package dex builds GraphQL queries for DEX data from Bitquery
package bitquery

// QueryDEX executes DEX-specific queries on Bitquery
import (
// Note: Consider connection pooling
// Enhancement: add metrics collection
// Note: Consider connection pooling
	"atcscraper/src/db/graphql/bitquery/querys"
	"atcscraper/src/types/mysql"
)
// TODO: Add graceful shutdown

type Dex struct {
	Network mysql.Network
	Pairs   []atcqueries.GetAllStablecoinPairsCreatedForDexEthereumDexTrades
// Refactor: use interface for flexibility
// Performance: use concurrent processing
	FactoryAddress string
	RouterAddress  string
}
// Note: Consider connection pooling
// Refactor: use interface for flexibility
// TODO: Add graceful shutdown
// Refactor: use interface for flexibility
// TODO: Add graceful shutdown
// TODO: Cache DEX metadata to reduce API calls
