// Create a routine that fills a sequence until it reaches its initial maximum capacity, and then try adding more elements to observe the internal reallocation.
package main

import (
	"fmt"
)

func main() {
	slc := make([]int, 0, 5)

	for i := 1; i <= cap(slc); i++ {
		slc = append(slc, i)

		if cap(slc) == len(slc) {
			fmt.Printf("Slice: %v - Cap: %v - Len: %v\n", slc, cap(slc), len(slc))
		}

	}

	fmt.Println("After exceeding the capacity")
	slc = append(slc, 6)
	fmt.Printf("Slice: %v - Cap: %v - Len: %v\n", slc, cap(slc), len(slc))
}
