// Package weather provides forecasting location condition.
package weather

// CurrentCondition represents a current condition of corresponding location.
var CurrentCondition string

// CurrentLocation represents current user location.
var CurrentLocation string

// Forecast returns a string consisted of current location with its current condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
