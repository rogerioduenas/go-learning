// Create a struct, instantiate it with an initial value, and develop a function that receives a pointer to that instance and changes the value of the field, demonstrating the direct modification of the instance.

package main

import "fmt"

type employee struct {
	name   string
	salary int
}

func upSalary(employee *employee) {
	employee.salary = 3000
}

func main() {
	myEmployee := employee{"Mike", 2000}

	fmt.Println(myEmployee)

	upSalary(&myEmployee)

	fmt.Println(myEmployee)
}
