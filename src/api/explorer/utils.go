package explorer

import (
	"fmt"
)

func BuildAPIEndpoint(APIEndpoint string, APIKey string, Module string, Action string, Address string) string {

	BuiltAPIEndpoint := fmt.Sprintf("%v/api?module=%v&action=%v&address=%v", APIEndpoint, Module, Action, Address)

	if APIKey != "" {
		BuiltAPIEndpoint = fmt.Sprintf("%v&apikey=%v", BuiltAPIEndpoint, APIKey)
	}

	return BuiltAPIEndpoint

}
