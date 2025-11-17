// Change the original value of the numeric variable and display the value obtained through the address variable again.

package main

import "fmt"

func main() {
	var var1 int = 10
	var var2 *int = &var1

	var1 = *var2 - 2

	fmt.Println(*var2)
}
