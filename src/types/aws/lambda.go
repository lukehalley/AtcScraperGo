package aws

type DecodeLambdaParams struct {
	RpcUrl          string `json:"rpc_url"`
	TxHash          string `json:"tx_hash"`
	ContractHash    string `json:"contract_hash"`
	ContractAbiDBId int64  `json:"contract_abi_db_id"`
}

type DecodeLambdaResponse struct {
	StatusCode int    `json:"statusCode"`
	Msg        string `json:"msg"`
	Body       struct {
		Function string     `json:"function"`
		Args     DecodeArgs `json:"args"`
	} `json:"body"`
}

type Schema struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	InternalType string `json:"internalType"`
}

type DecodeArgs struct {
	AmountIn     int64    `json:"amountIn"`
	AmountOutMin int      `json:"amountOutMin"`
	Path         []string `json:"path"`
	To           string   `json:"to"`
	Deadline     int      `json:"deadline"`
}
