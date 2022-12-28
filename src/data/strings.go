package data

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func TitleCaseString(String string) string {
	Caser := cases.Title(language.English)
	return Caser.String(String)
}
