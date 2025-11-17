// Implement a function that receives a memory address and redefines it to point to a different memory address.

package main

import (
	"fmt"
)

func changePtr(initialPtr **int, newPtr *int) {
	*initialPtr = newPtr
}

func main() {
	a := 10
	b := 0
	initialPtr := &a
	newPtr := &b

	fmt.Println(*initialPtr)
	changePtr(&initialPtr, newPtr)
	fmt.Println(*initialPtr)
}
