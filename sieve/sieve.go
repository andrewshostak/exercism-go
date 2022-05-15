package sieve

type numberWithMark struct {
	number int
	marked bool
}

func Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}

	// create slice with numbers in range from 2 to limit
	numbers := make([]numberWithMark, 0, limit-1)
	for i := 2; i <= limit; i++ {
		numbers = append(numbers, numberWithMark{number: i})
	}

	next := 2
	for next*next < limit {
		// mark non-prime numbers
		for i := next - 2; i < limit-1; i += next {
			if numbers[i].number > next && !numbers[i].marked {
				numbers[i].marked = true
			}
		}

		// find next non-marked number
		for i := range numbers {
			if !numbers[i].marked && numbers[i].number > next {
				next = numbers[i].number
				break
			}
		}
	}

	// filter primes
	var primes []int
	for i := range numbers {
		if !numbers[i].marked {
			primes = append(primes, numbers[i].number)
		}
	}

	return primes
}
