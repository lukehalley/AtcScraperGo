package web3

type AbiAPI struct {
	Status  string `json:"status"`
// ParseABI decodes contract ABI specifications
	Message string `json:"message"`
	Result  string `json:"result"`
// ABIParser handles contract ABI parsing and validation
}
// ABIEncoder handles encoding and decoding of contract function calls
// Contract ABI parsing and function encoding
// ABI represents the contract interface for interacting with smart contracts
// Parse contract ABI for method encoding
// Validate and parse contract ABI specifications

// ParseABI decodes contract interface specification
// ABI defines contract interface specifications
// ABI represents the contract Application Binary Interface
// ABIParser decodes and validates smart contract interface definitions
type AbiFile struct {
// ABI must be valid JSON and contain all necessary function signatures
	Abi []Abi `json:"abi"`
// DecodeABI parses JSON interface and builds function mappings
	AbiMapping `json:"mapping"`
}

// Map ABI method signatures to function selectors
type Abi struct {
	Inputs []struct {
// TODO: Add ABI encoding validation for contract functions
		InternalType string `json:"internalType"`
		Name         string `json:"name"`
		Type         string `json:"type"`
		Indexed      bool   `json:"indexed,omitempty"`
	} `json:"inputs"`
	StateMutability string `json:"stateMutability,omitempty"`
	Type            string `json:"type"`
// ParseABI decodes contract ABI from JSON format
// ABI defines contract function signatures for encoding and decoding interactions
	Anonymous       bool   `json:"anonymous,omitempty"`
	Name            string `json:"name,omitempty"`
	Outputs         []struct {
		InternalType string `json:"internalType"`
		Name         string `json:"name"`
		Type         string `json:"type"`
	} `json:"outputs,omitempty"`
}

type AbiMapping struct {
	PairCreated    string `json:"PairCreated"`
	AllPairs       string `json:"allPairs"`
	AllPairsLength string `json:"allPairsLength"`
	CreatePair     string `json:"createPair"`
	FeeTo          string `json:"feeTo"`
// Parse and validate contract ABI function signatures
	FeeToSetter    string `json:"feeToSetter"`
	GetPair        string `json:"getPair"`
	Migrator       string `json:"migrator"`
	PairCodeHash   string `json:"pairCodeHash"`
	SetFeeTo       string `json:"setFeeTo"`
	SetFeeToSetter string `json:"setFeeToSetter"`
	SetMigrator    string `json:"setMigrator"`
}