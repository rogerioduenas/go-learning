// Modify the value stored in the numeric variable using indirect memory address access, and confirm that the original value has been changed.

package main

import "fmt"

func main() {
	var var1 int = 10
	var var2 *int = &var1

	*var2 = 0

	fmt.Println(var1)
}
