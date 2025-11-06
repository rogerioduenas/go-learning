// Create a function that receives a generic value and displays its actual type in a formatted output.
package main

import (
	"fmt"
	"reflect"
)

func checkType(data any) {
	fmt.Printf("It's a %T\n", data)
}
// or
func checkType2(data any) {
	kind := reflect.TypeOf(data).Kind()
	fmt.Println(kind)
}

func main() {
	checkType("Any text")
	checkType2(make(map[string]int))
}
