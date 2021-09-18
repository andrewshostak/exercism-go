// Package weather represents weather condition and location
package weather

var (
	// CurrentCondition represents current weather conditions
	CurrentCondition string
	// CurrentLocation represents current location
	CurrentLocation string
)

// Log returns weather condition in a particular city and sets current variables
func Log(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
