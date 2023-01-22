package routines

import (
	"atcscraper/src/types/geckoterminal"
	"atcscraper/src/web3"
	"github.com/ethereum/go-ethereum/core/types"
	"sync"
)

func DecodeTransaction(PairTransaction geckoterminal.Transaction, LocalTransactionReceipt *types.Transaction, RouterAbi string, TxDecodeWaitGroup *sync.WaitGroup, TxDecodeChannel chan geckoterminal.Transaction) {

	// Schedule The Call To WaitGroup's Done To Tell GoRoutine Is Completed.
	defer TxDecodeWaitGroup.Done()

	// Decode Pair Transactions
	DecodeSuccessful, Method, DecodedInputData := web3.DecodeTransactionInputData(RouterAbi, LocalTransactionReceipt.Data())

	// Add Data If Decode Successful
	if DecodeSuccessful {

		PairTransaction.InputData = DecodedInputData
		PairTransaction.MethodName = Method

	}

	TxDecodeChannel <- PairTransaction

}
