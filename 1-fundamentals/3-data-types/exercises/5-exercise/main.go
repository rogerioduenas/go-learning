// Create a floating-point numeric variable with an explicit type, convert it to an integer, and display the integer value.
package main

import (
	"fmt"
)

func main(){
	var priceOfMyDog float32 = 249.90
	var PriceConvertedToDecimal = int(priceOfMyDog)
	fmt.Println(PriceConvertedToDecimal)
}