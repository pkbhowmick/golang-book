package main

import (
	"fmt"
	"math"
)

func findGreatestNumber(vals ...int) int {
	maxVal := - 1<<31
	for _,val := range vals{
		maxVal = int(math.Max(float64(maxVal),float64(val)))
	}
	return maxVal
}

func main()  {
	xs := []int{2,6,9,12,1,5}
	fmt.Println(findGreatestNumber(xs...))
}
