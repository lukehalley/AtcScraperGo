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

func CheckIfIntIsInList(IntList []int, IntToCheckInList int) (bool, int) {
	for MatchIndex, ListString := range IntList {

		if IntToCheckInList == ListString {
			return true, MatchIndex
		}
	}

	return false, -1
}