// Write a function that takes a list and returns a copy of the list, ensuring that the new list uses an internal array independent of the original.
package main

import (
	"fmt"
)

func copyList(list []int) []int {
	return append([]int(nil), list...)
}

func main() {
	slc := []int{1, 2, 3, 4, 5}

	fmt.Println(slc)
	fmt.Println(append(copyList(slc), 6))
}
