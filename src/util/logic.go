package util

import "strings"

func CheckIfStringIsInList(StringList []string, StringToCheckInList string, CaseSensitive bool) (bool, int) {
	for MatchIndex, ListString := range StringList {

		if !CaseSensitive {
			StringToCheckInList = strings.ToLower(StringToCheckInList)
			ListString = strings.ToLower(ListString)
		}

		if StringToCheckInList == ListString {
			return true, MatchIndex
		}
	}

	return false, -1
}
