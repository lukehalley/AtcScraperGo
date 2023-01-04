package requests

import (
	geckoterminal2 "atcscraper/src/api/geckoterminal"
	"atcscraper/src/requests"
	"atcscraper/src/types/geckoterminal"
	"encoding/json"
	"fmt"
	"regexp"
)

func GetGeckoterminalBuildID() string {

	RequestURL := geckoterminal2.BuildGTAPIURL("_next/data/")

	BodyString := requests.MakeGetRequestRAW(RequestURL)

	Regex := regexp.MustCompile(`"buildId":"(.*)","isFallback"`)

	TrimmedString := Regex.FindStringSubmatch(BodyString)

	BuildID := TrimmedString[1]

	return BuildID

}

func GetGeckoterminalNetworks(GTBuildID string) geckoterminal.GeckoTerminalNetworks {

	Endpoint := fmt.Sprintf("_next/data/%v/en/proof_of_reserves/exchanges.json", GTBuildID)

	RequestURL := geckoterminal2.BuildGTAPIURL(Endpoint)

	Body := requests.MakeGetRequestJSON(RequestURL)

	var Networks geckoterminal.GeckoTerminalNetworks
	if JSONError := json.Unmarshal(Body, &Networks); JSONError != nil {
		fmt.Println("Could Not Unmarshal GT Networks JSON: ", JSONError)
	}

	return Networks

}

func GetGeckoterminalDexsForNetwork(GTBuildID string, NetworkName string) geckoterminal.GeckoTerminalDexs {

	Endpoint := fmt.Sprintf("_next/data/%v/en/%v/pools.json", GTBuildID, NetworkName)

	RequestURL := geckoterminal2.BuildGTAPIURL(Endpoint)

	Body := requests.MakeGetRequestJSON(RequestURL)

	var Dexs geckoterminal.GeckoTerminalDexs
	if JSONError := json.Unmarshal(Body, &Dexs); JSONError != nil {
		fmt.Println("Could Not Unmarshal GT Dexs JSON: ", JSONError)
	}

	return Dexs

}

func GetGeckoterminalPairTransactionsAndTokens(NetworkName string, PairAddress string, Page int) geckoterminal.GeckoTerminalPairTransactions {

	Endpoint := fmt.Sprintf("api/p1/%v/pools/%v/swaps?include=from_token,to_token&page=%d", NetworkName, PairAddress, Page)

	RequestURL := geckoterminal2.BuildGTAPIURL(Endpoint)

	Body := requests.MakeGetRequestJSON(RequestURL)

	var TransactionsWithPairData geckoterminal.GeckoTerminalPairTransactions
	if JSONError := json.Unmarshal(Body, &TransactionsWithPairData); JSONError != nil {
		fmt.Println("Could Not Unmarshal GT Transactions And Pair JSON: ", JSONError)
	}

	return TransactionsWithPairData

}