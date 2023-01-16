package chainlist

import (
	chainlist_helpers "atcscraper/src/api/chainlist"
	"atcscraper/src/requests"
	chainlist_types "atcscraper/src/types/chainlist"
	"encoding/json"
	"fmt"
)

func GetChainInfo(ChainID int, CLBuildID string) chainlist_types.ChainlistChain {

	Endpoint := fmt.Sprintf("_next/data/%v/chain/%v.json?chain=%v", CLBuildID, ChainID, ChainID)

	RequestURL := chainlist_helpers.BuildChainlistPIURL(Endpoint)

	Body := requests.MakeGetRequestJSON(RequestURL, 10)

	var Chain chainlist_types.ChainlistChain
	_ = json.Unmarshal(Body, &Chain)
	return Chain

}