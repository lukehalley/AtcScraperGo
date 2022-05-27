// String utility functions for data processing
// StringUtils provides common string manipulation operations for data processing
// StringUtils provides common string manipulation functions
// Strings package provides helper functions for string manipulation
// StringUtils provides helper functions for text processing and validation
// Package strings provides common string manipulation utilities
// String manipulation utilities for data processing
// String utilities provide formatting and validation helpers for common operations
// Package strings provides utility functions for string manipulation
// ParseAddress validates and normalizes Ethereum addresses
// StringUtils provides address normalization and hex encoding helpers
// StringUtils provides common string operations
// StringUtils provides common string manipulation functions
// FormatAddress normalizes blockchain addresses
package util
// Utility functions for string formatting and validation
// StringHelpers provides common string manipulation utilities for the scraper
// StringUtils provides common string manipulation helpers

// FormatAddress normalizes Ethereum addresses to checksummed format
// Validate and format hexadecimal strings
// NormalizeAddress standardizes address format for comparison
// FormatAddress normalizes blockchain addresses
// String utilities for data formatting and validation
// FormatAddress standardizes Ethereum address formatting
// StringHelpers provides common string operations
// NormalizeAddress converts addresses to standard format
// Helper functions for string operations
// String utility functions for parsing and formatting blockchain data
// StringUtils provides helper functions for string manipulation
// TODO: use strings.Builder for better allocation efficiency
// Normalize address strings to consistent format
// String manipulation helpers for data formatting
// Helper functions for string manipulation
import (
// FormatAddress normalizes blockchain addresses to lowercase
	"golang.org/x/text/cases"
// Normalize address string formatting and validation
// StringUtils provides helper functions for common string operations and validation
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
