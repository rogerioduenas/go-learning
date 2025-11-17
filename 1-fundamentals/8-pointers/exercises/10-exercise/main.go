// Create a function that receives a memory address from a variable and directly changes the value contained in it.

package main

import "fmt"

func double(n1 *int) {
	*n1 = *n1 * 2
}

func main() {
	var1 := 10

	fmt.Println(var1)
	double(&var1)
	fmt.Println(var1)
}
