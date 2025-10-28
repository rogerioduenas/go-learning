// Create a program that contains a block of variables and a block of constants within the main function, and display all the values.
package main

import (
	"errors"
	"fmt"
)

func multipleDatas() {
	const (
		const1 string  = "const1"
		const2 float32 = 99.90
		const3 rune    = 10
	)
	var (
		var1 string = "var1"
		var2 error  = errors.New("Any error")
		var3        = []int{1, 2, 3}
	)
	fmt.Println(const1, const2, const3, var1, var2, var3)
}
func main() {

	multipleDatas()
}