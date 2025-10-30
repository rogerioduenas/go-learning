// Create a function that combines different types (integers, floating-point numbers, and booleans) and displays their dynamic types using reflection.
package main

import (
	"fmt"
	"reflect"
)

func showTypes(arg1 int, arg2 float64, arg3 bool) {
	listOfArgs := []any{arg1, arg2, arg3}
	for _, value := range listOfArgs {
		fmt.Println(reflect.TypeOf(value).Kind())
	}
}

func main() {
	showTypes(10, 29.90, true)
}
