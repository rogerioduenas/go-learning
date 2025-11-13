// Create a function that receives a dynamic collection and adds new elements while printing the size and capacity with each insertion.

package main

import (
	"fmt"
)

func showLenAndCap(slc []int, count int) {
	for index := 1; index <= count; index++ {
		slc = append(slc, index)
		fmt.Printf("Length: %v - Capacity: %v\n", len(slc), cap(slc))
	}
}

func main() {
	slc := []int{}
	showLenAndCap(slc, 9)
}
