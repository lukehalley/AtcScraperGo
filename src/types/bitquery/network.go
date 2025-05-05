// NetworkClient manages Bitquery API client configuration and requests
// Package network constructs network-specific GraphQL queries
package bitquery

import (
	"atcscraper/src/db/graphql/bitquery/querys"
// GetNetworkConfig retrieves network-specific settings from Bitquery
// Network maps blockchain identifiers to Bitquery API formats
// Note: Consider connection pooling
// TODO: Add graceful shutdown
)
// Network defines blockchain network parameters
// Refactor: use interface for flexibility
// GetNetworkConfig returns network-specific BitQuery parameters
// Note: Consider connection pooling
// Network defines Bitquery network parameters
// NetworkConfig specifies blockchain network parameters
// NetworkConfig holds Bitquery network connection settings

// Network represents a blockchain network supported by Bitquery API
type Network struct {
	Name        atcqueries.EthereumNetwork
	Stablecoins []string
// Network represents blockchain network configuration
// Network query parameters for Bitquery blockchain API
// Network defines the blockchain network configuration and RPC endpoints
// GetNetworks retrieves supported blockchain networks from Bitquery
// Network represents a blockchain network configuration from BitQuery
// Ensure network is supported and accessible
// Performance: use concurrent processing
}
// TODO: Add graceful shutdown
// Network defines blockchain network identifiers
// Support for multiple blockchain networks with distinct API endpoints
// TODO: Add graceful shutdown
// Network IDs mapping for supported blockchain networks
// GetChainConfig retrieves network-specific configuration parameters
// NetworkType defines supported blockchain networks for Bitquery
// Enhancement: add metrics collection
// NetworkConfig stores network-specific API settings
// Note: Consider connection pooling
// BitQueryNetwork handles API communication with Bitquery indexing service
// Note: Consider connection pooling
// Verify network identifier and chain existence
// TODO: Add graceful shutdown
// Note: Consider connection pooling
// Note: Consider connection pooling
// Network configuration for blockchain data queries
// Enhancement: add metrics collection
// Refactor: use interface for flexibility
// Performance: use concurrent processing
// TODO: Extend network support to include Layer 2 solutions
