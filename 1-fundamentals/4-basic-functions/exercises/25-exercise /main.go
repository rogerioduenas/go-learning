// Develop a function that returns an inner function capable of accumulating values ​​added to it with each call, returning the accumulated total. Demonstrate with multiple sequential calls.
package main

import "fmt"

func fn() func(int) int {
	var x int
	return func(number int) int {
		x += number
		return x
	}
}

func main() {
	acc := fn()
	fmt.Println(acc(5))
	fmt.Println(acc(5))
	fmt.Println(acc(5))
}
