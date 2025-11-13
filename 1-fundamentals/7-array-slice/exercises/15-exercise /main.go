// Implement logic that adds multiple values ​​to a dynamic collection at runtime.

package main

import (
	"fmt"
)

func fillSlice(slc []int, count int) []int {
	for index := 1; index <= count; index++ {
		slc = append(slc, index)
	}
	return slc
}

func main() {
	slc := []int{}
	fmt.Println(fillSlice(slc, 5))
}
