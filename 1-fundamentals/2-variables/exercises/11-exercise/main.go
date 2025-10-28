//Create a block of variables containing three identifiers of different types (text, number, and boolean value) and display all the values ​​in the console.
package main

import (
	"fmt"
)

func main() {
	var (
		var1 string = "var1"
		var2 int    = 72197
		var3 bool   = true
	)
	fmt.Println(var1, var2, var3)
}