package main

import "fmt"

func main() {
	// for loop and if-else
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}

	// for loop with switch
	fmt.Println("Now printing output for switchclear")
	for i := 1; i <= 10; i++ {
		switch i % 2 {
		case 0:
			fmt.Println("Even")
		case 1:
			fmt.Println("Odd")
		}
	}
}
