// Create a block of constants with three different types and display their values.
package main

import (
	"fmt"
)

func main() {
	const (
		const1 string  = "const1"
		const2 float32 = 99.90
		const3 rune    = 10
	)
	fmt.Println(const1, const2, const3)
}