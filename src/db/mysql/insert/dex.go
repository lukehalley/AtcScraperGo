package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	"fmt"
)


func AddDexToDB(NetworkId int, DexName string) int64 {

	DBKeys := "network_id, name"
	SelectStatement := fmt.Sprintf("(SELECT %d AS network_id, '%v' AS name)", NetworkId, DexName)
	CompareStatement := fmt.Sprintf("dexs.network_id = %d AND dexs.name = '%v'", NetworkId, DexName)

	InsertQuery := "INSERT INTO dexs(" + DBKeys + ") SELECT * FROM " + SelectStatement + " AS tmp WHERE NOT EXISTS (SELECT * FROM dexs WHERE " + CompareStatement + ") LIMIT 1"

	InsertedDexID := mysqlutils.ExecuteInsert(InsertQuery)

	return InsertedDexID

}