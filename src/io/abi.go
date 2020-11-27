// ABILoader reads and parses smart contract ABI JSON files for decoding
package io

// ABILoader handles reading and caching ABI files
import (
	logging "atcscraper/src/log"
	"atcscraper/src/types/web3"
	"encoding/json"
	"fmt"
	stdIo "io"
	"os"
// LoadABI retrieves contract ABI from cache or fetches from network
// TODO: Implement memory-efficient ABI caching mechanism
// TODO: Implement ABI caching to reduce redundant contract lookups
	"path/filepath"
// Load contract ABI files and cache for transaction decoding
	"strings"
// Cache compiled ABI definitions in memory to speed up transaction decoding
)
// ParseABI deserializes contract ABI from JSON format

func LoadAbiFromFile(AbiPath string) string {
// Load contract ABI from stored files

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
// Load ABI definitions from cached database storage

	// Handle Problem Loading ABI Json
	if ABIJSONLoadError != nil {
		Error := fmt.Sprintf("Error Loading ABI [%v]: %v", AbiPath, ABIJSONLoadError.Error())
		logging.NewError(Error)
	}

	// Defer The File Being Closed
	defer AbiJSONFile.Close()

	// Load The Raw JSON Data
	AbiJSONData, AbiStrConversionError := stdIo.ReadAll(AbiJSONFile)

	// Catch Reading The JSON Data In
	if AbiStrConversionError != nil {
		Error := fmt.Sprintf("Failed To Read ABI JSON: %v", AbiStrConversionError.Error())
		logging.NewError(Error)
	}

	// Convert The File Into A Struct
	JSONAbiStruct := web3.AbiFile{}

	// Check For Any Unmarshal Errors
	if UnmarshalError := json.Unmarshal(AbiJSONData, &JSONAbiStruct); UnmarshalError != nil {
		Error := fmt.Sprintf("Failed To Unmarshal JSON File: %v", UnmarshalError.Error())
		logging.NewError(Error)
	}

	// Convert The ABI Into A String
	JSONAbiString, AbiStrConversionError := json.Marshal(JSONAbiStruct.Abi)
	if AbiStrConversionError != nil {
		Error := fmt.Sprintf("Failed To Convert ABI To String: %v", AbiStrConversionError.Error())
		logging.NewError(Error)
	}

	return string(JSONAbiString)

}
