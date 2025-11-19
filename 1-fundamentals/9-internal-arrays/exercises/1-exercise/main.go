// Create a dynamic sequence of integers with a defined initial size and a capacity greater than that size.
package main

import (
	"fmt"
)

func main() {
	slc := make([]int, 3, 10)
	fmt.Println(slc, len(slc), cap(slc))
}
