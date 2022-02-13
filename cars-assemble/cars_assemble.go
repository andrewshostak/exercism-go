package cars

const carPrice = 10000
const tenCarsPrice = 95000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) / 100 * successRate
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRatePerHour int, successRate float64) int {
	return int(float64(productionRatePerHour) / 60 / 100 * successRate)
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	carsWithDefaultPrice := carsCount % 10
	carsWithEconomizedPrice := carsCount / 10
	return uint(carsWithEconomizedPrice*tenCarsPrice + carsWithDefaultPrice*carPrice)

}
