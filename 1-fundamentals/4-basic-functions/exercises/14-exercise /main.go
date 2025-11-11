// Create a function that, when called, returns another function. The returned function should be able to be invoked later to calculate and provide a final result.
package main

import (
	"fmt"
)

func makeDiscount(percent float64) func(price float64) float64 {
	return func(price float64) float64 {
		return price * (1 - percent/100)
	}
}

func main() {
	tenPercentDiscount := makeDiscount(10)

	fmt.Println(tenPercentDiscount(200))
	fmt.Println(tenPercentDiscount(50))
}
