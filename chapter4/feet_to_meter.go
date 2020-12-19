package main

import "fmt"

func main() {
	var f float64
	fmt.Print("Enter value(feet): ")
	fmt.Scanf("%f", &f)
	m := f * 0.3048
	fmt.Printf("Meter = %.2f\n", m)
}
