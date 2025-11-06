// Develop a function that identifies the type of a variable and returns a string describing it.
package main

import (
	"fmt"
	"reflect"
)

func checkTypes(arg any) string {
	kind := reflect.TypeOf(arg).Kind()
	formatted := fmt.Sprintf("My argument \"%v\" is a %s", arg, kind)

	return formatted
}

func main() {
	fmt.Println(checkTypes([2]int{1,2}))
	fmt.Println(checkTypes([]int{1,2}))
	fmt.Println(checkTypes("Any text"))
}
