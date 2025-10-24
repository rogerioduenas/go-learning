package main

import (
	"fmt"
	"modulo/helper"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("writing from the main file")
	helper.Write()

	erro := checkmail.ValidateFormat("rogerio@gmail.com")
	fmt.Println(erro)
}
