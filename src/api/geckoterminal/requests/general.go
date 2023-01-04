package requests

import (
	geckoterminal_helpers "atcscraper/src/api/geckoterminal"
	"atcscraper/src/requests"
	"regexp"
)

func GetGeckoterminalBuildID() string {

	RequestURL := geckoterminal_helpers.BuildGTAPIURL("_next/data/", false)

	BodyString := requests.MakeGetRequestRAW(RequestURL)

	Regex := regexp.MustCompile(`"buildId":"(.*)","isFallback"`)

	TrimmedString := Regex.FindStringSubmatch(BodyString)

	BuildID := TrimmedString[1]

	return BuildID

}