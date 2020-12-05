package main

import "fmt"

func main() {
	// Logical operators

	// Arithmetic Operators +,-,/,*, % 
	// Bitwise Operators &,|, ^,<<,>>
	// Assignment Operators  =, +=, -=,*=,%=,&=,^=,|=
	// Misc Operators


	p := 34
	q := 20

	
	result1 := p & q
	fmt.Printf("\n Result of p & q = %d", result1)
	
	result2 := p | q
	fmt.Printf("\n Result of p & q = %d", result2)

	//XOR
	result3 := p ^ q  
	fmt.Printf("\n Result of p ^ q = %d", result3)

	result4 := p << 1  
	fmt.Printf("\n Resilt of p << 1 = %d\n", result4)

	// Misc Operators

	a := 4
	// Using address of Operation(&) and
	// pointer indirection(*) operator

	b := &a
	fmt.Println(*b)
	*b = 7
	fmt.Println(a)



}