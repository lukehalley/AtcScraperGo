package atchelpers

import (
	"atcscraper/src/types/bitquery"
	"net/http"
)

func CreateHTTPClientWithAuth(APIKey string) http.Client {

	HTTPClient := http.Client{
		Transport: &bitquery.AuthedTransport{
			Key:     APIKey,
			Wrapped: http.DefaultTransport,
		},
	}

	return HTTPClient

}
