package pythagorean

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	result := make([]Triplet, 0)
	for a := min; a <= max-2; a++ {
		for b := a + 1; b <= max-1; b++ {
			for c := b + 1; c <= max; c++ {
				if isTriplet(a, b, c) {
					result = append(result, Triplet{a, b, c})
				}
			}
		}
	}

	return result
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	result := make([]Triplet, 0)
	aMax := int(p/3) - 1
	for a := 1; a <= aMax; a++ {
		bMax := (p - a) / 2
		for b := a + 1; b <= bMax; b++ {
			c := p - (a + b)
			if isTriplet(a, b, c) {
				result = append(result, Triplet{a, b, c})
			}
		}
	}

	return result
}

func isTriplet(a, b, c int) bool {
	if c < b || b < a {
		return false
	}

	return a*a+b*b == c*c
}
