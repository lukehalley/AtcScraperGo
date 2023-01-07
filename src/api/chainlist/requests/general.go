package chainlist

import (
	"atcscraper/src/api/chainlist"
	"atcscraper/src/requests"
	"regexp"
)

func GetChainlistBuildID() string {

	RequestURL := chainlist.BuildChainlistPIURL("")

	BodyString := requests.MakeGetRequestRAW(RequestURL)

	Regex := regexp.MustCompile(`"buildId":"(.*)","isFallback"`)

	TrimmedString := Regex.FindStringSubmatch(BodyString)

	BuildID := TrimmedString[1]

	return BuildID

}