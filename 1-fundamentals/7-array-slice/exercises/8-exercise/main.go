// Create a dynamic collection, initialize it with some values, add a new value to this dynamic collection, and display the result.

package main

import "fmt"

func main() {
	slc := []int{1, 2, 3}
	fmt.Println(slc)
	slc = append(slc, 4)
	fmt.Println(slc)
}
