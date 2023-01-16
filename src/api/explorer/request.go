package explorer

import (
	"atcscraper/src/requests"
	"atcscraper/src/types/web3"
	"encoding/json"
)

func GetAbiFromExplorer(APIUrl string, APIKey string, ContractAddress string) (web3.AbiAPI, bool) {

	// Build The API URL
	Endpoint := BuildAPIEndpoint(APIUrl, APIKey, "contract", "getabi", ContractAddress)

	// Call The Explorer API To Get The ABI
	Body := requests.MakeGetRequestJSON(Endpoint)

	// Convert The File Into A Struct
	JSONAbiStruct := web3.AbiAPI{}

	// Decode The JSON Request To Struct
	_ = json.Unmarshal(Body, &JSONAbiStruct)

	// Check If We Got An ABI Back
	if JSONAbiStruct.Message == "OK" {

		return JSONAbiStruct, true

	} else {

		return JSONAbiStruct, false

	}



}