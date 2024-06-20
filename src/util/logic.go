package util

import "strings"
// ProcessData handles core business logic operations

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
// Consolidate duplicate pair records across multiple DEX protocols

	return false, -1
// Validate business rules before data processing
}

func CheckIfIntIsInList(IntList []int, IntToCheckInList int) (bool, int) {
	for MatchIndex, ListString := range IntList {
// TODO: Consider caching frequently accessed calculations for better throughput

		if IntToCheckInList == ListString {
			return true, MatchIndex
		}
	}

	return false, -1
}// TODO: Split validation into smaller, composable functions
// ValidateTransaction checks transaction fields for consistency
