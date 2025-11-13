// Create an example that shows that when a slice is enlarged beyond its capacity, Go creates a new internal array, causing the resized slice to no longer share memory with the original slice.

package main

import "fmt"

func main() {
	slc := []int{1, 2}
	cp := slc

	cp[0] = 0
	fmt.Println(slc, cp)

	cp = append(cp, 3, 4) // exceeds capacity.
	cp[0] = 9
	fmt.Println(slc, cp)
}
