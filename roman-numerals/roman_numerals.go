package romannumerals

import (
	"fmt"
)

type rule struct {
	roman  string
	arabic int
}

func ToRomanNumeral(input int) (string, error) {
	min, max := 1, 3000
	if input < min || input > max {
		return "", fmt.Errorf("input should be between %d and %d", min, max)
	}

	rules := []rule{
		{roman: "M", arabic: 1000},
		{roman: "CM", arabic: 900},
		{roman: "D", arabic: 500},
		{roman: "CD", arabic: 400},
		{roman: "C", arabic: 100},
		{roman: "XC", arabic: 90},
		{roman: "L", arabic: 50},
		{roman: "XL", arabic: 40},
		{roman: "XXX", arabic: 30},
		{roman: "XX", arabic: 20},
		{roman: "X", arabic: 10},
		{roman: "IX", arabic: 9},
		{roman: "V", arabic: 5},
		{roman: "IV", arabic: 4},
		{roman: "I", arabic: 1},
	}

	num := input
	romanNumeral := ""
	for _, r := range rules {
		for num >= r.arabic {
			num -= r.arabic
			romanNumeral += r.roman
		}
	}

	return romanNumeral, nil
}
