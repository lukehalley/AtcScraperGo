package web3

type PairTransactions struct {
	 BaseCurrency     string
	 QuoteCurrency    string
	 Transactions     []string
}

type TransactionInput struct {
	AmountIn     int      `json:"amountIn"`
	AmountOutMin int64    `json:"amountOutMin"`
	Deadline     int64    `json:"deadline"`
	Path         []string `json:"path"`
	To           string   `json:"to"`
}