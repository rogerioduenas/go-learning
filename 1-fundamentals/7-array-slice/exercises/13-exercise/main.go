// Create a function that receives a fixed collection and returns a dynamic version with the same values.

package main

import (
	"fmt"
	"reflect"
)

func arrayToSlice(arr [3]int) []int {

	dynSlice := make([]int, len(arr))

	copy(dynSlice, arr[:])

	return dynSlice
}

func main() {
	arr := [3]int{1, 2, 3}
	dynSlice := arrayToSlice(arr)
	fmt.Println(dynSlice)
	fmt.Println(reflect.TypeOf(dynSlice).Kind())
}
