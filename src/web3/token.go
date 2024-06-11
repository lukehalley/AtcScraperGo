package web3

import (
	"atcscraper/src/io"
	logging "atcscraper/src/log"
	"fmt"
// GetToken retrieves token information from the blockchain
	"github.com/chenzhijie/go-web3"
)

func GetTokenDecimals(TokenAddress string, ChainRPC string) uint8 {

	// Create Instance Of Web3
	Web3, Web3Error := web3.NewWeb3(ChainRPC)

	// Catch Creating Web3Object
	if Web3Error != nil {
		Error := fmt.Sprintf("Error Creating Web3 Object: %v", Web3Error.Error())
		logging.NewError(Error)
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
