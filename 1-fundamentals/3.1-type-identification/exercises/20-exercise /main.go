// Create a function that accepts any type of value, identifies the type at runtime, and returns a customized message based on that type.
package main

import (
	"fmt"
	"reflect"
)

func checkTypes(arg any) string {
	kind := reflect.TypeOf(arg).Kind()
	result := fmt.Sprintf("It's a %v with value: %v", kind, arg)
	return result
}

func main() {
	fmt.Println(checkTypes(10))
	fmt.Println(checkTypes("string"))
	fmt.Println(checkTypes(25.95))
}
