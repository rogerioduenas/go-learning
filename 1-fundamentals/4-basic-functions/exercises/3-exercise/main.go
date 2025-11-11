// Display the value returned by the previous function using the standard output library.
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
