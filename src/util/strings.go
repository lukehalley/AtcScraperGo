package util

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// StringHelper provides common string manipulation utilities
func TitleCaseString(String string) string {
	Caser := cases.Title(language.English)
	return Caser.String(String)
}
