package bob

import (
	"strings"
	"unicode"
)

// Hey returns bob response
func Hey(remark string) string {
	trimmed := strings.TrimSpace(remark)
	if len(trimmed) == 0 {
		return "Fine. Be that way!"
	}

	capitalized := isCapitalized(trimmed)
	question := isQuestion(trimmed)

	if question && capitalized {
		return "Calm down, I know what I'm doing!"
	} else if question && !capitalized {
		return "Sure."
	} else if capitalized && !question {
		return "Whoa, chill out!"
	}

	return "Whatever."
}

func isCapitalized(s string) bool {
	letters := ""
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters += string(r)
		}
	}

	if len(letters) == 0 {
		return false
	}

	return strings.ToUpper(letters) == letters
}

func isQuestion(s string) bool {
	return s[len(s) - 1] == '?'
}
