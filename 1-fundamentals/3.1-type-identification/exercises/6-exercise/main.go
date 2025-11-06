// Develop a function that uses a generic type and displays the underlying type of the received value.
package main

import (
	"fmt"
	"reflect"
)

func checkTypes[T any](arg T) {
		fmt.Println(reflect.TypeOf(arg).Kind())
}

func main() {
	checkTypes(true)
	checkTypes("any text")
	checkTypes(25.50)
}
