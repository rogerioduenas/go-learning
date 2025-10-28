// Create two variables, display their original values, swap their contents, and display them again.
package main

import (
	"fmt"
)

func main() {
	var var1, var2 = "var1", "var2"
	fmt.Println(var1, var2)
	var1, var2 = var2, var1
	fmt.Println(var1, var2)
}