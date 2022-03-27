package web3

import (
	"atcscraper/src/io"
	"github.com/chenzhijie/go-web3"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

func GetPairAddress(BaseCurrencyTokenAddress string, QuoteCurrencyTokenAddress string, FactoryAddress string, NetworkRPC string) string {

	// Create Instance Of Web3
	Web3, Web3Error := web3.NewWeb3(NetworkRPC)

	// Catch Creating Web3Object
	if Web3Error != nil {
		log.Fatal(Web3Error)
	}

	// Load Factory ABI
	FactoryAbi := io.LoadAbi("IUniswapV2Factory.json")

	// Create Factory Contract Object
	FactoryContract, FactoryContractError := Web3.Eth.NewContract(FactoryAbi, FactoryAddress)
	if FactoryContractError != nil {
		log.Fatal(FactoryContractError)
	}

	// Call 'getPair'
	PairAddress, PairAddressError := FactoryContract.Call("getPair", common.HexToAddress(BaseCurrencyTokenAddress), common.HexToAddress(QuoteCurrencyTokenAddress))

	// Catch Any Call Errors
	if PairAddressError != nil {
		log.Fatalf("Error Calling 'getPair': %v", PairAddressError)
	}

	return GetHexAddress(PairAddress)

}