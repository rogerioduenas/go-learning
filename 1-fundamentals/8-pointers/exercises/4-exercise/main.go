// Create a numeric variable and another variable intended to store a memory address compatible with it.

package main

import "fmt"

func main() {
	var var1 int = 10
	var var2 *int = &var1

	fmt.Println(var1, var2)
}
