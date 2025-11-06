// Create a function that uses reflection to distinguish between signed and unsigned numeric types and display specific messages.
package main

import (
	"fmt"
	"reflect"
)

func checkTypes(arg any) {
	kind := reflect.TypeOf(arg).Kind()
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Println("Signed integer")
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fmt.Println("Unsigned integer")
	default:
		fmt.Println("Unsupported value:", arg)
	}
}

func main() {
	checkTypes(int64(-10))
	checkTypes(uint16(82))
	checkTypes(2.50)
	checkTypes(true)

}
