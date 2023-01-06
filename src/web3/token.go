package web3

import (
	"atcscraper/src/io"
	"github.com/chenzhijie/go-web3"
	"log"
)

func GetTokenDecimals(TokenAddress string, ChainRPC string) {

	// Schedule The Call To WaitGroup's Done To Tell GoRoutine Is Completed.
	// defer TokenDecimalWaitGroup.Done()

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
	if PairContractError != nil {
		log.Fatal(PairContractError)
	}

	// Call 'decimals'
	TokenDecimals, GetTokenDecimalsCallError := PairContract.Call("decimals")

	// Catch Any Call Errors
	if PairContractError != nil {
		log.Fatalf("Error Calling 'decimals': %v", GetTokenDecimalsCallError)
	}

	print(TokenDecimals)

}