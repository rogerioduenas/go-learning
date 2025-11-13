// Create a collection of elements, and then a clipping of that collection. Modify an element in the original collection and view the clipping again to see how the change affects the shared elements.

package main

import "fmt"

func main() {

	slc := []string{"a", "b", "c", "d"}

	sliceSlc := slc[0:2]
	fmt.Println(sliceSlc)

	slc[0] = "--"
	fmt.Println(sliceSlc)

	sliceSlc[1] = "--"
	fmt.Println(sliceSlc)

	fmt.Println(slc)
}
