package scrabble

import (
	"github.com/stretchr/stew/slice"
	"strings"
)

var pointsMap = map[int][]string{
	1:  {"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
	2:  {"D", "G"},
	3:  {"B", "C", "M", "P"},
	4:  {"F", "H", "V", "W", "Y"},
	5:  {"K"},
	8:  {"J", "X"},
	10: {"Q", "Z"},
}

func Score(input string) int {
	sum := 0
	prepared := strings.ToUpper(input)
	for key, letters := range pointsMap {
		for _, a := range prepared {
			if slice.Contains(letters, string(a)) {
				sum += key
			}
		}
	}

	return sum
}