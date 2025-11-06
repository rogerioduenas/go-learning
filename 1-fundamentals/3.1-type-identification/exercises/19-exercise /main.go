// Implement a function that combines explicit type checking and reflection to double-check the type of a value.
package main

import (
	"errors"
	"fmt"
	"reflect"
)

func checkTypes(arg any) (string, error) {
	switch arg.(type) {
	case int:
		if reflect.TypeOf(arg).Kind() == reflect.Int {
			return "It's an int", nil
		}
	case string:
		if reflect.TypeOf(arg).Kind() == reflect.String {
			return "It's a string", nil
		}
	}
	return "", errors.New("Unsupported value")
}

func main() {
	fmt.Println(checkTypes(10))
	fmt.Println(checkTypes("string"))
	fmt.Println(checkTypes(25.95))
}
