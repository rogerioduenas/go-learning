// Combine named and anonymous function calls so that the value returned by one serves as an argument for the other.
package main

import (
	"fmt"
)

func multiple() int {
	return 10
}

var result = func(f func() int) int {
	value := f()
	return value * 5
}(multiple)

func main() {
	fmt.Println(result)
}
