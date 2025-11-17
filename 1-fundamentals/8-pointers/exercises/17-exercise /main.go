// Create an integer variable, a pointer to it, and another pointer that stores the address of that pointer. Implement a function that receives this pointer to the pointer and makes the original pointer point to another integer variable.

package main

import (
	"fmt"
)

func modifyPtr(ptr **int, newAddress *int) {
	*ptr = newAddress
}

func main() {
	a := 111
	b := 222
	ptr := &a
	ptr2 := &ptr

	fmt.Printf("Before: %v --- %v\n", ptr, *ptr)
	modifyPtr(ptr2, &b)
	fmt.Printf("After: %v --- %v\n", ptr, *ptr)
}
