package atchelpers

import (
	"atcscraper/src/types"
	"net/http"
)

func CreateHTTPClientWithAuth(APIKey string) http.Client {

	HTTPClient := http.Client{
		Transport: &atctypes.AuthedTransport{
			Key:     APIKey,
			Wrapped: http.DefaultTransport,
		},
	}

	return HTTPClient

}
