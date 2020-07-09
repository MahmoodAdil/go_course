package main

import (
	"fmt"
	"strconv"
)

//Person struct
type Person struct {
	/*	firstName string
		lastName  string
		city      string
		gender    string*/
	firstName, lastName, city, gender string
	age                               int
}

//Greeting method (value reciever)
func (p Person) greeting() string {
	return "Hello, my name is " + p.firstName + " and I am " + strconv.Itoa(p.age) + " years old"
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "F" {
		p.lastName = spouseLastName
	}
}

//hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}
func main() {

	//Init person using struct
	personOne := Person{firstName: "Alice", lastName: "Paul", city: "Dublin", gender: "M", age: 23}
	personTwo := Person{firstName: "Bob", lastName: "Mick", city: "London", gender: "M", age: 34}
	personThree := Person{firstName: "Ann", lastName: "Fred", city: "NYC", gender: "F", age: 40}
	//Alternative
	personFour := Person{"Marie", "Jon", "London", "F", 40}
	fmt.Println(personOne)
	fmt.Println(personTwo)
	fmt.Println(personThree)
	fmt.Println(personFour)
	/*----OUT-PUT----
	{Alice Paul Dublin M 23}
	{Bob Mick London M 34}
	{Ann Jon London F 40}
	{Sean Jon London F 40}
	*/
	fmt.Println(personOne.firstName)
	personOne.age++
	fmt.Println(personOne) //{Alice Paul Dublin M 24}
	fmt.Println(personOne.greeting())
	personOne.hasBirthday()
	fmt.Println(personOne)
	personThree.getMarried("Alice")
	personTwo.getMarried("Alice")
	fmt.Println(personThree) //{Ann Alice NYC F 40}
	fmt.Println(personTwo)   //Bob Mick London M 34}
}
