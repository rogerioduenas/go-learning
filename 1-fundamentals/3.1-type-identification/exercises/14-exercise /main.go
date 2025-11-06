// Create a function that compares two generic values ​​and displays messages indicating whether or not they are of the same type.
package main

import (
	"fmt"
	"reflect"
)

func checkTypes(arg1, arg2 any) {
	if reflect.TypeOf(arg1) == reflect.TypeOf(arg2) {
		fmt.Println("They are the same type")
		return
	}
	fmt.Println("They aren't the same type")
}

func main() {
	checkTypes(1, 2)
	checkTypes(1, true)
}
