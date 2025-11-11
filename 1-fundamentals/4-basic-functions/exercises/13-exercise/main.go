// Implement an anonymous function that uses a previously defined function within its body before returning a value.
package main

import (
	"fmt"
)

func discount(percent int) int {
	return percent
}

func main() {
	finalPrice := func(fullPrice float64, discountFunc func(int) int) float64 {
		return fullPrice * (1 - float64(discountFunc(7))/100)
	}(19.90, discount)

	fmt.Println(finalPrice)
}
