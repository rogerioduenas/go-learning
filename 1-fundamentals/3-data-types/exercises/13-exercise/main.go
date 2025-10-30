// Create a function that receives and returns a numerical value of a specific type, displaying the return value in the main program.
package main

import (
	"fmt"
)

func simpleReturn(number int) int{
	return number
}
func main() {
	fmt.Println(simpleReturn(10))
}
