package main

import "fmt"

func swap(xPtr, yPtr *int) {
	temp := *xPtr
	*xPtr = *yPtr
	*yPtr = temp
}

func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println(x, y)
}
