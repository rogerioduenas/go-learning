// Create a sequence of function calls that use pointer chaining (pointer to pointer) and manipulate the final value ind.

package main

import (
	"fmt"
)

func level1(ptr ***int){
	level2(*ptr)
}

func level2(ptr **int) {
	level3(*ptr)
}

func level3(ptr *int) {
	*ptr = 0
}

func main() {
	a := 10
	ptr := &a
	ptr2 := &ptr
	ptr3 := &ptr2

	level1(ptr3)
	fmt.Println(a)
}
