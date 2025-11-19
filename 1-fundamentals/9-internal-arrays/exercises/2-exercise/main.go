// Create a dynamic sequence and display its current values, along with its size and capacity.
package main

import (
	"fmt"
)

func main() {
	slc := make([]int, 3, 10)
	fmt.Println(slc, len(slc), cap(slc))
}
