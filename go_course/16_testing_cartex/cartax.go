package main

import "fmt"

/* global variable declaration */
var valid bool
var yearOfManufacture int
var co2Pollution int
var tax int

//var todayYear = 2020

//carTax function
func carTax(year int) {
	if year >= 1900 /*&& year <= todayYear*/ {
		valid = true
		yearOfManufacture = year
	} else {
		valid = false
		co2Pollution = -1000
		tax = -1
	}
}

func setCO2(rating int) {
	if valid {
		co2Pollution = rating * 1000
	}
}
func getCO2() int {
	return co2Pollution / 1000
}

func getManufactureYear() int {
	if valid {
		return yearOfManufacture
	}
	return -1
}

func readAndZeroTax() int {
	temp := tax
	tax = 0
	return temp
}

func calculateTax(newSystem bool) int {
	if valid {
		if (yearOfManufacture <= 2010) ||
			((yearOfManufacture > 2010) && (newSystem)) {
			if newSystem {
				tax = 10 * (co2Pollution / 1000)
			} else {
				tax = 3000
			}
		}
	}
	return tax
}

func main() {
	carTax(2050)
	setCO2(-10)
	calculateTax(true)
	fmt.Println(tax)

}
