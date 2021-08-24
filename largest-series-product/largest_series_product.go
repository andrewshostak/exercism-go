package lsproduct

import (
	"errors"
	"strconv"
	"strings"
)

func LargestSeriesProduct(digits string, span int) (int, error) {
	if span < 0 {
		return -1, errors.New("span must be greater than zero")
	}

	if span > len(digits) {
		return -1, errors.New("span must be smaller than string length")
	}

	splitted := strings.Split(digits, "")
	ints := make([]int, 0, len(splitted))
	for i := range splitted {
		asInt, err := strconv.Atoi(splitted[i])
		if err != nil {
			return -1, errors.New("digits input must only contain digits")
		}
		ints = append(ints, asInt)
	}

	product := 0
	for i := 0; i+span <= len(ints); i++ {
		sum := 1
		part := ints[i : i+span]
		for j := range part {
			sum *= part[j]
		}
		if sum > product {
			product = sum
		}
	}

	return product, nil
}
