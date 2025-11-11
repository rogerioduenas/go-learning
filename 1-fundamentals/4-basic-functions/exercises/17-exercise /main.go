// Create a function that receives a numerical value and returns a function capable of receiving two numbers and performing a mathematical operation combining the original value with the two parameters. Execute the returned function with different arguments.
package main

import (
	"fmt"
)

func f(n1 int) func(int, int) int {
	return func(i1, i2 int) int {
		return n1 + i1 + i2
	}
}

func main() {
	x := f(2)
	fmt.Println(x(1, 1))
	fmt.Println(x(5, 3))
	fmt.Println(x(10, -4))
}
