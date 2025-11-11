// Create a function that takes a function as an argument (func(int) int) and an integer. Internally, add 5 to the integer, call the callback function, and return the result. Demonstrate how the outer function "captures" the modified variable before passing it to the callback.
package main

import "fmt"

func process(number int) int {
	return number
}

func apply(process func(int) int, number int) int {
	modified := number + 5
	return process(modified)
}

func main() {
	fmt.Println(apply(process, 5))
	fmt.Println(apply(process, 20))
}
