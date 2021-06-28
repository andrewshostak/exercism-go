package isogram

import "strings"

func IsIsogram(input string) bool {
	prepared := strings.ToUpper(input)
	
	lettersMap := map[int32]bool{}
	for _, r := range prepared {
		if r < 'A' || r > 'Z' {
			continue
		}

		if _, ok := lettersMap[r]; !ok {
			lettersMap[r] = true
		} else {
			return false
		}
	}
	
	return true
}