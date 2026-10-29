package sieve

type candidate struct {
	value  int
	marked bool
}

func Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}

	candidates := make([]candidate, 0)
	primes := make([]int, 0)
	for i := 0; i <= limit; i++ {
		candidates = append(candidates, candidate{value: i})
	}

	for i, v := range candidates {
		if v.value < 2 {
			continue
		}

		if v.marked {
			continue
		}

		for j := v.value + i; j <= limit; j += i {
			candidates[j].marked = true
		}
	}

	for _, v := range candidates[2:] {
		if !v.marked {
			primes = append(primes, v.value)
		}
	}

	return primes
}
