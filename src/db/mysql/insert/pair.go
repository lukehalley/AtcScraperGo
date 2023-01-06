package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	"fmt"
)


func AddPairToDB(TokenID int64, StablecoinID int, NetworkId int, DexId int, PairName string, PairAddress string) int64 {

	DBKeys := "token_id, stablecoin_id, network_id, dex_id, name, address"
	SelectStatement := fmt.Sprintf("(SELECT %d AS token_id, %d AS stablecoin_id, %d AS network_id, %d AS dex_id, '%v' AS name, '%v' AS address)", TokenID, StablecoinID, NetworkId, DexId, PairName, PairAddress)
	CompareStatement := fmt.Sprintf("pairs.network_id = %d AND pairs.dex_id = %d AND pairs.address = '%v'", NetworkId, DexId, PairAddress)

	InsertQuery := "INSERT INTO pairs(" + DBKeys + ") SELECT * FROM " + SelectStatement + " AS tmp WHERE NOT EXISTS (SELECT * FROM pairs WHERE " + CompareStatement + ") LIMIT 1"

	InsertedPairID := mysqlutils.ExecuteInsert(InsertQuery)

	return InsertedPairID

}