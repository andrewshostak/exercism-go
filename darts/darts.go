package darts

import (
	"math"
	"sort"
)

var radiusPointsRules = map[int]int{1: 10, 5: 5, 10: 1}
var radiuses = make([]int, 0, len(radiusPointsRules))

const centerX float64 = 0
const centerY float64 = 0

func init() {
	for radius := range radiusPointsRules {
		radiuses = append(radiuses, radius)
	}
	sort.Ints(radiuses)
}

func Score(x, y float64) int {
	a := math.Pow(x-centerX, 2) + math.Pow(y-centerY, 2)

	for _, radius := range radiuses {
		if a <= math.Pow(float64(radius), 2) {
			return radiusPointsRules[radius]
		}
	}

	return 0
}
