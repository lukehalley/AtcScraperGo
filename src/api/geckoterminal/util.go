package geckoterminal

import (
	"atcscraper/src/env"
	"fmt"
)

func BuildGTAPIURL(Endpoint string, IsApp bool) string {

	// Get Base URL
	GtBaseUrl := ""
	if IsApp {
		GtBaseUrl = env.LoadEnv("GT_APP_URL")
	} else {
		GtBaseUrl = env.LoadEnv("GT_BASE_URL")
	}


	// Build URL
	FullURL := fmt.Sprintf("%v/%v", GtBaseUrl, Endpoint)

	return FullURL

}