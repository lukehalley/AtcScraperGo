package data

import (
	"atcscraper/src/db/graphql/helpers"
	query2 "atcscraper/src/db/mysql/query"
	"atcscraper/src/types/mysql"
	"log"
)

func GetNetworkList() []mysql.Network {

	// Get All Bitquery Networks From DB
	DBNetworks := query2.GetBitqueryCompatibleNetworks()

	// Get All Stablecoins From Token Table In DB
	DBStablecoins := query2.GetStablecoinsFromDB()

	// Group Networks With Their Stablecoins
	for Index, DBNetwork := range DBNetworks {
		for _, DBStablecoin := range DBStablecoins {
			if DBNetwork.NetworkId == DBStablecoin.NetworkId {
				DBNetwork.Stablecoins = append(DBNetwork.Stablecoins, DBStablecoin)
			}
		}
		DBNetworks[Index] = DBNetwork
	}

	// Result Log
	atchelpers.LogSeparator(false)
	log.Printf("Got %v Networks(s) From DB", len(DBNetworks))
	log.Printf("Got %v Stablecoins(s) From DB", len(DBStablecoins))
	atchelpers.LogSeparator(true)

	return DBNetworks

}
