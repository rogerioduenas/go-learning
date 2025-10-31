package main

import "fmt"

func main() {

	fmt.Println("//array---------------------------")
	var arr [5]int
	fmt.Println(arr)

	var arr2 [3]string
	arr2[0] = "any text"
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	var arr4 = [3]string{"one", "two", "three"}
	fmt.Println(arr4)

	arr4[0] = "modified item"
	fmt.Println(arr4)

	arr5 := [...]int{1, 2, 3}
	fmt.Println(arr5)
	// arr5[3] = 4 --------error

	fmt.Println("//slice---------------------------")

	slc := []int{}
	fmt.Println(slc)

	slc2 := []int{1, 2, 3}
	fmt.Println(slc2)

	slc2 = append(slc2, 4)
	fmt.Println(slc2)

	slc3 := []string{"a", "b", "c", "d", "e"}
	slc4 := slc3[1:3] //Shows from pos 1 to pos 2, pos 3 is excluded.
	fmt.Println(slc4)
}
