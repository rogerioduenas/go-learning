// Create a function that receives values ​​in sequence and displays the type of each one using a dynamic type detection technique.
package main

import (
	"fmt"
	"reflect"
)

func checkTypes(args ...any) {
	for _, value := range args {
		valueType := reflect.TypeOf(value).Kind()
		fmt.Println(valueType)
	}
}

func main() {
	checkTypes(10.5, true, []any{1, "a", -2.5})
}
