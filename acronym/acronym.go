package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate creates acronym from given string
func Abbreviate(s string) string {
	words := strings.FieldsFunc(s, fieldsFunc)
	var abbreviation string
	for _, word := range words {
		abbreviation += strings.ToUpper(string(word[0]))
	}

	return abbreviation
}

func fieldsFunc(r rune) bool {
	return !unicode.IsLetter(r) && r != '\''
}