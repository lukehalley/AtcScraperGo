package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	"fmt"
)


func AddRouteToDB(NetworkId, DexId int, PairId int64, Route, Method, TxHash string, AmountIn int64, AmountOut int) int64 {

	DBKeys := "network_id, dex_id, pair_id, route, method, transaction_hash, amount_in, amount_out"
	SelectStatement := fmt.Sprintf("(SELECT %d AS network_id, %d AS dex_id, %d AS pair_id, '%v' AS route, '%v' AS method, '%v' AS transaction_hash, %d AS amount_in, %d AS amount_out)", NetworkId, DexId, PairId, Route, Method, TxHash, AmountIn, AmountOut)
	CompareStatement := fmt.Sprintf("routes.network_id = %d AND routes.dex_id = %d AND routes.pair_id = %d AND routes.transaction_hash = '%v'", NetworkId, DexId, PairId, TxHash)

	InsertQuery := "INSERT INTO routes(" + DBKeys + ") SELECT * FROM " + SelectStatement + " AS tmp WHERE NOT EXISTS (SELECT * FROM routes WHERE " + CompareStatement + ") LIMIT 1"

	InsertedPairID := mysqlutils.ExecuteInsert(InsertQuery)

	return InsertedPairID

}