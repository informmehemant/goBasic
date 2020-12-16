package main

import "fmt"

// Function
func mul(a1, a2 int) int {
	res := a1 * a2
	fmt.Println("Result:  ", res)
	return 0
}

func show() {
	fmt.Println("Hello!, GeeksforGeeks ")
}

// main function

func main() {
	mul(23, 45)
	fmt.Println("Start")
	defer fmt.Println("End")
	defer mul(34, 56)
	defer mul(23, 56)
	show()
	
}
