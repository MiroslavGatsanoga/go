// Package weather provides utilities for a
// weather station's forecast program.
package weather

var (
	// CurrentCondition describes the current weather condition.
	CurrentCondition string
	// CurrentLocation describes a specified location.
	CurrentLocation string
)

// Forecast returns a statement about the weather
// based on a given city and weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
