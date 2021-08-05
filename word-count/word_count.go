package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	inLowerCase := strings.ToLower(phrase)

	frequency := make(Frequency)
	word := ""
	for _, v := range inLowerCase {
		if unicode.IsLetter(v) || unicode.IsNumber(v) || v == '\'' {
			word += string(v)
			continue
		}

		if word != "" {
			frequency[strings.Trim(word, "'")]++
			word = ""
		}
	}

	if word != "" {
		frequency[strings.Trim(word, "'")]++
	}

	return frequency
}
