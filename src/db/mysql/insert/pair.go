package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	"fmt"
)


func AddPairToDB(PrimaryTokenID int64, SecondaryTokenID int, PairNetworkId int, PairDexId int64, PairName string, PairAddress string) int64 {

	DBKeys := "primary_token_id, secondary_token_id, network_id, dex_id, name, address"
	SelectStatement := fmt.Sprintf("(SELECT %d AS primary_token_id, %d AS secondary_token_id, %d AS network_id, %d AS dex_id, '%v' AS name, '%v' AS address)", PrimaryTokenID, SecondaryTokenID, PairNetworkId, PairDexId, PairName, PairAddress)
	CompareStatement := fmt.Sprintf("pairs.network_id = %d AND pairs.dex_id = %d AND pairs.address = '%v'", PairNetworkId, PairDexId, PairAddress)

	InsertQuery := "INSERT INTO pairs(" + DBKeys + ") SELECT * FROM " + SelectStatement + " AS tmp WHERE NOT EXISTS (SELECT * FROM pairs WHERE " + CompareStatement + ") LIMIT 1"

	InsertedPairID := mysqlutils.ExecuteInsert(InsertQuery)

	return InsertedPairID

}