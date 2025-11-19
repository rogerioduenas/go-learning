// Create a dynamic sequence by explicitly setting the same value for size and capacity. Display its properties to compare its behavior with other sequences.
package main

import (
	"fmt"
)

func main() {
	slc := make([]int, 5)
	slc2 := make([]int, 3, 4)
	slc3 := []int{}


	fmt.Println(slc, len(slc), cap(slc))
	fmt.Println(slc2, len(slc2), cap(slc2))
	fmt.Println(slc3, len(slc3), cap(slc3))
}
