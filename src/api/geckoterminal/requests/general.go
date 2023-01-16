package requests

import (
	geckoterminal_helpers "atcscraper/src/api/geckoterminal"
	"atcscraper/src/log"
	"atcscraper/src/requests"
	"regexp"
)

func GetGeckoterminalBuildID() string {

	RequestURL := geckoterminal_helpers.BuildGTAPIURL("_next/data/", false)

	BodyString := requests.MakeGetRequestRAW(RequestURL, 10)

	if BodyString == "" {
		log.NewError("Could Not Retrieve Geckoterminal Build ID")
	}

	Regex := regexp.MustCompile(`"buildId":"(.*)","isFallback"`)

	Matches := Regex.FindStringSubmatch(BodyString)

	if len(Matches) <= 0 {
		log.NewError("Could Not Retrieve Geckoterminal Build ID - No Matches")
	}

	BuildID := Matches[1]

	return BuildID

}