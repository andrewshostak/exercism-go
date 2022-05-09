package allyourbase

import (
	"errors"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if err := validateInput(inputBase, inputDigits, outputBase); err != nil {
		return nil, err
	}

	if inputBase == 10 {
		return fromDecimal(inputDigits, outputBase), nil
	} else if outputBase == 10 {
		return toDecimal(inputBase, inputDigits), nil
	}

	decimal := toDecimal(inputBase, inputDigits)
	return fromDecimal(decimal, outputBase), nil
}

func fromDecimal(inputDigits []int, outputBase int) []int {
	digit := sliceToInt(inputDigits)

	if digit == 0 {
		return []int{0}
	}

	result := make([]int, 0)
	for digit != 0 {
		result = append(result, digit%outputBase)
		digit = digit / outputBase
	}

	return reverseSlice(result)
}

func toDecimal(inputBase int, inputDigits []int) []int {
	var sum int
	for i, digit := range inputDigits {
		q := len(inputDigits) - 1 - i // 2, 1, 0 for 3-digit number
		result := digit * int(math.Pow(float64(inputBase), float64(q)))
		sum += result
	}

	if len(inputDigits) == 0 {
		return []int{0}
	}

	return intToSlice(sum)
}

func validateInput(inputBase int, inputDigits []int, outputBase int) error {
	if inputBase < 2 {
		return errors.New("input base must be >= 2")
	}

	for _, digit := range inputDigits {
		if digit < 0 || inputBase <= digit {
			return errors.New("all digits must satisfy 0 <= d < input base")
		}
	}

	if outputBase < 2 {
		return errors.New("output base must be >= 2")
	}

	return nil
}

func sliceToInt(s []int) int {
	res := 0
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= 10
	}

	return res
}

func intToSlice(i int) []int {
	digits := make([]int, 0)
	for i != 0 {
		digits = append(digits, i%10)
		i = i / 10
	}

	return reverseSlice(digits)
}

func reverseSlice(digits []int) []int {
	reversed := make([]int, 0, len(digits))
	for i := len(digits) - 1; i >= 0; i-- {
		reversed = append(reversed, digits[i])
	}

	return reversed
}
