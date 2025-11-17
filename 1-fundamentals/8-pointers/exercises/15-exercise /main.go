// Create an example that clearly demonstrates the difference between changing the content pointed to and changing the address itself stored in a pointer variable.

package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 777

	ptr := &a

	fmt.Println(ptr)
	fmt.Println(*ptr)
	fmt.Println(a)

	fmt.Println("----------")

	*ptr = 5
	fmt.Println(*ptr)
	fmt.Println(a)

	fmt.Println("----------")

	ptr = &b
	fmt.Println(ptr)
	fmt.Println(*ptr)
	fmt.Println(b)

	fmt.Println("----------")

	*ptr = 000
	fmt.Println(*ptr)
	fmt.Println(b)
}
