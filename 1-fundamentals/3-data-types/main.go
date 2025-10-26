package main

import (
	"errors"
	"fmt"
)

func main() {

	//int8, int16, int32 and int64
	var number int = 10
	fmt.Println(number)

	//accepts only positive numbers and zero
	//uint8, uint16, uint32 and uint64
	var number2 uint = 20
	fmt.Println(number2)

	//alias
	var number3 rune = /*int32*/ 12345
	fmt.Println(number3)

	var number4 byte = /*uint8*/ 255
	fmt.Println(number4)

	//float32 and float64
	var number5 float32 = 99.90
	fmt.Println(number5)

	var str string = "uhsuashua"
	fmt.Println(str)

	var boolean bool = true
	fmt.Println(boolean)

	//stdlib errors
	var err error = errors.New("Any error")
	fmt.Println(err)
}
