// Demonstrate, through a clipping exercise, that slices share the same memory area as the underlying array.

package main

import (
	"fmt"
)

func main() {
	slc := []int{1, 2, 3, 4, 5}
	partSlc := slc[0:2]
	fmt.Println(slc)
	fmt.Println(partSlc)

	partSlc[0] = 0
	partSlc[1] = 0

	fmt.Println(slc)
	fmt.Println(partSlc)

}
