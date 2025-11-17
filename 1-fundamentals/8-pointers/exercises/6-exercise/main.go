// Display in the output the contents of the address variable and also the value it points to.

package main

import "fmt"

func main() {
	var var1 int = 10
	var var2 *int = &var1

	fmt.Println(var2, *var2)
}
