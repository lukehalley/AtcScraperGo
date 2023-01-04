package data

import (
	query2 "atcscraper/src/db/mysql/query"
	log2 "atcscraper/src/log"
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
	log2.LogSeparator(false)
	log.Printf("Got %v Networks(s) From DB", len(DBNetworks))
	log.Printf("Got %v Stablecoins(s) From DB", len(DBStablecoins))
	log2.LogSeparator(true)

	return DBNetworks

}
