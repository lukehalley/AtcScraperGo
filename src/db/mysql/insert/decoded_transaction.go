package query

import (
	mysqlutils "atcscraper/src/db/mysql/utils"
	"fmt"
)


func AddDecodedTransactionToDB(NetworkId, DexId int, PairId int64, Method, TxHash string, Parameters string) int64 {

	DBKeys := "network_id, dex_id, pair_id, method, transaction_hash, parameters"
	SelectStatement := fmt.Sprintf("(SELECT %d AS network_id, %d AS dex_id, %d AS pair_id, '%v' AS method, '%v' AS transaction_hash, '%v' AS parameters)", NetworkId, DexId, PairId, Method, TxHash, Parameters)
	CompareStatement := fmt.Sprintf("decoded_transactions.network_id = %d AND decoded_transactions.dex_id = %d AND decoded_transactions.pair_id = %d AND decoded_transactions.transaction_hash = '%v'", NetworkId, DexId, PairId, TxHash)

	InsertQuery := "INSERT INTO decoded_transactions(" + DBKeys + ") SELECT * FROM " + SelectStatement + " AS tmp WHERE NOT EXISTS (SELECT * FROM decoded_transactions WHERE " + CompareStatement + ") LIMIT 1"

	InsertedPairID := mysqlutils.ExecuteInsert(InsertQuery)

	return InsertedPairID

}