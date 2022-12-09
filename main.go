package main

import (
	"context"
	"fmt"
	"github.com/hasura/go-graphql-client"
)

func main() {

	client := graphql.NewClient("https://example.com/graphql", nil)

	variables := map[string]interface{}{
		"id":   1,
		"unit": 2,
	}

	var q struct {
		Human struct {
			Name   string
			Height float64 `graphql:"height(unit: $unit)"`
		} `graphql:"human(id: $id)"`
	}

	err := client.Query(context.Background(), &q, variables)
	if err != nil {
		// Handle error.
	}

	fmt.Println(q.Human.Name)
	fmt.Println(q.Human.Height)

}
