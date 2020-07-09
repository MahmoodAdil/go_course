package main

import "fmt"

/*
 anonymous functions, which can form closures. A
 nonymous functions are useful when you want to define a function inline
 without having to name it.*/
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func main() {
	sum := adder()

	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}

	newSum := adder()
	fmt.Println(newSum(88))
}
