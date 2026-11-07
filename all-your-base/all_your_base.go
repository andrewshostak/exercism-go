package allyourbase

import "errors"

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}

	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	for _, v := range inputDigits {
		if v >= inputBase || v < 0 {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
	}

	asDecimal := toDecimal(inputBase, inputDigits)
	asOutputBase := decimalToBase(asDecimal, outputBase)

	return asOutputBase, nil
}

func toDecimal(inputBase int, inputDigits []int) int {
	result := 0
	for _, v := range inputDigits {
		result = result*inputBase + v
	}

	return result
}

func decimalToBase(decimal, base int) []int {
	if decimal == 0 {
		return []int{0}
	}

	digits := []int{}
	for decimal > 0 {
		digit := decimal % base
		digits = append(digits, digit)
		decimal = decimal / base
	}

	reversed := make([]int, 0, len(digits))
	for i := len(digits) - 1; i >= 0; i-- {
		reversed = append(reversed, digits[i])
	}

	return reversed
}
