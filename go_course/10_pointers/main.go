package main

import "fmt"

func main() {

	// pointer allow you to point memory address of value
	a := 5
	b := &a // point to a
	fmt.Println(a, b)

	//use * to read value from memory address
	fmt.Println(*b)
	fmt.Println(*&a)

	// change value with pointer
	*b = 10
	fmt.Println(a, b)
}
