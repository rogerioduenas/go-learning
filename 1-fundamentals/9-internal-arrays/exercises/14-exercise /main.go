// Implement a code snippet that monitors and displays the capacity of a sequence before and after each expansion to visualize the runtime's reallocation strategy.
package main

import (
	"fmt"
)

func main() {
	slc := []int{1}
	prevCap := cap(slc)

	for i := 0; i <= 30; i++ {
		slc = append(slc, i)
		if cap(slc) != prevCap {
			fmt.Printf("Length: %d, Capacity changed: %d -> %d\n", len(slc), prevCap, cap(slc))
			prevCap = cap(slc)
		}
	}
}
