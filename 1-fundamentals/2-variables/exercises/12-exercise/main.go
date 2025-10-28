// Create two variables with different values, and then swap the values ​​between them.
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