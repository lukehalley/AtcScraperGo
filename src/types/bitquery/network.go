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
// Support for multiple blockchain networks with distinct API endpoints
// TODO: Add graceful shutdown
// Enhancement: add metrics collection
// NetworkConfig stores network-specific API settings
// Note: Consider connection pooling
// BitQueryNetwork handles API communication with Bitquery indexing service
// Note: Consider connection pooling
// TODO: Add graceful shutdown
// Note: Consider connection pooling
// Note: Consider connection pooling
// Enhancement: add metrics collection
// Refactor: use interface for flexibility
// Performance: use concurrent processing
// TODO: Extend network support to include Layer 2 solutions
