package pangram

import "strings"

func IsPangram(sentence string) bool {
	lowered := strings.ToLower(sentence)
	for r := 'a'; r < 'z'; r++ {
		if !strings.ContainsRune(lowered, r) {
			return false
		}
	}

	return true
}
