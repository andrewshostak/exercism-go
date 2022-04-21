package lasagna

func PreparationTime(layers []string, avgTime int) int {
	if avgTime == 0 {
		avgTime = 2
	}

	return len(layers) * avgTime
}

func Quantities(layers []string) (int, float64) {
	var noodlesGm int
	var sauceL float64
	for _, layer := range layers {
		if layer == "noodles" {
			noodlesGm += 50
		} else if layer == "sauce" {
			sauceL += 0.2
		}
	}

	return noodlesGm, sauceL
}

func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities2Pt []float64, amount int) []float64 {
	modified := make([]float64, len(quantities2Pt))
	for i := range quantities2Pt {
		modified[i] = quantities2Pt[i] / 2 * float64(amount)
	}

	return modified
}
