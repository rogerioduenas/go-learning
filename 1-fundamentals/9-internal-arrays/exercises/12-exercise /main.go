// Develop an experiment that clearly demonstrates the difference between copying elements of a sequence and simply referencing the same internal array.
package main

import (
	"fmt"
)

func main() {
	slc := []int{1, 2, 3, 4, 5}
	refOfList := slc
	fmt.Printf("%p %p\n", &refOfList[0], &slc[0])

	fmt.Println("----------")

	slc2 := []int{1, 2, 3, 4, 5}
	onlyCopyOfList := append([]int(nil), slc2...)
	fmt.Printf("%p %p\n", &onlyCopyOfList[0], &slc2[0])
}
