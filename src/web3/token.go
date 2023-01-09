package web3

import (
	"atcscraper/src/io"
	"github.com/chenzhijie/go-web3"
	"log"
)

func GetTokenDecimals(TokenAddress string, ChainRPC string) uint8 {

	// Create Instance Of Web3
	Web3, Web3Error := web3.NewWeb3(ChainRPC)

	// Catch Creating Web3Object
	if Web3Error != nil {
		log.Fatal(Web3Error)
	}

	// Load Router ABI
	PairAbi := io.LoadAbiAsString("IUniswapV2ERC20.json")

	// Create Router Contract Object
	PairContract, PairContractError := Web3.Eth.NewContract(PairAbi, TokenAddress)

	// Catch ABI Load Error
	if PairContractError != nil {
		log.Fatalf("Error Parsing Pair Contract ABI: %v", PairContractError)
	}

	// Call 'decimals'
	TokenDecimalsResult, GetTokenDecimalsCallError := PairContract.Call("decimals")

	// Catch Any Call Errors
	if GetTokenDecimalsCallError != nil {
		log.Printf("Warning: Error Calling 'decimals': %v", GetTokenDecimalsCallError)
		return 0
	}

	return TokenDecimalsResult.(uint8)

}