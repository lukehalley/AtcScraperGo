package requests

import (
	geckoterminal_helpers "atcscraper/src/api/geckoterminal"
	"atcscraper/src/requests"
	"atcscraper/src/types/geckoterminal"
	"encoding/json"
	"fmt"
)

func GetGeckoterminalNetworks(GTBuildID string) geckoterminal.GeckoTerminalNetworks {

	Endpoint := fmt.Sprintf("_next/data/%v/en/proof_of_reserves/exchanges.json", GTBuildID)

	RequestURL := geckoterminal_helpers.BuildGTAPIURL(Endpoint, false)

	Body := requests.MakeGetRequestJSON(RequestURL)

	var Networks geckoterminal.GeckoTerminalNetworks
	if JSONError := json.Unmarshal(Body, &Networks); JSONError != nil {
		fmt.Println("Could Not Unmarshal GT Networks JSON: ", JSONError)
	}

	return Networks

}