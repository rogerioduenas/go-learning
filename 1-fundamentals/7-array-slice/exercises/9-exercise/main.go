// Create a collection of text elements, and then create a second collection that contains only a selection of text elements from the first, limiting the range to two positions.

package main

import "fmt"

func main() {
	slc := []string{"a", "b", "c", "d"}
	sliceSlc := slc[0:2]
	fmt.Println(sliceSlc)
}
