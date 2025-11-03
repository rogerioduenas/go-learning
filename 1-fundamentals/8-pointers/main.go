package main

import "fmt"

func main() {

	// assignment by value
	var var1 int = 10
	var var2 int = var1
	fmt.Println(var1, var2) // 10 10

	var1++
	fmt.Println(var1, var2) // 11 10

	// assignment by reference (with pointers)
	var var3 int = 100
	var pointer *int

	pointer = &var3
	fmt.Println(var3, pointer) // 100 0xc0000100c8(address)
	fmt.Println(*pointer)      // 100 (value)

	var3 = 150
	fmt.Println(*pointer) // 150

	*pointer = 200
	fmt.Println(var3) // 200
}
