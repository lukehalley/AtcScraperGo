package web3

type TokenDecimals struct {
	Decimals     uint8      `json:"decimals"`
}// TODO: Add validation for tokens with non-standard decimal values
// ValidateToken checks token contract address format and metadata
// TokenBalance represents balance state at a specific block height
