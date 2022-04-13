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

func DecodeTransactionInputData(ABI string, TransactionInputData []byte) (string, web3.TransactionInput) {

	DecodedInputData := web3.TransactionInput{}

	ABIObject, ABIReadError := abi.JSON(strings.NewReader(ABI))
	if ABIReadError != nil {
		return "", DecodedInputData
	}

	MethodSigData := TransactionInputData[:4]
	Method, MethodError := ABIObject.MethodById(MethodSigData)
	if MethodError != nil {
		return "", DecodedInputData
	}

	InputsSigData := TransactionInputData[4:]
	InputsMap := make(map[string]interface{})
	if UnpackError := Method.Inputs.UnpackIntoMap(InputsMap, InputsSigData); UnpackError != nil {
		return "", DecodedInputData
	}

	InputDataJSON := MapToJson(InputsMap)
	JSONParseError := json.Unmarshal([]byte(InputDataJSON), &DecodedInputData)
	if JSONParseError != nil {
		return "", DecodedInputData
	}

	return Method.Name, DecodedInputData

}
