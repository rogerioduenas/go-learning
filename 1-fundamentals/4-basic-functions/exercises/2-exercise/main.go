// Create a variable with an explicit type and assign to it the result of the function created earlier, within the program's entry point (main).
package main

import (
	"fmt"
)

func sum(n1, n2 int8) int8 {
	return n1 + n2
}

func main() {
	var result int8 = sum(10, 20)
	fmt.Println(result)
}
