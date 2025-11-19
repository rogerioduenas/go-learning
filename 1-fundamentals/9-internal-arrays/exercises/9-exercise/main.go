// Create a sequence from another existing sequence and modify elements in the original sequence to analyze whether both share the same internal array.
package main

import (
	"fmt"
)

func main() {
	slc := make([]int, 3, 5)
	fmt.Println(slc)

	slc2 := slc
	fmt.Println(slc2)

	slc[0] = 1
	fmt.Printf("Original Slice: %v -- Copy Slice: %v\n",slc, slc2)
}
