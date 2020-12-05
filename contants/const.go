package main

import "fmt"

// PI Constant used as global constants
const PI = 3.14 // Constant PI
func main() {
	const trueConst = true
	const GFG = "GeeksforGeeks"
	fmt.Println("Hello", GFG)
	fmt.Println("Happy", PI, "Day")
	const Correct = true
	fmt.Println("Go rules? ", Correct)
	
	// var VariableName Type
	// or create a type

	type myBool bool
	var defaultBool = trueConst // allowed
	var customBool myBool = trueConst // allowed
	// defaultBool = customBool // not allowed
	fmt.Println(defaultBool)
	fmt.Println(customBool) 
}
