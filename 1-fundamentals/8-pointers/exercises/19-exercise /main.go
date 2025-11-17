// Create two separate string variables. Create a pointer that points to the first variable and a double pointer that points to that pointer. Write a function that replaces the address stored in the single pointer to point to the second variable, using only the double pointer.
package main

import (
	"fmt"
)

func modifyPtr(ptr **string, newPtr *string) {
	*ptr = newPtr
}

func main() {
	a := "a"
	b := "b"
	ptr := &a
	ptr2 := &ptr

	fmt.Println(ptr, *ptr)

	modifyPtr(ptr2, &b)

	fmt.Println(ptr, *ptr)
}
