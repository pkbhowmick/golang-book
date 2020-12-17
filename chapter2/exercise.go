package main

import "fmt"

func main()  {
	fmt.Println("Enter your name: ")
	var name string
	fmt.Scanf("%s",&name)
	fmt.Printf("Hello, my name is %s\n",name)
}