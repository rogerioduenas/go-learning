// Implement a function that takes two integers as input and returns a specific function depending on a condition (e.g., if greater than 10, return addition; if less than 10, return subtraction). Call the returned function with the original arguments.
package main

import "fmt"

func fn(n1, n2 int) func() int {
	if n1 >= 10 {
		return func() int {
			return n1 + n2
		}
	}
	return func() int {
		return n1 - n2
	}
}

func main() {
	fmt.Println(fn(10, 5)())
	fmt.Println(fn(9, 5)())
}
