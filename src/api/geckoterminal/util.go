package geckoterminal

import (
	"atcscraper/src/env"
	"fmt"
)

func BuildGTAPIURL(Endpoint string) string {

	// Get Base URL
	GtBaseUrl := env.LoadEnv("GT_BASE_URL")

	return fmt.Sprintf("%v/%v", GtBaseUrl, Endpoint)


}