// Create a function that uses reflection to identify the type of a value and return both the name and the category of the type.
package main

import (
	"fmt"
	"reflect"
)

func checkTypes(arg any) (string, string) {
	tp := reflect.TypeOf(arg)
	kind := tp.Kind()
	return tp.String(), kind.String()
}

func main() {
	fmt.Println(checkTypes([1]int{}))
}
