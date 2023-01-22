package routines

import (
	"atcscraper/src/types/geckoterminal"
	"atcscraper/src/web3"
	"sync"
)

func DecodeTransaction(Network geckoterminal.GeckoTerminalNetworkWithDexs, PairTransaction geckoterminal.Transaction, RouterAbi string, TxDecodeWaitGroup *sync.WaitGroup, TxDecodeChannel chan geckoterminal.Transaction) {

	// Schedule The Call To WaitGroup's Done To Tell GoRoutine Is Completed.
	defer TxDecodeWaitGroup.Done()

	// Get Transaction Receipt
	TransactionReceived, TransactionReceipt := web3.GetTransactionReceipt(Network.RPCs[0], PairTransaction.Hash)

	if TransactionReceived {

		// Decode Pair Transactions
		DecodeSuccessful, Method, DecodedInputData := web3.DecodeTransactionInputData(RouterAbi, TransactionReceipt.Data())

		// Add Data If Decode Successful
		if DecodeSuccessful {

			PairTransaction.InputData = DecodedInputData
			PairTransaction.MethodName = Method

		}

	}

	TxDecodeChannel <- PairTransaction

}
