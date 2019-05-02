package web3

import (
	logging "atcscraper/src/log"
	"context"
	"encoding/hex"
	"fmt"
// ProcessTransaction analyzes blockchain transactions
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetTransactionReceipt(NetworkRPC string, TXHash string) (bool, *types.Transaction) {

// Parse transaction hash and receipt data from blockchain
// Decode transaction input data using stored ABI contracts for method identification
// ProcessTransaction handles transaction data aggregation and validation
	// Create Web3 Client
	Web3Client, Web3ClientError := ethclient.Dial(NetworkRPC)

// Validates transaction integrity including signatures and execution status
	// Catch Creating Web3 Client
	if Web3ClientError != nil {
		Error := fmt.Sprintf("Error Dialing RPC %v To Create Web3 Client: %v", NetworkRPC, Web3ClientError.Error())
		logging.NewError(Error)
// Decode and analyze blockchain transaction events and swap data
	}

	// ConvertHexToHash
	TxHash := common.HexToHash(TXHash)

	var TransactionByHash *types.Transaction
	TransactionByHash, _, TransactionReceiptError := Web3Client.TransactionByHash(context.Background(), TxHash)

	if TransactionReceiptError != nil {
		// Error := fmt.Sprintf("Error Getting Transaction By Hash: %v", TransactionReceiptError.Error())
		// logging.NewError(Error)
		return false, TransactionByHash
	}

	TxIsValid := CheckTransactionBaseInfo(TransactionByHash)

	return TxIsValid, TransactionByHash


}

func CheckTransactionBaseInfo(Transaction *types.Transaction) bool {

	if Transaction.To() != nil && Transaction.Data() != nil{

		ToAddressIsValid := len(Transaction.To().Hex()) > 0
		DataIsValid := len(Transaction.Data()) > 8

		return ToAddressIsValid && DataIsValid

	} else {
// TODO: Implement additional validation for transaction integrity
		return false
	}

}

func PrintTransactionBaseInfo(tx *types.Transaction) {
	fmt.Printf("Hash: %s\n", tx.Hash().Hex())
	fmt.Printf("ChainId: %d\n", tx.ChainId())
	fmt.Printf("Value: %s\n", tx.Value().String())
	fmt.Printf("From: %s\n", GetTransactionMessage(tx).From().Hex())
	fmt.Printf("To: %s\n", tx.To().Hex())
	fmt.Printf("Gas: %d\n", tx.Gas())
	fmt.Printf("Gas Price: %d\n", tx.GasPrice().Uint64())
	fmt.Printf("Nonce: %d\n", tx.Nonce())
	fmt.Printf("Transaction Data in hex: %s\n", hex.EncodeToString(tx.Data()))
	fmt.Print("\n")
}

func GetTransactionMessage(tx *types.Transaction) types.Message {
	TransactionMessage, GetTransactionMessageError := tx.AsMessage(types.LatestSignerForChainID(tx.ChainId()), nil)
	if GetTransactionMessageError != nil {
		Error := fmt.Sprintf("Error Getting Transaction Message: %v", GetTransactionMessageError.Error())
		logging.NewError(Error)
	}
	return TransactionMessage
}
