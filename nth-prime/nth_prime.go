package nthprime

import (
	"errors"
)

func Nth(n int) (int, error) {
	foundPrimes := 0
	for i := 2; foundPrimes < n; i++ {
		isPrime := IsPrime(i)
		if isPrime {
			foundPrimes++
		}

		if foundPrimes == n {
			return i, nil
		}
	}

	return 0, errors.New("n must be greater than 0")
}

func IsPrime(n int) bool {
	for j := 2; j <= n; j++ {
		if n == j {
			continue
		}

		if n%j == 0 {
			return false
		}
	}

	return true
}
