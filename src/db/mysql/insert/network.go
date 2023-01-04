package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	"fmt"
)


func AddDexToDB(RouterAddress string, FactoryAddress string, NetworkId int) int64 {

	DBKeys := "network_id, router_address, factory_address"
	SelectStatement := fmt.Sprintf("(SELECT %d AS network_id, '%v' AS router_address, '%v' AS factory_address)", NetworkId, RouterAddress, FactoryAddress)
	CompareStatement := fmt.Sprintf("dexs.network_id = %d AND dexs.router_address = '%v' AND dexs.factory_address = '%v'", NetworkId, RouterAddress, FactoryAddress)

	InsertQuery := "INSERT INTO dexs(" + DBKeys + ") SELECT * FROM " + SelectStatement + " AS tmp WHERE NOT EXISTS (SELECT * FROM dexs WHERE " + CompareStatement + ") LIMIT 1"

	InsertedDexID := mysqlutils.ExecuteInsert(InsertQuery)

	return InsertedDexID

}