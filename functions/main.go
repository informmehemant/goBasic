package main

import "fmt"

// simple function functions 

//  call by value
func area(length, width int) int {
	Ar := length * width
	return Ar
}

// call by reference
func swap(a, b *int) int  {
	var o int 
	o = *a 
	*a = *b
	*b = o
	return 0
}


func main() {
	var a, b int = 2, 3
	fmt.Printf("Area of rectange is : %d", area(2,3))
	// call by reference

	fmt.Printf("not swaped a = %d , b = %d", a,b)
	swap(&a, &b)
	fmt.Printf("swaped a = %d , b = %d", a,b)

}