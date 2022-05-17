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
	FactoryAbi := io.LoadAbiAsString("IUniswapV2Factory.json")

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

func GetPairFactoryAddress(PairAddress string, NetworkRPC string) string {

	NetworkRPC = "123" + NetworkRPC

	// Create Instance Of Web3
	Web3, Web3Error := web3.NewWeb3(NetworkRPC)

	// Catch Creating Web3Object
	if Web3Error != nil {
		log.Fatal(Web3Error)
	}

	// Load Factory ABI
	PairAbi := io.LoadAbiAsString("IUniswapV2Pair.json")

	// Create Pair Contract Object
	PairContract, PairContractError := Web3.Eth.NewContract(PairAbi, PairAddress)
	if PairContractError != nil {
		log.Fatal(PairContractError)
	}

	// Call 'getPair'
	FactoryAddress, FactoryAddressError := PairContract.Call("factory")

	// Catch Any Call Errors
	if FactoryAddressError != nil {
		return ""
	}

	return GetHexAddress(FactoryAddress)

}

func GetPairTokenAddresses(PairAddress string, NetworkRPC string) (string, string) {

	// Create Instance Of Web3
	Web3, Web3Error := web3.NewWeb3(NetworkRPC)

	// Catch Creating Web3Object
	if Web3Error != nil {
		log.Fatal(Web3Error)
	}

	// Load Factory ABI
	PairAbi := io.LoadAbiAsString("IUniswapV2Pair.json")

	// Create Pair Contract Object
	PairContract, PairContractError := Web3.Eth.NewContract(PairAbi, PairAddress)
	if PairContractError != nil {
		log.Fatal(PairContractError)
	}

	// Call 'token0'
	BaseCurrencyAddress, BaseCurrencyAddressError := PairContract.Call("token0")

	// Catch Any Call Errors
	if BaseCurrencyAddressError != nil {
		return "", ""
	}

	// Call 'token1'
	QuoteCurrencyAddress, QuoteCurrencyAddressError := PairContract.Call("token1")

	// Catch Any Call Errors
	if QuoteCurrencyAddressError != nil {
		return "", ""
	}

	return GetHexAddress(BaseCurrencyAddress), GetHexAddress(QuoteCurrencyAddress)

}