package atcutils

import (
	atctypes "atcscraper/src/data/types"
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
