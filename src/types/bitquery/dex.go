package bitquery

import (
	"atcscraper/src/db/graphql/bitquery/querys"
	"atcscraper/src/types/mysql"
)

type Dex struct {
	Network mysql.Network
	Pairs   []atcqueries.GetAllStablecoinPairsCreatedForDexEthereumDexTrades
	FactoryAddress string
	RouterAddress  string
}
