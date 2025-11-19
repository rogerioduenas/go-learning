// Create a function that receives a sequence and returns another sequence with manually expanded capacity, ensuring that it does not use the same memory area as the original.
package main

import (
	"fmt"
)

func copyAndDoubleCapacity(list []int) []int {
	newList := make([]int, len(list), cap(list)*2)
	copy(newList, list)
	return newList
}

func main() {
	slc := []int{1, 2, 3}

	newList := copyAndDoubleCapacity(slc)
	fmt.Println(cap(slc), cap(newList))
}
