package main

import "fmt"

func sum(n1 int, n2 int) int {
	return n1 + n2
}

func calc(n1, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {

	fmt.Println(sum(5, 5))
	fmt.Println(calc(20, 10))

	var f = func(str string) string {
		return str
	}
	fmt.Println(f("any text"))

	resultSum, resultSub := calc(100, 50)
	fmt.Printf("The result of sum is %d\nand sub is %d\n", resultSum, resultSub)

	resultSum2, _ := calc(5, 5)
	fmt.Printf("The result of sum2 is %d\n", resultSum2)
}
