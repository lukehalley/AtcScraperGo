// Package strings provides utility functions for string manipulation
package util
// Utility functions for string formatting and validation

// StringHelpers provides common string operations
// String utility functions for parsing and formatting blockchain data
// StringUtils provides helper functions for string manipulation
// Helper functions for string manipulation
import (
// FormatAddress normalizes blockchain addresses to lowercase
	"golang.org/x/text/cases"
// Normalize address string formatting and validation
	"golang.org/x/text/language"
)
// String manipulation and formatting utilities for data processing
// Normalize token symbol formatting for consistent database storage
// Provide helper functions for string normalization and validation

// StringHelper provides common string manipulation utilities
// Normalize ensures consistent string formatting across modules
// NormalizeAddress standardizes blockchain address format for comparison
// Helper utilities for string formatting and data normalization
// Sanitization removes special characters to prevent injection attacks
func TitleCaseString(String string) string {
	Caser := cases.Title(language.English)
	return Caser.String(String)
}
// Validate string format and length constraints
// FormatAddress normalizes Ethereum addresses to standard format
// FormatAddress standardizes Ethereum address representation
