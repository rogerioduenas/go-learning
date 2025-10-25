package main

import "fmt"

func main() {
	// Explicit variable declaration with defined type
	var var1 string = "var1"
	// Explicit constant declaration with defined type
	const const1 string = "const1"
	// Short variable declaration (type inferred automatically)
	var2 := "var2"
	fmt.Println(var1, const1, var2)

	// Block declaration of multiple variables with explicit type
	var (
		var3 string = "var3"
		var4 string = "var4"
	)
	// Block declaration of multiple constants with inferred type
	const (
		const2 = "const2"
		const3 = "const3"
	)
	fmt.Println(var3, var4, const2, const3)

	// Multiple variable declaration on the same line with inferred type
	var var5, var6 = "var5", "var6"
	// Multiple variable declaration on the same line with explicit type
	var var7, var8 string = "var7", "var8"
	// Short multiple variable declaration with inferred type
	var9, var10 := "var9", "var10"
	fmt.Println(var5, var6, var7, var8, var9, var10)

	// Multiple constant declaration with inferred type
	const const4, const5 = "const4", "const5"
	// Multiple constant declaration with explicit type
	const const6, const7 string = "const6", "const7"
	fmt.Println(const4, const5, const6, const7)

	var varA, varB string = "A", "B"
	// Swap values between variables
	varA, varB = varB, varA
	fmt.Println(varA, varB)
}
