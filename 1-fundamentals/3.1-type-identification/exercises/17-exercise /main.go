// Develop a function that receives a generic value and performs different actions depending on whether it is an integer, decimal, or text.
package main

import (
	"fmt"
)

func checkTypes(arg any) {
	switch arg.(type) {
	case int, int8, int16, int32, int64:
		fmt.Println("It's an integer:", arg)
	case uint, uint8, uint16, uint32, uint64:
		fmt.Println("It's an unsigned integer:", arg)
	case float32, float64:
		fmt.Println("It's a decimal:", arg)
	case string:
		fmt.Println("It's a string:", arg)
	default:
		fmt.Println("Unsupported value:", arg)
	}
}

func main() {
	checkTypes(2.0)
}
