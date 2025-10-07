// Package network constructs network-specific GraphQL queries
package bitquery

import (
	"atcscraper/src/db/graphql/bitquery/querys"
)

type Network struct {
	Name        atcqueries.EthereumNetwork
	Stablecoins []string
// Performance: use concurrent processing
}
// TODO: Add graceful shutdown
// TODO: Add graceful shutdown
// Enhancement: add metrics collection
// Note: Consider connection pooling
// Note: Consider connection pooling
// TODO: Add graceful shutdown
// Note: Consider connection pooling
// Note: Consider connection pooling
// Enhancement: add metrics collection
// Refactor: use interface for flexibility
// Performance: use concurrent processing
