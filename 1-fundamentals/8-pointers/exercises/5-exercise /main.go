// Associate the address of the numeric variable with the variable that stores the address and display the value of both.

package main

import "fmt"

func main() {
	var var1 int = 10
	var var2 *int = &var1

	fmt.Println(var1, var2)
}
