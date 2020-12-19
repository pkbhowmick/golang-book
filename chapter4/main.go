package main

import "fmt"

const PI float64 = 3.1416

var str string = "Global variable"

func main() {
	// Variables
	var x string
	x = "Hello World"
	fmt.Println(x)
	y := "Hello"
	fmt.Println(x == y)

	// Understand variable scope
	f()
}

func f() {
	fmt.Println(str)
	fmt.Println(PI)
	str = "another string"
	fmt.Println(str)
	// PI = 3.14 results in compile error
}
