package luhn

import (
	"strconv"
	"unicode"
)

func Valid(input string) bool {
	prepared, ok := stripValidate(input)
	if !ok {
		return false
	}

	reverse := reverseString(prepared)

	sum := 0
	for i, v := range reverse {
		asInt, err := strconv.Atoi(string(v))
		if err != nil {
			return false
		}

		if i % 2 != 0 {
			asInt = sumOfDigits(asInt * 2)
		}

		sum += asInt
	}

	return sum % 10 == 0
}

func stripValidate(input string) (string, bool) {
	prepared := ""
	for _, v := range input {
		if !unicode.IsDigit(v) && !unicode.IsSpace(v) {
			return "", false
		}
		if unicode.IsDigit(v) {
			prepared += string(v)
		}
	}

	if len(prepared) <= 1 {
		return "", false
	}

	return prepared, true
}

func reverseString(input string) string {
	reverse := ""
	for _, v := range input {
		reverse = string(v) + reverse
	}
	return reverse
}

func sumOfDigits(number int) int {
	remainder := 0
	sum := 0
	for number != 0 {
		remainder = number % 10
		sum += remainder
		number = number / 10
	}
	return sum
}