package atbash

import (
	"bytes"
	"strings"
)

const (
	firstLetter = 'a'
	lastLetter  = 'z'

	firstNumber = '0'
	lastNumber  = '9'
)

func Atbash(s string) string {
	inLowerCase := strings.ToLower(s)
	var encoded []rune
	for _, r := range inLowerCase {
		if isLetter(r) {
			encoded = append(encoded, encodeLetter(r))
			continue
		}

		if isNumber(r) {
			encoded = append(encoded, r)
		}
	}

	var b bytes.Buffer
	for i, r := range encoded {
		if i != 0 && i%5 == 0 {
			b.WriteString(" ")
		}

		b.WriteRune(r)
	}

	return b.String()
}

func isLetter(r rune) bool {
	return r >= firstLetter && r <= lastLetter
}

func isNumber(r rune) bool {
	return r >= firstNumber && r <= lastNumber
}

func encodeLetter(r rune) rune {
	return lastLetter - (r - firstLetter)
}
