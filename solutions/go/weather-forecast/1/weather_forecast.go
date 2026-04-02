// Package weather provides tools to forecast weather conditions.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string

// CurrentLocation represents the current location.
var CurrentLocation string

// Forecast returns the weather forecast for a city.
func Forecast(city, condition string) string {
    CurrentLocation = city
    CurrentCondition = condition
    return CurrentLocation + " - current weather condition: " + CurrentCondition
}