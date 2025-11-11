// Develop a function that receives a multiplication factor and returns a function that accepts a slice of integers, applying the factor to each element and returning a new slice.
package main

import (
	"fmt"
)

func pow(pow int) func([]int) []int {
	return func(slc []int) []int {
		newSlice := make([]int, len(slc))
		for index, value := range slc {
			newSlice[index] = value * pow
		}
		return newSlice
	}
}

func main() {
	slc := []int{2, 3, 4}
	fmt.Println(pow(2)(slc))
}
