// Implement a function that receives a value and displays its type, recognizing only int, string, and bool. For other types, display "Unsupported type".
package main

import (
	"fmt"
)

func checkTypes(arg any) {
	switch arg.(type) {
	case int, string, bool:
		fmt.Printf("It's a %T\n", arg)
	default:
		fmt.Println("Unsupported type")
	}
}

func main() {
	checkTypes(1)
	checkTypes("a")
	checkTypes(true)
	checkTypes([]int{})
}
