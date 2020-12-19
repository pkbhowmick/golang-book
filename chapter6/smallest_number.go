package main

import (
	"fmt"
	"math"
)

func main() {
	list := make([]int, 0)

	//take input
	fmt.Println("Enter five integers: ")
	for i := 1; i <= 5; i++ {
		var val int
		fmt.Scanf("%d", &val)
		list = append(list, val)
	}

	//find the smallest number in the list
	minVal := 1<<31 - 1

	for _, val := range list {
		minVal = int(math.Min(float64(minVal), float64(val)))
	}
	fmt.Println("Smallest number:", minVal)
}
