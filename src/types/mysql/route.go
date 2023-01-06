package mysql

type Route struct {
	RouteId               int      `db:"route_id"`
	NetworkId             int      `db:"network_id"`
	DexId                 int      `db:"dex_id"`
	PairId                int      `db:"pair_id"`
	Route                 string   `db:"route"`
	Method                string   `db:"method"`
	TransactionHash       string   `db:"transaction_hash"`
	BlockNumber           int      `db:"block_number"`
	AmountIn              int      `db:"amount_in"`
	AmountOut             int      `db:"amount_out"`
	TxTimestamp           int      `db:"tx_timestamp"`
}
