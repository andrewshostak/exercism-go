package lasagna

func OvenTime() int {
	return 40
}

func RemainingOvenTime(actualMinutes int) int {
	return OvenTime() - actualMinutes
}

func PreparationTime(numberOfLayers int) int {
	return 2 * numberOfLayers
}

func ElapsedTime(numberOfLayers, actualMinutes int) int {
	return PreparationTime(numberOfLayers) + actualMinutes
}
