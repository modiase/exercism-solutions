// Package weather describes current weather conditions.
package weather

// CurrentCondition describes the current conditions at the current location.
var CurrentCondition string
// CurrentLocation describes the current location for which the conditions will be given.
var CurrentLocation string

// Forecast provides a forecast for the current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
