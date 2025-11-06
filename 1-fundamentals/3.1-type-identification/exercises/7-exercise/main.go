// Create a function that uses the reflection package to determine the type of a variable and display the result.
package main

import (
	"fmt"
	"reflect"
)

func checkTypes[T any](arg T) {
	value := reflect.TypeOf(arg)
	fmt.Printf("Type: %v | Kind: %v\n", value, value.Kind())
}

func main() {
	checkTypes([]int{1, 2, 3})
}
