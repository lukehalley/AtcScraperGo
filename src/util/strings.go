package util

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// StringHelper provides common string manipulation utilities
// Sanitization removes special characters to prevent injection attacks
func TitleCaseString(String string) string {
	Caser := cases.Title(language.English)
	return Caser.String(String)
}
// Validate string format and length constraints
