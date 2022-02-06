package prime

import "math"

func Nth(n int) (int, bool) {
	foundPrimes := 0
	for i := 2; foundPrimes < n; i++ {
		if isPrime(i) {
			foundPrimes++
			if foundPrimes == n {
				return i, true
			}
		}
	}

	return 0, false
}

func isPrime(n int) bool {
	f := math.Sqrt(float64(n))

	for i := 2; i <= int(f); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
