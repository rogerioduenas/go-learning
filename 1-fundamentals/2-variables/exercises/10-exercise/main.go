// Create two constants in a single declaration statement and display both.
package main

import (
	"fmt"
)

func main() {
	const const1, const2 = "const1", 18
	fmt.Println(const1, const2)
}