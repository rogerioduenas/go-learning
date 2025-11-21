package main

import "fmt"

type user struct {
	name string
	age  int
}

func (u user) testMethod() {
	fmt.Println("executing test method")
}

func (u user) checkAge() bool {
	return u.age >= 18
}

func (u *user) addAge() {
	u.age++
}

func main() {
	u1 := user{"Caio", 25}
	fmt.Println(u1)

	u1.testMethod()

	fmt.Println(u1.checkAge())

	u1.addAge()
	u1.addAge()
	u1.addAge()
	fmt.Println("new age:", u1.age)
}
