package cars

const individualCarCost = 10000
const groupedCarCost = 9500

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(float64(productionRate) * float64(successRate/100))
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(float64(productionRate) * float64(successRate/100) / 60)
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	groupedCars := carsCount - (carsCount % 10)

	return uint((groupedCars * groupedCarCost) + ((carsCount % 10) * individualCarCost))
}
