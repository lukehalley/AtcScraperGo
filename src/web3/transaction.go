package web3



import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func GetTransactionReceipt(NetworkRPC string, TXHash string) *types.Transaction {

	// Create Web3 Client
	Web3Client, Web3ClientError := ethclient.Dial(NetworkRPC)

	// Catch Creating Web3 Client
	if Web3ClientError != nil {
		log.Fatal(Web3ClientError)
	}

	// ConvertHexToHash
	TxHash := common.HexToHash(TXHash)

	var TransactionReceipt *types.Transaction
	TransactionReceipt, _, TransactionReceiptError := Web3Client.TransactionByHash(context.Background(), TxHash)
	if TransactionReceiptError != nil {
		log.Fatal(TransactionReceiptError)
	}

	return TransactionReceipt

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
	msg, err := tx.AsMessage(types.LatestSignerForChainID(tx.ChainId()), nil)
	if err != nil {
		log.Fatal(err)
	}
	return msg
}
