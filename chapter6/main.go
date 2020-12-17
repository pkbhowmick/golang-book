package main

import "fmt"

func main()  {
	//***** Array *******
	var arr [5]float64

	// taking input
	fmt.Println("Enter five values to print average: ")
	for i:=0; i <len(arr); i++{
		fmt.Scanf("%f",&arr[i])
	}

	//print average
	var total float64 = 0
	for i:=0; i<len(arr); i++{
		total+=arr[i]
	}
	avg := total/float64(len(arr))
	fmt.Println("Average:",avg)

	total = 0

	for _, val := range arr {
		total+=val
	}
	avg = total/float64(len(arr))
	fmt.Println("Average:",avg)

	//***** Slices **********

	slice1 := make([]float64,0)
	//var slice1 []float64

	//append some integers in the slice
	for _,val := range arr {
		slice1 = append(slice1,val)
	}

	//print slice1
	fmt.Println(slice1)

	slice2 := make([]float64, 2)
	copy(slice2,slice1)
	fmt.Println(slice1,slice2)

	//****** map ******
	map1 := make(map[string]int)
	map1["key1"] = 10
	map1["key2"] = 20

	if val, ok := map1["key1"]; ok {
		fmt.Println(val,ok)
	}
	if val, ok := map1["key3"]; ok {
		fmt.Println(val,ok)
	}

	// loop through all keys
	for k,v := range map1 {
		fmt.Println("key:",k,"val:",v)
	}


}
