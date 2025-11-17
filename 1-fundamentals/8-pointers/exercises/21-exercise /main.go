// Create two distinct integer variables. Create a simple pointer pointing to one of them. Create a double pointer storing the address of this simple pointer. Create a triple pointer storing the address of the double pointer. Implement a function that receives the triple pointer and makes the simple pointer point to the other variable.
package main

import (
	"fmt"
)

func modifyPtr(ptr ***int, newPtr *int) {
	**ptr = newPtr
}

func main() {
	a := 1
	b := 2
	ptr := &a
	ptr2 := &ptr
	ptr3 := &ptr2

	fmt.Println(ptr, *ptr)
	modifyPtr(ptr3, &b)
	fmt.Println(ptr, *ptr)
}
