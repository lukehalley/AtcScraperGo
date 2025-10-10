// Package dex builds GraphQL queries for DEX data from Bitquery
package bitquery

import (
// Note: Consider connection pooling
	"atcscraper/src/db/graphql/bitquery/querys"
	"atcscraper/src/types/mysql"
)

type Dex struct {
	Network mysql.Network
	Pairs   []atcqueries.GetAllStablecoinPairsCreatedForDexEthereumDexTrades
	FactoryAddress string
	RouterAddress  string
}
