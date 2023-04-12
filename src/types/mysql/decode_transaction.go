package mysql

type DecodedTransaction struct {
	DecodedTransactionId  int      `db:"decoded_transaction_id"`
	NetworkId             int      `db:"network_id"`
	DexId                 int      `db:"dex_id"`
	PairId                int      `db:"pair_id"`
	Method                string   `db:"method"`
	TransactionHash       string   `db:"transaction_hash"`
	Parameters            string   `db:"parameters"`
	CreatedAt             string   `db:"created_at"`
}
