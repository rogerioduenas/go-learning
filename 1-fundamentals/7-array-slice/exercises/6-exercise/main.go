// Create two different collections, each initialized differently. Display the contents of both to observe the differences between slices and arrays.

package main

import "fmt"

func main() {
	slc := []int{1, 2, 3}
	array := [3]string{"a", "b", "c"}

	fmt.Println(slc)
	fmt.Println(array)
}
