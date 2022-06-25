package chainlist

import (
	"atcscraper/src/api/chainlist"
	"atcscraper/src/log"
	"atcscraper/src/requests"
	"regexp"
)

func GetChainlistBuildID() string {

	RequestURL := chainlist.BuildChainlistPIURL("")

	BodyString := requests.MakeGetRequestRAW(RequestURL, 10)

	if BodyString == "" {
		log.NewError("Could Not Retrieve Chainlist Build ID")
	}

	Regex := regexp.MustCompile(`"buildId":"(.*)","isFallback"`)

	Matches := Regex.FindStringSubmatch(BodyString)

	if len(Matches) <= 0 {
		log.NewError("Could Not Retrieve Chainlist Build ID - No Matches")
	}

	BuildID := Matches[1]

	return BuildID

}