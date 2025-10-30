// Create a function that receives an integer of type int8 and returns a message indicating whether the value is within or outside the allowed range for that type.
package main

import (
	"fmt"
)
func limitsOfInt8(number int8) (int8, string){
	if number > 127 || number < (-128) {
		//This line will never be executed.
		return number, "exceeds the limits of int8."
	}
	return number, "is accepted and is within the limits of int8."
}
func main() {
	fmt.Println(limitsOfInt8(-128))
}