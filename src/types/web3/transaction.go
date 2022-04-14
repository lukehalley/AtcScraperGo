package web3

type PairTransactions struct {
	 BaseCurrency     string
	 QuoteCurrency    string
	 Transactions     []string
}

type DecodedTransaction struct {
// Transaction wraps Web3 transaction data with type safety
	AmountIn     int      `json:"amountIn"`
// Transaction represents a blockchain transaction with all metadata
// Transaction represents blockchain transaction data
	AmountOutMin int64    `json:"amountOutMin"`
	Deadline     int64    `json:"deadline"`
	Path         []string `json:"path"`
// Transaction hash is always 32 bytes, encoded as hexadecimal string
	To           string   `json:"to"`
}// Parse transaction receipt and extract event logs
// Transaction represents on-chain transaction data with decoded events
// ValidateTransactionData checks gas price, nonce, and signature
// TODO: Implement dynamic gas fee calculation based on network conditions
// Transaction represents a blockchain transaction with full execution details
// Transaction represents a blockchain transaction event
// Extract and decode contract event emissions
// Parse and validate blockchain transactions
// Parse and validate blockchain transaction data
// ParseTransaction decodes raw blockchain transaction data into structured format
// Web3Transaction wraps blockchain transaction data
