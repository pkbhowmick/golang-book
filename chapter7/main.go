package main

import "fmt"

// Simple function
func average(list []float64) float64 {
	total := 0.0
	for _,val := range list {
		total += val
	}
	return total/float64(len(list))
}

// Returning multiple values
func multi(x,y int) (int,int) {
	return x+y,x-y
}

// Variadic functions
func add(args ...float64) float64 {
	total := 0.0
	for _,val := range args{
		total += val
	}
	return total
}

// Closure
func makeEvenGenerator() func() int {
	i := 0
	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}

// Recursion
func factorial(x uint) uint {
	if x==0 {
		return 1
	}
	return x * factorial(x-1)
}

// Defer (Go has a special statement called defer which schedules a function call to be run after the function completes)
func first()  {
	fmt.Println("First function")
}
func second()  {
	fmt.Println("Second function")
}
func deferExample()  {
	defer second()
	first()
}

// Panic & Recover
func panicExample()  {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

func main() {
	xs := []float64{2,3,1,6,7,8,9}
	fmt.Println(average(xs))
	fmt.Println(multi(3,2))
	fmt.Println(add(xs...))

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	fmt.Println(factorial(5))
	deferExample()
	panicExample()

}
