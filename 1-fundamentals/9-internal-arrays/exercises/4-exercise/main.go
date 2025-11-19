// Create a dynamic sequence with a defined initial capacity. Continue adding elements until you exceed that capacity and observe how the capacity is automatically adjusted.
package main

import (
	"fmt"
)

func main() {
	slc := make([]int, 0, 1)

	for i := 0; i < 10; i++ {
		slc = append(slc, 1)
		fmt.Println(slc, len(slc), cap(slc))
	}

}
