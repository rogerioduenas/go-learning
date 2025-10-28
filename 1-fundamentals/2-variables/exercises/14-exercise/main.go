// Create a program that declares two variables and two constants within the same function and displays them all in a single output statement.
package main

import (
	"fmt"
)

func main() {

	var var1, var2 = "var1", true
	const const1, const2 = "const1", 2.50

	fmt.Println(var1, var2, const1, const2)
}