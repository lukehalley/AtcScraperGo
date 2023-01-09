package web3

import (
	"atcscraper/src/types/web3"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
)

func MapToJson(param map[string]interface{}) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}

func DecodeTransactionInputData(ABI string, TransactionInputData []byte) (bool, string, web3.DecodedTransaction) {

	DecodedInputData := web3.DecodedTransaction{}

	ABIObject, ABIReadError := abi.JSON(strings.NewReader(ABI))
	if ABIReadError != nil {
		return false, "", DecodedInputData
	}

	if len(TransactionInputData) <= 8 {
		return false, "", DecodedInputData
	}

	MethodSigData := TransactionInputData[:4]
	Method, MethodError := ABIObject.MethodById(MethodSigData)
	if MethodError != nil {
		return false, "", DecodedInputData
	}

	InputsSigData := TransactionInputData[4:]
	InputsMap := make(map[string]interface{})
	if UnpackError := Method.Inputs.UnpackIntoMap(InputsMap, InputsSigData); UnpackError != nil {
		return false, "", DecodedInputData
	}

	InputDataJSON := MapToJson(InputsMap)
	JSONParseError := json.Unmarshal([]byte(InputDataJSON), &DecodedInputData)
	if JSONParseError != nil {
		return false, "", DecodedInputData
	}

	return true, Method.Name, DecodedInputData

}
