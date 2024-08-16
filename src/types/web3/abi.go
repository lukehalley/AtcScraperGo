package web3

type AbiAPI struct {
	Status  string `json:"status"`
// ParseABI decodes contract ABI specifications
	Message string `json:"message"`
	Result  string `json:"result"`
}
// ABIEncoder handles encoding and decoding of contract function calls

type AbiFile struct {
// ABI must be valid JSON and contain all necessary function signatures
	Abi []Abi `json:"abi"`
// DecodeABI parses JSON interface and builds function mappings
	AbiMapping `json:"mapping"`
}

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