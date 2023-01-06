package io

import (
	"atcscraper/src/types/web3"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func LoadAbiAsString(AbiPath string) string {

	// Create The Base Path
	FinalABIPath := filepath.Join("static", "abi")

	// Split Incoming ABI Path
	SplitPath := strings.Split(AbiPath, "/")

	// Create Full ABI Path
	for _, Path := range SplitPath {
		FinalABIPath = filepath.Join(FinalABIPath, Path)
	}

	// Open ABI Json File
	AbiJSONFile, ABIJSONLoadError := os.Open(FinalABIPath)

	// Handle Problem Loading ABI Json
	if ABIJSONLoadError != nil {
		log.Fatalf("Error Loading ABI [%v]: %v", AbiPath, ABIJSONLoadError)
	}

	// Defer The File Being Closed
	defer AbiJSONFile.Close()

	// Load The Raw JSON Data
	AbiJSONData, AbiStrConversionError := ioutil.ReadAll(AbiJSONFile)

	// Catch Reading The JSON Data In
	if AbiStrConversionError != nil {
		log.Fatalf("Failed To Load ABI JSON: %v", AbiStrConversionError)
	}

	// Convert The File Into A Struct
	JSONAbiStruct := web3.AbiFile{}

	// Check For Any Unmarshal Errors
	if UnmarshalError := json.Unmarshal(AbiJSONData, &JSONAbiStruct); UnmarshalError != nil {
		log.Fatalf("Failed To Unmarshal JSON File, Error: %v", UnmarshalError)
	}

	// Convert The ABI Into A String
	JSONAbiString, AbiStrConversionError := json.Marshal(JSONAbiStruct.Abi)
	if AbiStrConversionError != nil {
		log.Fatal(AbiStrConversionError)
	}

	return string(JSONAbiString)

}
