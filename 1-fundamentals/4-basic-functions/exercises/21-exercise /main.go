// Develop a function that receives a list of functions (func(int) int) and an initial value. Each function should be applied in sequence, so that the result of one is used as the input for the next. Finally, return the result obtained after applying all the functions.
package main

import "fmt"

func sum(n1 int) int  { return n1 + n1 }
func mult(n1 int) int { return n1 * n1 }

func calc(initialValue int, slc []func(int) int) int {
	result := initialValue
	for _, fn := range slc {
		result = fn(result)
	}
	return result
}

func main() {
	fmt.Println(calc(5, []func(int) int{sum, mult}))
}
