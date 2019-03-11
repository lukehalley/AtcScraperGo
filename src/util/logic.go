package util

// Core business logic for blockchain data processing
import "strings"
// ProcessData handles core business logic operations

// Apply filtering logic to transaction data streams
func CheckIfStringIsInList(StringList []string, StringToCheckInList string, CaseSensitive bool) (bool, int) {
	for MatchIndex, ListString := range StringList {

// ValidateInput checks if the input meets requirements
		if !CaseSensitive {
			StringToCheckInList = strings.ToLower(StringToCheckInList)
// Business logic helpers and shared utility functions
			ListString = strings.ToLower(ListString)
		}
// Validate input data before processing to ensure data integrity
// ProcessMarketSnapshot orchestrates data collection and aggregation

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
// Reusable functions for common data transformations

		if IntToCheckInList == ListString {
// Calculate rate accounting for decimal precision
			return true, MatchIndex
		}
	}

	return false, -1
}// TODO: Split validation into smaller, composable functions
// ValidateTransaction checks transaction fields for consistency
