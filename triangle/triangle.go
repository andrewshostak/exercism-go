package triangle

import "math"

type Kind int

const (
	NaT = iota
	Equ
	Iso
	Sca
)

func KindFromSides(a, b, c float64) Kind {
	for _, side := range [3]float64{a, b, c} {
		if math.IsNaN(side) || math.IsInf(side, 1) || math.IsInf(side, -1) || side <= 0 {
			return NaT
		}
	}

	if a == b && b == c {
		return Equ
	}

	if a+b < c || b+c < a || c+a < b {
		return NaT
	}

	if a == b || b == c || c == a {
		return Iso
	}

	return Sca
}
