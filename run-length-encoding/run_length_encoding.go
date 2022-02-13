package encode

import (
	"fmt"
	"strconv"
)

func RunLengthEncode(input string) string {
	var output string
	var group []rune

	for i, r := range input {
		if i == 0 {
			group = append(group, r)
			continue
		}

		lastGroupChar := group[len(group)-1]
		if lastGroupChar != r {
			output += encodeGroup(group)
			group = []rune{r}
			continue
		}

		group = append(group, r)
	}

	output += encodeGroup(group)
	return output
}

func RunLengthDecode(input string) string {
	var groups []string
	var group []byte

	var isPreviousNumber bool
	for i := 0; i < len(input); i++ {
		if i == 0 {
			group = append(group, input[i])
			isPreviousNumber = isNumber(input[i])
			continue
		}

		isNum := isNumber(input[i])

		switch true {
		case isNum && !isPreviousNumber:
			groups = append(groups, string(group))
			group = []byte{input[i]}
		case isNum && isPreviousNumber:
			group = append(group, input[i])
		case !isNum && isPreviousNumber:
			group = append(group, input[i])
		case !isNum && !isPreviousNumber:
			groups = append(groups, string(group))
			group = []byte{input[i]}
		}

		isPreviousNumber = isNum
	}

	if len(input) > 0 {
		groups = append(groups, string(group))
	}

	var output string
	for i := range groups {
		output += decodeGroup(groups[i])
	}

	return output
}

func encodeGroup(group []rune) string {
	if len(group) > 1 {
		return fmt.Sprintf("%d%s", len(group), string(group[0]))
	} else if len(group) == 1 {
		return string(group[0])
	}

	return ""
}

func isNumber(n byte) bool {
	return n >= '0' && n <= '9'
}

func decodeGroup(s string) string {
	lastCharacter := s[len(s)-1]
	numberOfCharacters, err := strconv.Atoi(s[:len(s)-1])
	if err != nil {
		return string(lastCharacter)
	}

	var decoded string
	for i := 0; i < numberOfCharacters; i++ {
		decoded += string(lastCharacter)
	}

	return decoded
}
