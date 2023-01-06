package requests

import (
	geckoterminal_helpers "atcscraper/src/api/geckoterminal"
	"atcscraper/src/requests"
	"atcscraper/src/types/geckoterminal"
	"encoding/json"
	"fmt"
)

func GetGeckoterminalDexsForNetwork(GTBuildID string, NetworkName string) geckoterminal.GeckoTerminalDexs {

	Endpoint := fmt.Sprintf("_next/data/%v/en/%v/pools.json", GTBuildID, NetworkName)

	RequestURL := geckoterminal_helpers.BuildGTAPIURL(Endpoint, false)

	Body := requests.MakeGetRequestJSON(RequestURL)

	var Dexs geckoterminal.GeckoTerminalDexs
	_ = json.Unmarshal(Body, &Dexs)

	return Dexs

}