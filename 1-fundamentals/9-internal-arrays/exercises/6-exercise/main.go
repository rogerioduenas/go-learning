// Create a function that receives a dynamic sequence and returns its current size and capacity.
package main

import (
	"fmt"
)

func analyzeList(list []int) (int, int) {
	return len(list), cap(list)
}

func main() {

	slc := []int{1, 2, 3}

	fmt.Println(analyzeList(slc))
}
