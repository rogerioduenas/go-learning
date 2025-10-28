// Create a block of variables with three different identifiers, all of type text, and display their values.
package main

import (
	"fmt"
)

func main() {
	var var1, var2, var3 string = "var1", "var2", "var3"
	fmt.Println(var1, var2, var3)
}