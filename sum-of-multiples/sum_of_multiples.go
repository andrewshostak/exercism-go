package summultiples

func SumMultiples(limit int, divisors ...int) int {
	var sum int
	multiples := make(map[int]struct{})
	for n := 1; n < limit; n++ {
		for _, divisor := range divisors {
			if divisor == 0 || n%divisor != 0 {
				continue
			}

			_, ok := multiples[n]
			if !ok {
				multiples[n] = struct{}{}
				sum += n
			}
		}
	}

	return sum
}
