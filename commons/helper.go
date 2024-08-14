package commons

import "math"

func CalculateWeightedAverage(totalRemain float64, weightedAverage float64, newValue float64, quantity float64) float64 {
	return (totalRemain*weightedAverage + newValue*quantity) / (totalRemain + quantity)
}

func CalculatePercentage(number float64, percentage float64) float64 {
	if number < 0 {
		return 0
	}
	return number * percentage
}

func RoundUpTwoDigits(val float64) float64 {
	ratio := math.Pow(10, float64(2))
	return math.Round(val*ratio) / ratio
}
