package main

import (
	atcqueries "atcscraper/src/bitquery/schema"
	"context"
	"fmt"
	"github.com/Khan/genqlient/graphql"
	"net/http"
	"time"
)

type authedTransport struct {
	key     string
	wrapped http.RoundTripper
}

func (Auth *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("X-API-KEY", Auth.key)
	return Auth.wrapped.RoundTrip(req)
}

func main() {

	BitqueryApiKey := "BQYZVuAE5Tl7zOvKBmQajtsyaH5JaqmG"

	HTTPClient := http.Client{
		Transport: &authedTransport{
			key:     BitqueryApiKey,
			wrapped: http.DefaultTransport,
		},
	}
	graphqlClient := graphql.NewClient("https://graphql.bitquery.io", &HTTPClient)

	FromTime, TimeParseError := time.Parse(
		time.RFC3339,
		"2022-12-18T15:04:05Z")

	if TimeParseError != nil {
		fmt.Println(TimeParseError)
	}

	QueryResult, QueryErr := atcqueries.GetTopNDexsForNetwork(context.Background(), graphqlClient, "klaytn", 25, 0, FromTime)

	if QueryErr != nil {
		return
	}

	fmt.Println("QueryResult", QueryResult)
}
