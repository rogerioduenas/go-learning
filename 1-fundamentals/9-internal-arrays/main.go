package main

import "fmt"

func main() {
	slc := make([]int, 2, 3)
	fmt.Println(slc)                   // [0, 0]
	fmt.Println("Size:", len(slc))     // len 2
	fmt.Println("Capacity:", cap(slc)) // cap 3

	fmt.Println("----------------")

	slc = append(slc, 5)
	fmt.Println(slc) // [0, 0, 5]

	fmt.Println("----------------")

	slc = append(slc, 6)
	fmt.Println(slc)                   // [0, 0, 5, 6]
	fmt.Println("Size:", len(slc))     // len 4
	fmt.Println("Capacity:", cap(slc)) // cap 6
}
