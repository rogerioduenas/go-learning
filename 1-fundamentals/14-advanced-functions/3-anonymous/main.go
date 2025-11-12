package main

import "fmt"

func main() {

	func() {
		fmt.Println("Hello buddy")
	}()

	sum := func(n1, n2 int) (result int) {
		result = n1 + n2
		return
	}(2, 2)
	fmt.Println(sum)

	sub := func(n1, n2 int) (result int) {
		result = n1 - n2
		return
	}
	fmt.Println(sub(10, 5))
}
