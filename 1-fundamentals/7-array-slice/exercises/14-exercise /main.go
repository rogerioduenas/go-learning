// Create a function that receives a dynamic collection and returns a new collection containing only a portion of the original elements.

package main

import (
	"fmt"
)

func subSlice(slc []int, start int, end int) []int {

	if start < 0 {
		start = 0
	}
	if end > len(slc) {
		end = len(slc)
	}
	if start > end {
		start = end
	}

	listLength := end - start
	clippedList := make([]int, listLength)

	copy(clippedList, slc[start:end])

	return clippedList
}

func main() {
	slc := []int{1, 2, 3, 4, 5}
	fmt.Println(subSlice(slc, 2, 4))
}
