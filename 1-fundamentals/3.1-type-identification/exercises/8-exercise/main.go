// Create a function that identifies whether a generic value is an integer, text, or other type, displaying a message for each case.
package main

import (
	"fmt"
)

func checkTypes(args ...any) {
	for _, value := range args {
		switch value.(type) {
		case int:
			fmt.Printf("It's a int: %v\n", value)
		case string:
			fmt.Printf("It's a string: %v\n", value)
		default:
			fmt.Printf("It's not an integer or a string: %v\n", value)
		}
	}
}

func main() {
	checkTypes([]int{1, 2, 3}, 10, 2.0, "any text", true)
}
