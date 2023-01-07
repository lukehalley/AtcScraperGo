package requests

import (
	"net/http"
	"time"
)

func CheckURLIsReachable(URL string) bool {
	Timeout := 3 * time.Second
	HTTPClient := http.Client{
		Timeout: Timeout,
	}
	Dial, DialError := HTTPClient.Get(URL)
	if Dial != nil {
		if Dial.StatusCode == 200 && DialError == nil {
			return true
		} else {
			return false
		}
	} else {
		return false
	}

}
