// Implement a function that receives a sequence, inserts elements progressively, and displays how the capacity increases non-linearly with each insertion.
package main

import (
	"fmt"
)

func fn(list []int) {
	for i := 1; i <= 10; i++ {
		list = append(list, i)

		fmt.Printf("Length: %v -- Capacity: %v\n", len(list), cap(list))
	}

}

func main() {
	slc := make([]int, 0, 1)
	fn(slc)
}
