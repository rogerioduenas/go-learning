// Implement a function that accepts different numeric types as parameters and displays the value and type of each argument.
package main

import (
	"fmt"
)

func multipleTypesOfNumbers(numbers ...any) {
	for _, value := range numbers {
		switch value.(type) {
		case int:
			fmt.Println("int", value)
		case float64:
			fmt.Println("float64", value)
		default:
			fmt.Println("It isn't a number")
		}
	}
}

func main() {
	multipleTypesOfNumbers(10, 25.20, -200)
}
