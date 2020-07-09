package main

import "fmt"

func greeting(name string) /*return type*/ string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func getMultiply(num1, num2 int) int {
	return num1 * num2
}

func main() {
	fmt.Println(greeting("Adil"))
	fmt.Println(getSum(2, 4))
	fmt.Println(getMultiply(2, 4))
}
