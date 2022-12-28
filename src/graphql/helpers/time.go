package atchelpers

import (
	"fmt"
	"time"
)

func GenerateGraphQLTimestamp(Timestamp string) time.Time {

	FromTime, TimeParseError := time.Parse(
		time.RFC3339,
		Timestamp)

	if TimeParseError != nil {
		fmt.Println(TimeParseError)
	}

	return FromTime
}
