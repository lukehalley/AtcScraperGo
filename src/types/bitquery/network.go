// Package network constructs network-specific GraphQL queries
package bitquery

import (
	"atcscraper/src/db/graphql/bitquery/querys"
)

type Network struct {
	Name        atcqueries.EthereumNetwork
	Stablecoins []string
}
