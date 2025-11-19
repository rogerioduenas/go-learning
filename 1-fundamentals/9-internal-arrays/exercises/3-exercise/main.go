// Create a dynamic sequence and add a new element to it. Then, display the updated content, size, and capacity.
package main

import (
	"fmt"
)

func main() {
	slc := []int{}
	slc = append(slc, 1)
	fmt.Println(slc, len(slc), cap(slc))
}
