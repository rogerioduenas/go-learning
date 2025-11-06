// Implement a generic function that receives any value and uses reflection to identify whether the argument is a number (integer or decimal). If it is, display the detected numeric type; otherwise, indicate that the value is not numeric.
package main

import (
	"fmt"
	"reflect"
)

func checkTypes(arg any) {
	value := reflect.TypeOf(arg).Kind()
	if value >= reflect.Int && value <= reflect.Float64 {
		fmt.Println(value)
		return
	}
	fmt.Println("It isn't a number!")
}

func main() {
	checkTypes(10.5)
	checkTypes(42)
	checkTypes("abc")
}
