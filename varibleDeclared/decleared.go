package main

import "fmt"

func main() {
	// concept of variable declaration

	// initialized without the explicit type
	var myvariable1 = 20
	var myvariable2 = "GeeksforGeeks"

	fmt.Printf("The value of myvariable1 is : %d \n", myvariable1)
	fmt.Printf("The value of myvariable2 is : %s \n", myvariable2)

	// variable declared and initialized without expression

	var myVarInt int
	var myVarStr string

	fmt.Printf("The value of myVarInt %d \n", myVarInt)
	fmt.Printf("The value of myVarStr %s\n", myVarStr)

	// Multiple variables of the same type
	// are declared and initialized

	var myVarMulHomo1, myVarMulHomo2, myVarMulHomo3 int = 2, 4, 6
	var myVarMulHetero1, myVarMulHetero2, myVarMulHetero3 = 2, "GeeksforGeeks", 4
	fmt.Printf("Homogenous variable declarations %d %d %d\n", myVarMulHomo1, myVarMulHomo2, myVarMulHomo3)
	fmt.Printf("Heterogenous declarations %d, %s, %d\n", myVarMulHetero1, myVarMulHetero2, myVarMulHetero3)

}
