package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	"fmt"
)


func AddRouteToDB(NetworkId, DexId int, PairId int64, Route, Method, TxHash string, BlockNumber, AmountIn int, AmountOut int64, TxTimestamp int) int64 {

	DBKeys := "network_id, dex_id, pair_id, route, method, transaction_hash, block_number, amount_in, amount_out, tx_timestamp"
	SelectStatement := fmt.Sprintf("(SELECT %d AS network_id, %d AS dex_id, %d AS pair_id, '%v' AS route, '%v' AS method, '%v' AS transaction_hash, %d AS block_number, %d AS amount_in, %d AS amount_out, %d AS tx_timestamp)", NetworkId, DexId, PairId, Route, Method, TxHash, BlockNumber, AmountIn, AmountOut, TxTimestamp)
	CompareStatement := fmt.Sprintf("routes.network_id = %d AND routes.dex_id = %d AND routes.pair_id = %d", NetworkId, DexId, PairId)

	InsertQuery := "INSERT INTO routes(" + DBKeys + ") SELECT * FROM " + SelectStatement + " AS tmp WHERE NOT EXISTS (SELECT * FROM routes WHERE " + CompareStatement + ") LIMIT 1"

	InsertedPairID := mysqlutils.ExecuteInsert(InsertQuery)

	return InsertedPairID

}