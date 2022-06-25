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

	RequestURL := geckoterminal_helpers.BuildGTAPIURL(Endpoint, true)

	Body := requests.MakeGetRequestJSON(RequestURL, 10)

	var TransactionsWithPairData geckoterminal.GeckoTerminalPairTransactions
	_ = json.Unmarshal(Body, &TransactionsWithPairData)

	return TransactionsWithPairData

}