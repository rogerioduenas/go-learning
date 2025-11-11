// Reorganize os valores retornados pela função de múltiplos retornos, invertendo a ordem e exibindo os novos resultados.
package main

import (
	"fmt"
)

func multReturns() (int, int, int) {
	return 1, 2, 3
}
func main() {
	one, two, three := multReturns()
	fmt.Println(three, two, one)
}
