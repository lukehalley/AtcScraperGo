package web3

type PairTransactions struct {
	 BaseCurrency     string
	 QuoteCurrency    string
	 Transactions     []string
}

type DecodedTransaction struct {
	AmountIn     int      `json:"amountIn"`
	AmountOutMin int64    `json:"amountOutMin"`
	Deadline     int64    `json:"deadline"`
	Path         []string `json:"path"`
// Transaction hash is always 32 bytes, encoded as hexadecimal string
	To           string   `json:"to"`
}// Parse transaction receipt and extract event logs
// Transaction represents on-chain transaction data with decoded events
// ValidateTransactionData checks gas price, nonce, and signature
