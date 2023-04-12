package routines

import (
	"atcscraper/src/aws/lambda"
	"atcscraper/src/types/geckoterminal"
	"atcscraper/src/web3"
	"sync"
)

func DecodeTransaction(Network geckoterminal.GeckoTerminalNetworkWithDexs, PairTransaction geckoterminal.Transaction, RouterAddress string, RouterAbiDbId int64, TxDecodeWaitGroup *sync.WaitGroup, TxDecodeChannel chan geckoterminal.Transaction) {

	// Schedule The Call To WaitGroup's Done To Tell GoRoutine Is Completed.
	defer TxDecodeWaitGroup.Done()

	// Get Transaction Receipt
	TransactionReceived, TransactionReceipt := web3.GetTransactionReceipt(Network.RPCs[0], PairTransaction.Hash)

	if TransactionReceipt != nil {

		if TransactionReceipt.To() != nil {

			TransactionToDex := TransactionReceipt.To().Hex()
			TransactionData := TransactionReceipt.Data()

			MethodSignature := web3.GetMethodSignature(TransactionData)
			GetMethodParams := web3.GetMethodParams(TransactionData)

			if TransactionToDex != "" && MethodSignature != "" && GetMethodParams != "" {

				TransactionRightAddress := TransactionToDex == RouterAddress

				if TransactionReceived && TransactionRightAddress {

					// Decode Pair Transactions
					TxHash := TransactionReceipt.Hash().Hex()

					DecodeLambdaResult := lambda.CallDecodeLambda(Network.RPCs[0], TxHash, RouterAddress, RouterAbiDbId)

					if DecodeLambdaResult.StatusCode == 200 {
						PairTransaction.DecodeResults = DecodeLambdaResult.Body
					}

				}
			}
		}
	}

	TxDecodeChannel <- PairTransaction

}
