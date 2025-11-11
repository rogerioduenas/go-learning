// Create a function with two parameters of a small numeric type and a return value of the same type, performing a simple operation between the two values.
package main

import (
	"fmt"
)

func sum(n1, n2 int8) int8 {
	return n1 + n2
}

func main() {
	fmt.Println(sum(10, 20))
}
