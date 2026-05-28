// Package weather provides weather reports.
package weather

var (
    // CurrentCondition is condition at your are.
	CurrentCondition string

	// CurrentLocation is area you lived in.
	CurrentLocation  string
)

// Forecast is the report that given to public.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
