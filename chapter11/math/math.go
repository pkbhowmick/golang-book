package math

//Finds the average of a series of numbers
func Average(xs []float64) float64 {
	total := float64(0)
	for _,val := range xs{
		total += val
	}
	return total / float64(len(xs))
}

// Finds minimum number from a series of numbers
func Min(xs []float64) float64 {
	minVal := float64(1<<31 - 1)
	for _,val := range xs {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

//Finds maximum number from a series of numbers
func Max(xs []float64) float64 {
	maxVal := float64(- 1<<31)
	for _,val := range xs {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}
