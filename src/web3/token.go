package web3

import (
	"atcscraper/src/io"
	logging "atcscraper/src/log"
	"fmt"
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

	// Create Router Contract Object
	PairContract, PairContractError := Web3.Eth.NewContract(PairAbi, TokenAddress)

	// Catch ABI Load Error
	if PairContractError != nil {
		Error := fmt.Sprintf("Error Parsing Pair Contract ABI: %v", PairContractError.Error())
		logging.NewError(Error)
	}

	// Call 'decimals'
	TokenDecimalsResult, GetTokenDecimalsCallError := PairContract.Call("decimals")

	// Catch Any Call Errors
	if GetTokenDecimalsCallError != nil {
		// log.Printf("Warning: Error Calling 'decimals': %v", GetTokenDecimalsCallError)
		return 0
	}

	return TokenDecimalsResult.(uint8)

}