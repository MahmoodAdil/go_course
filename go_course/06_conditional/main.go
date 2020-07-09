package main

import "fmt"

func main() {
	x := 10
	y := 10
	if x <= y {

		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	colour := "red"
	//if else
	if colour == "red" {
		fmt.Println("colour is red")
	} else if colour == "blue" {
		fmt.Println("colour is blue")
	} else {
		fmt.Println("colour is NOT Red or Blue")
	}

	//switch
	switch colour {
	case "red":
		fmt.Println("colour is red")
	case "blue":
		fmt.Println("colour is blue")
	default:
		fmt.Println("colour is NOT Red or Blue")

	}
}
