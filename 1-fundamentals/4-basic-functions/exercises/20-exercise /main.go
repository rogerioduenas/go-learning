// Create a function that takes an integer and a conditional function (func(int) bool) as arguments. If the conditional function returns true, perform an operation and return the result; otherwise, return another fixed value.
package main

import "fmt"

func checkValue(value int, condition func(int) bool) bool {
	if condition(value) {
		return true
	}
	return false
}

func isAdult(age int) bool {
	if age >= 21 {
		return true
	}
	return false
}

func main() {
	fmt.Println(checkValue(21, isAdult))
}
