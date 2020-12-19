package main

import (
	"fmt"
	m "golang-book/chapter11/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	min := m.Min(xs)
	max := m.Max(xs)
	fmt.Println(avg)
	fmt.Println(min)
	fmt.Println(max)
}
