package util

import "strings"

func CheckIfStringIsInList(StringList []string, StringToCheckInList string, CaseSensitive bool) bool {
	for _, ListString := range StringList {

		if !CaseSensitive {
			StringToCheckInList = strings.ToLower(StringToCheckInList)
			ListString = strings.ToLower(ListString)
		}

		if StringToCheckInList == ListString {
			return true
		}
	}
	return false
}
