package lsproduct

import (
	"errors"
)

func LargestSeriesProduct(digits string, span int) (int, error) {
	if span < 0 {
		return -1, errors.New("span must not be negative")
	}

	if span == 0 {
		return 1, nil
	}

	if span > len(digits) {
		return -1, errors.New("span must be smaller than string length")
	}

	ints := make([]int, 0, len(digits))
	for _, s := range digits {
		if s < '0' || s > '9' {
			return 0, errors.New("digits input must only contain digits")
		}

		ints = append(ints, int(s-'0'))
	}

	largest := 0
	for i := range ints {
		product := 1
		start := i
		end := i + span
		if end > len(ints) {
			break
		}
		for _, digit := range ints[start:end] {
			product *= digit
		}
		if product > largest {
			largest = product
		}
	}

	return largest, nil
}
