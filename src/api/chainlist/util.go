package chainlist

import (
	"atcscraper/src/env"
	"fmt"
)

func BuildChainlistPIURL(Endpoint string) string {

	// Get Base URL
	ChainlistBaseUrl := env.LoadEnv("CL_BASE_URL")

	// Build URL
	FullURL := fmt.Sprintf("%v/%v", ChainlistBaseUrl, Endpoint)

	return FullURL

}