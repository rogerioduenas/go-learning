package main

import "fmt"

func approved(grade int) string {

	defer func() {
		if recover() != nil {
			fmt.Println("go study during the vacation")
		}
	}()

	if grade > 5 {
		return "passed the year"
	}
	panic("needs recovery classes")
}

func main() {
	fmt.Println(approved(4))
}
