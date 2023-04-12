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
	Body       []DecodeResult `json:"body"`
}

type Schema struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	InternalType string `json:"internalType"`
}

//type DecodeResult struct {
//	FunctionName       string                 `json:"functionName"`
//	FunctionParameters map[string]interface{} `json:"functionParameters"`
//}

type DecodeArgs struct {
	AmountIn     int64    `json:"amountIn"`
	AmountOutMin int      `json:"amountOutMin"`
	Path         []string `json:"path"`
	To           string   `json:"to"`
	Deadline     int      `json:"deadline"`
}

type DecodedInput struct {
	AmountIn     float64 `json:"_amountIn"`
	AmountOutMin float64 `json:"_amountOutMin"`
	Path         float64 `json:"_path"`
	To           float64 `json:"_to"`
	Deadline     float64 `json:"_deadline"`
}

type DecodeResult struct {
	FunctionName       string `json:"functionName"`
	FunctionParameters []struct {
		Name  string `json:"name"`
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"functionParameters"`
}