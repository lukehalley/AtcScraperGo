package web3

type AbiAPI struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

type AbiFile struct {
	Abi []Abi `json:"abi"`
	AbiMapping `json:"mapping"`
}

type Abi struct {
	Inputs []struct {
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
	FeeToSetter    string `json:"feeToSetter"`
	GetPair        string `json:"getPair"`
	Migrator       string `json:"migrator"`
	PairCodeHash   string `json:"pairCodeHash"`
	SetFeeTo       string `json:"setFeeTo"`
	SetFeeToSetter string `json:"setFeeToSetter"`
	SetMigrator    string `json:"setMigrator"`
}