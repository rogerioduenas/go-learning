// Create an integer variable, a simple pointer to it, a double pointer that stores the address of that simple pointer, and a triple pointer that stores the address of the double pointer. Implement a function that receives the triple pointer and changes its original integer value.
package main

import (
	"fmt"
)

func modifyValue(ptr ***int) {
	***ptr = 0
}

func main() {
	var a int = 1
	var ptr *int = &a
	var ptr2 **int = &ptr
	var ptr3 ***int = &ptr2

	modifyValue(ptr3)

	fmt.Println(a)
}
