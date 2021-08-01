package isbn

import (
	"strconv"
	"strings"
	"unicode"
)

func IsValidISBN(s string) bool {
	prepared := strings.ReplaceAll(s, "-", "")
	if len(prepared) != 10 {
		return false
	}

	checkCharacter := prepared[len(prepared)-1]
	if checkCharacter != 'X' && !unicode.IsNumber(rune(checkCharacter)) {
		return false
	}

	j := 10
	sum := 0
	for _, v := range prepared {
		if v == 'X' {
			sum += j * 10
			continue
		}

		if !unicode.IsNumber(v) {
			return false
		}

		asInt, err := strconv.Atoi(string(v))
		if err != nil {
			return false
		}
		sum += j * asInt
		j--
	}

	return sum%11 == 0
}
