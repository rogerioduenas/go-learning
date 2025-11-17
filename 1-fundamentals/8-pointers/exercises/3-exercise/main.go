// Modify the value of the original variable and display both again to see if the second one was affected.

package main

import "fmt"

func main() {
	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, var2)

	var1 = 5

	fmt.Println(var1, var2)
}
