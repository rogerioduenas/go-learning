package main

import "fmt"

func calcMath(n1, n2 int) (sum, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return // sum and sub have been returned
}

func main() {
	fmt.Println(calcMath(10, 5))
	fmt.Println(calcMath(5, 10))
	fmt.Println(calcMath(10, 10))
}
