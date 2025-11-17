// Create an integer variable and a pointer that points to it. Then, create a pointer to that pointer. Implement a function that, using only the pointer to the pointer, changes the value of the original variable.
package main

import (
	"fmt"
)

func modifyVar(ptr **int) {
	**ptr = 0
}

func main() {
	a := 1
	ptr := &a
	ptr2 := &ptr

	modifyVar(ptr2)

	fmt.Println(a)
}
