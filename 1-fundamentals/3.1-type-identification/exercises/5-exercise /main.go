// Create a code snippet that calls a type checking function with values ​​of different types and displays the results.
package main

import (
	"fmt"
)

func checkTypes(args ...any) {
	for _, value := range args {
		fmt.Printf("%T\n", value)
	}
}

func setArgs(function func(...any), args ...any) {
	function(args...)
}

func main() {
	setArgs(checkTypes, 1, "A", true)
}
