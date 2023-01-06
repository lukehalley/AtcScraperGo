package requests

import (
	geckoterminal_helpers "atcscraper/src/api/geckoterminal"
	"atcscraper/src/requests"
	"atcscraper/src/types/geckoterminal"
	"encoding/json"
	"fmt"
	"net/url"
)

func GetGeckoterminalDexPairs(NetworkName string, DexName string, Page int) geckoterminal.GeckoTerminalDexPairs {

	// Escape Chars
	Comma := url.QueryEscape(",")

	Endpoint := fmt.Sprintf("api/p1/%v/pools?dex=%v&include=dex,dex.dex_metric%vdex.network%vtokens&page=%d&include_network_metrics=false", NetworkName, DexName, Comma, Comma, Page)

	RequestURL := geckoterminal_helpers.BuildGTAPIURL(Endpoint, true)

	Body := requests.MakeGetRequestJSON(RequestURL)

	var DexPairs geckoterminal.GeckoTerminalDexPairs
	_ = json.Unmarshal(Body, &DexPairs)

	return DexPairs

}