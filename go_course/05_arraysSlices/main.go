package main

import "fmt"

func main() {
	//Arrays
	var fruitArr [2]string

	//assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	//declare and assign array
	cars := [2]string{"BMW", "IBM"}

	//clice
	carsMake := []string{"BMW", "IBM", "Toyota", "Opel"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])
	fmt.Println(cars)
	fmt.Println(cars[1])
	fmt.Println(carsMake)
	//get length
	fmt.Println(len(carsMake))
	//range
	fmt.Println(carsMake[1:2])
}
