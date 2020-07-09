package main

import "fmt"

func main() {
	// Define Map
	//emails := make(map[Key-Type]Value-Type)
	emails := make(map[string]string)
	//map-name["key"] = "value"
	emails["adil"] = "adil@mahmood.com"
	emails["alice"] = "alice@gmail.com"
	emails["bob"] = "bob@gmail.com"

	//decalre map and add Key-Value
	users := map[string]string{
		"adil":  "adil@mahmood.ie",
		"alice": "alice@yahoo.com",
		"bob":   "bob@yahoo.com",
	}
	users["mike"] = "mike@yahoo.com"

	fmt.Println(emails)
	fmt.Println(emails["bob"])
	//length of map
	fmt.Println(len(emails))
	//delete
	delete(emails, "bob")
	fmt.Println(emails)

	fmt.Println(users)

}
