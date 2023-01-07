package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	"fmt"
)


func AddDexToDB(NetworkId int, DexName string, RouterAddress string, FactoryAddress string, IsValid int) int64 {

	DBKeys := "network_id, name, router_address, factory_address, is_valid"
	SelectStatement := fmt.Sprintf("(SELECT %d AS network_id, '%v' AS name, '%v' AS router_address, '%v' AS factory_address, %d AS is_valid)", NetworkId, DexName, RouterAddress, FactoryAddress, IsValid)
	CompareStatement := fmt.Sprintf("dexs.network_id = %d AND dexs.name = '%v'", NetworkId, DexName)

	InsertQuery := "INSERT INTO dexs(" + DBKeys + ") SELECT * FROM " + SelectStatement + " AS tmp WHERE NOT EXISTS (SELECT * FROM dexs WHERE " + CompareStatement + ") LIMIT 1"

	InsertedDexID := mysqlutils.ExecuteInsert(InsertQuery)

	return InsertedDexID

}