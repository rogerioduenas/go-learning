// Declare a floating-point variable with an explicit type and assign it the result of a mathematical operation.
package main

import (
	"fmt"
)

func main() {
	var discount float32 = 0.10
	var price float32 = 19.90
	finalPrice := price * (1 - discount)
	fmt.Println(finalPrice)
}
