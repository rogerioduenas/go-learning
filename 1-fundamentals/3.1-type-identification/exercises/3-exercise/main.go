// Create a function that receives different types of values ​​and displays the type of each one on separate lines.
package main

import (
	"fmt"
)

func checkTypes(args ...any) {
	for _, value := range args {
		fmt.Printf("The argument \"%v\" is of type %T\n", value, value)
	}
}
func main() {
	checkTypes("a", true, 10)
}
