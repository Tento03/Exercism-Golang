package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var result=float64(productionRate)/100 * successRate
    return result
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var result=float64(productionRate) * (successRate/100)
    return int(result/60)
}

// CalculateCost works out the cost of producing the given number of cars.
// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	individual := carsCount % 10
	group := carsCount / 10

	result := (individual * 10000) + (group * 95000)
	return uint(result)
}
