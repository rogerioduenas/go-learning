// Create two numeric variables with the same type and assign the value of one to the other in such a way that they represent an independent copy.

package main

import "fmt"

func main() {
	var var1 int = 10
	var var2 int = var1

	fmt.Println(var2 - 5)
	fmt.Println(var1)
}
