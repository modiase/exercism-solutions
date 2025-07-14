package cars

// The size of batches
const batchSize = 10
// The price to produce batchSize cars
const batchPrice = 95000
// The price to produce one car
const unitPrice = 10000

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100.0)
}


func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60
}

func CalculateCost(carsCount int) uint {
	quotient, remainder := carsCount/10, carsCount%10
    return uint(quotient*batchPrice + remainder*unitPrice)

}
