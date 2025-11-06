// Develop a function that receives a generic value and checks if it is an integer, displaying a corresponding message.
package main

import (
	"fmt"
	"reflect"
)

func isInt(data any) {
	if reflect.TypeOf(data).Kind() == reflect.Int {
		fmt.Println("It's a int.")
		return
	}
	fmt.Println("It isn't a int.")
}

// or
func isInt2(data any) {
	if _, ok := data.(int); ok {
		fmt.Println("It's an int.")
		return
	}
	fmt.Println("It isn't an int.")
}

func main() {
	isInt(-10)
	isInt2(25.50)
}
