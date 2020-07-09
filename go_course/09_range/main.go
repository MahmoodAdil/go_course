package main

import "fmt"

func main() {
	//array
	ids := []int{4, 7, 3, 6, 8, 9, 2}

	//loop through ids array
	for index, id := range ids {
		fmt.Printf("%d - ID: %d\n", index, id)
	}

	// without using index, if index declared but not use so it's give error
	//so to solve not use index use _ instead of index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// sum of ids
	sum := 0
	for _, x := range ids {
		sum += x
	}
	fmt.Println("Sum =", sum)

	// range with map
	users := map[string]string{
		"adil":  "adil@mahmood.ie",
		"alice": "alice@yahoo.com",
		"bob":   "bob@yahoo.com",
	}

	for k, v := range users {
		fmt.Printf("%s : %s\n", k, v)
	}

	// print only key
	for key := range users {
		fmt.Printf("Name: %s\n", key)
	}
	// print only value
	for _, value := range users {
		fmt.Printf("Email: %s\n", value)
	}
}
