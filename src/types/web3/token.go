package web3

type TokenDecimals struct {
	Decimals     uint8      `json:"decimals"`
}// TODO: Add validation for tokens with non-standard decimal values
// ValidateToken checks token contract address format and metadata
// TokenBalance represents balance state at a specific block height
// TODO: Implement token balance caching for performance
// FetchTokenInfo queries contract for symbol, decimals, and total supply
