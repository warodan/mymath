package mymath

func Average(numbers []float64) float64 {
	var sum float64
	var count float64
	for _, number := range numbers {
		sum += number
		count++
	}
	return sum / count
}
