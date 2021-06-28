package grains

import "errors"

const numberOfSquares int = 64

func Square(squareNumber int) (uint64, error) {
	if squareNumber < 1 || squareNumber > numberOfSquares {
		return 0, errors.New("invalid square number")
	}

	return 1 << uint64(squareNumber - 1), nil
}

func Total() uint64 {
	return 1 << numberOfSquares - 1
}

