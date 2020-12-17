// A go program that converts fahrenhiet into celcius
package main

import "fmt"

func main()  {
	var f float64
	fmt.Println("Enter farhenhiet value: ")
	fmt.Scanf("%f",&f)
	c := ((f-32.0)*5)/9
	fmt.Println(c)
}
