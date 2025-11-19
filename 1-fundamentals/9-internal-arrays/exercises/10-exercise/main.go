// Create a newly initialized sequence and dynamically change its size and capacity, displaying the resulting behavior.
package main

import (
	"fmt"
)

func main() {
	slc := make([]int, 0)

	for i := 1; i <= 10; i++{
		slc = append(slc, i)
		fmt.Printf("Length: %v  -- Capacity: %v\n",len(slc), cap(slc))
	}
}
