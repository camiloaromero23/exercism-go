// Package weather provides general utilities for weather forecasting
package weather

// Stores the current weather condition value for a specified location
var CurrentCondition string

// Stores the current location value
var CurrentLocation string

// Forecast returns the current forecast for a specific city
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
