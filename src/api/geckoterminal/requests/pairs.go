package requests

import (
	geckoterminal_helpers "atcscraper/src/api/geckoterminal"
	"atcscraper/src/requests"
	"atcscraper/src/types/geckoterminal"
	"encoding/json"
	"fmt"
)

func GetGeckoterminalPairTransactionsAndTokens(NetworkName string, PairAddress string, Page int) geckoterminal.GeckoTerminalPairTransactions {

	Endpoint := fmt.Sprintf("api/p1/%v/pools/%v/swaps?include=from_token,to_token&page=%d", NetworkName, PairAddress, Page)

	RequestURL := geckoterminal_helpers.BuildGTAPIURL(Endpoint)

	Body := requests.MakeGetRequestJSON(RequestURL)

	var TransactionsWithPairData geckoterminal.GeckoTerminalPairTransactions
	if JSONError := json.Unmarshal(Body, &TransactionsWithPairData); JSONError != nil {
		fmt.Println("Could Not Unmarshal GT Transactions And Pair JSON: ", JSONError)
	}

	return TransactionsWithPairData

}