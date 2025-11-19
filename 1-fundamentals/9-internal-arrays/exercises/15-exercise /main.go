// Create a simulator that adds elements to a sequence, displaying size and capacity at each iteration. The program should predict the next capacity expansion and stop the process if that expansion exceeds a defined maximum limit.
package main

import (
	"fmt"
)

func main() {
	slc := make([]int, 0)
	maxCap := 10

	for i := 0; i < 100; i++ {
		if cap(slc) == len(slc) && cap(slc)*2 > maxCap {
			break
		}

		slc = append(slc, i)
		fmt.Printf("Len: %v -- Capacity: %v\n", len(slc), cap(slc))

	}

	fmt.Println(slc, cap(slc))
}
