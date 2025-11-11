// Create a function that receives two numerical parameters and returns two results of the same type, representing two different operations with those values. At the program's entry point (main), call this function and store the two results in separate variables.
package main

import (
	"fmt"
)

func calc(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	sum, sub := calc(30, 20)
	fmt.Println(sum, sub)
}
