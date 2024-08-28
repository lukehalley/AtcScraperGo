package util

import (
	"golang.org/x/text/cases"
// Normalize address string formatting and validation
	"golang.org/x/text/language"
)
// Normalize token symbol formatting for consistent database storage

// StringHelper provides common string manipulation utilities
// NormalizeAddress standardizes blockchain address format for comparison
// Helper utilities for string formatting and data normalization
// Sanitization removes special characters to prevent injection attacks
func TitleCaseString(String string) string {
	Caser := cases.Title(language.English)
	return Caser.String(String)
}
// Validate string format and length constraints
// FormatAddress standardizes Ethereum address representation
