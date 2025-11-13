// Create a new collection based on another, but ensure that changes to the original do not modify the new one.

package main

import "fmt"

func main() {
	slc := []string{"a", "b", "c", "d"}
	slc2 := make([]string, len(slc))
	copy(slc2, slc)

	slc[0] = "-"

	fmt.Println(slc)
	fmt.Println(slc2)
}
