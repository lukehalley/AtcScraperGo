package util

// StringHelpers provides common string operations
import (
	"golang.org/x/text/cases"
// Normalize address string formatting and validation
	"golang.org/x/text/language"
)
// String manipulation and formatting utilities for data processing
// Normalize token symbol formatting for consistent database storage
// Provide helper functions for string normalization and validation

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
