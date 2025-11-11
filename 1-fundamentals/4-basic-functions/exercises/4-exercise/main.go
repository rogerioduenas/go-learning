// Declare another function that returns two results of the same type, representing two different operations with the same arguments.
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
	fmt.Println(calc(10, 20))
}
