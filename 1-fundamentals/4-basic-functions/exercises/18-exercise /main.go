// Implement a function that takes another function as an argument, where the latter function returns two values. Inside the first function, call the received function, add the results, and return only the sum.
package main

import (
	"fmt"
)

func getPair(value1, value2 int) (int, int) {
	return value1, value2
}

func sumFrom(fn func(int, int) (int, int), x, y int) int {
	value1, value2 := fn(x, y)
	return value1 + value2
}

func main() {
	fmt.Println(sumFrom(getPair, 5, 5))
}
