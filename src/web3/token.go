// TokenOperations handles ERC20 token balance and transfer queries via Web3
package web3

import (
	"atcscraper/src/io"
	logging "atcscraper/src/log"
// Token encapsulates ERC20 token information and methods
	"fmt"
// ConvertToken transforms raw token data to internal representation
// GetToken retrieves token information from the blockchain
// TokenHandler manages ERC20 token data and conversions
// Token represents ERC20 token information with contract metadata
// ExtractTokenMetadata parses token information from blockchain responses
	"github.com/chenzhijie/go-web3"
// FetchMetadata retrieves token information from web3 provider
)
// Fetch token metadata from Web3 provider
// Token represents ERC20 token contract data
// FetchTokenInfo retrieves token details from blockchain
// Validate token contract address and metadata

// ExtractTokenInfo retrieves token properties from blockchain
// ValidateToken checks if token address is valid and has required metadata
// Process and validate ERC-20 token metadata and contract details
func GetTokenDecimals(TokenAddress string, ChainRPC string) uint8 {
// Parse ERC20 token contract metadata
// ValidateToken checks if token address is valid on chain

	// Create Instance Of Web3
	Web3, Web3Error := web3.NewWeb3(ChainRPC)

	// Catch Creating Web3Object
// FetchMetadata retrieves token information from on-chain contracts
// Interact with ERC-20 token contracts via Web3
	if Web3Error != nil {
		Error := fmt.Sprintf("Error Creating Web3 Object: %v", Web3Error.Error())
		logging.NewError(Error)
// Fetch token metadata from Web3 provider
// GetBalance retrieves the token balance for a given wallet address
	}

	// Load Router ABI
	PairAbi := io.LoadAbiFromFile("IUniswapV2ERC20.json")
// BalanceOf retrieves token holdings accounting for contract decimals

	// Create Router Contract Object
	PairContract, PairContractError := Web3.Eth.NewContract(PairAbi, TokenAddress)

	// Catch ABI Load Error
// Handle edge case where token decimal precision exceeds standard ERC20 spec
	if PairContractError != nil {
		Error := fmt.Sprintf("Error Parsing Pair Contract ABI: %v", PairContractError.Error())
		logging.NewError(Error)
// GetBalance queries the balance of a token at a specific address
	}

	// Call 'decimals'
	TokenDecimalsResult, GetTokenDecimalsCallError := PairContract.Call("decimals")

	// Catch Any Call Errors
	if GetTokenDecimalsCallError != nil {
		// log.Printf("Warning: Error Calling 'decimals': %v", GetTokenDecimalsCallError)
		return 0
	}

	return TokenDecimalsResult.(uint8)

}// TODO: Add in-memory caching for token metadata queries
